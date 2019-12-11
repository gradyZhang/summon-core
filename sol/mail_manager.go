package sol

import (
	"sync"
	"database/sql"
	"github.com/gradyZhang/summon-core/aliens/common/cache"
	"github.com/gradyZhang/summon-core/aliens/log"
	"strings"
	"time"
	"encoding/json"
	"strconv"
)

const (
	MAIL_SQL_CMD_INSERT int32 = 1
	MAIL_SQL_CMD_UPDATE int32 = 2
)

const (
	MAIL_SP_SELECT_MAIL      string = "CALL sp_select_mail(?,?)"                 //搜索邮件数据
	MAIL_SP_SELECT_MAIL_USER string = "CALL sp_select_mail_user( ?,? )"          //搜索玩家的个人邮件领取状态
	MAIL_SP_INSERT_MAIL      string = "CALL sp_insert_mail(?,?,?,?,?,?,?,?,?,?,?)" // 插入邮件
	MAIL_SP_UPDATE_MAIL             = "CALL sp_update_mail(?,?)"                 // 更新邮件记录
)

const (
	REDIS_MAIL        = "rM"          //邮件 redis 的key
	REDIS_MAIL_USER   = "rMU"         //邮件 redis 用户的 key
	REDIS_MAIL_GLOBAL = "rMailGolbal" //邮件  邮件全局id
)

const (
	REDIS_FILED_MAIL      = "fM" // 邮件redis的filed
	REDIS_FILED_MAIL_USER = "fMU"
)

const (
	MAIL_FLAG_UNREAD   int32 = 0 //未读
	MAIL_FLAG_READ     int32 = 1 //已读
	MAIL_FLAG_RECEIVED int32 = 2 //已领取
	MAIL_FLAG_DELET    int32 = 3 //已删除
)

type MailFlag struct {
	ID   int64 //邮件唯一ID
	Flag int32 //邮件操作标识
}

func (m *MailFlag) parse(str string) { //解析
	arrStr := strings.Split(str, "-")
	if arrStr == nil {
		return
	}
	if len(arrStr) < 2 {
		return
	}
	m.ID, _ = strconv.ParseInt(arrStr[0], 10, 64)
	flag, _ := strconv.Atoi(arrStr[1])
	m.Flag = int32(flag)
}

func (m *MailFlag) form() string {
	return strconv.FormatInt(m.ID, 10) + "-" + strconv.Itoa(int(m.Flag)) + "#"
}

type Mail struct {
	ID         int64   //邮件唯一ID
	SrcUid     int64   //发送者 uid
	SrcName    string  //发送者昵称
	DecUid     []int64 //目标的 uid ( 注: 群发邮件为 len = 0 ) 可能是多人邮件
	Title      string  //标题
	Content    string  //内容
	Gift       string  //附件内容
	SendTime   int64   //时间戳
	ExpiryTime int64   //到期时间
	Type       int32   // 邮件类型
	Condition  string  // 条件
}

type MailManager struct {
	dbBase      *sql.DB                 //sql 句柄
	dbBaseInfo  string                  //sql 信息
	redisClient *cache.RedisCacheClient //redis 句柄

	channel     chan interface{}
	isChanClose bool

	maxId int64 //邮件的最大 id
	sync.RWMutex
}

type MailSqlCmd struct {
	//邮件数据库指令结构
	Type int32         //命令类型
	Args []interface{} //参数数据
}

var mailMgr *MailManager
var onceMail sync.Once

//单例
func GetInsMailMgr() *MailManager {
	onceMail.Do(func() {
		mailMgr = &MailManager{}
	})
	return mailMgr
}

//初始化
func (mgr *MailManager) Init(sqlInfo string, redisInfo string,redisPw string, dbIdx int,sentinelAddrs []string) {
	mgr.maxId = 0
	mgr.isChanClose = true

	dbBase, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		log.Debug(">>mail mamanger >>> %+v", err)
	}
	mgr.dbBase = dbBase
	mgr.dbBaseInfo = sqlInfo

	mgr.redisClient = &cache.RedisCacheClient{
		MaxIdle:     10,
		MaxActive:   200,
		Address:     redisInfo,
		IdleTimeout: 180 * time.Second,
		DBIdx:       dbIdx,
		Password:redisPw,
		MasterName:MASTER_NAME,
		SentinelAddr:sentinelAddrs,
	}
	if len(sentinelAddrs) <= 0 {
		mgr.redisClient.Start()
	} else {
		mgr.redisClient.StartSentinel()
	}
}

//获取邮件最大值
func (mgr *MailManager) GetMaxID() int64 {
	return mgr.maxId
}

//加载所有数据【注：只允许一台服务器调用， 其他服务器 只允许调用 init 不允许调用 此函数 】
func (mgr *MailManager) ReadAllDB() {
	mgr.maxId = 0
	if mgr.redisClient == nil {
		return
	}
	mgr.redisClient.FlushDB()
	//mgr.redisClient.DelData( REDIS_MAIL )
	//mgr.redisClient.DelData( REDIS_MAIL_USER )

	//mgr.redisClient.HSet( REDIS_MAIL, "Loaded", false) //设置为加载未完成

	count := 0
	limit := 1000
	for ; ; { //读取所有邮件内容
		mailsStr := mgr.query(MAIL_SP_SELECT_MAIL, mgr.maxId, limit)
		for _, mailStr := range mailsStr {
			mail := mgr.parseMail(mailStr)
			if mail == nil {
				continue
			}
			if len(mail.DecUid) > 0 && mail.DecUid[0] == 0 { // 表示是全服邮件, 写内存 方便验证 刚上线的玩家是否有新的全服邮件
				strId := strconv.FormatInt(mail.ID, 10)
				mgr.redisClient.SAdd(REDIS_MAIL_GLOBAL, strId)
			}
			if mgr.maxId < mail.ID {
				mgr.maxId = mail.ID
			}
			jsMail, _ := json.Marshal(mail)
			if len(jsMail) > 0 { //成功生成 json
				if mgr.redisClient != nil { //写到 redis 内
					strId := strconv.FormatInt(mail.ID, 10)
					mgr.redisClient.HSet(REDIS_MAIL+strId, REDIS_FILED_MAIL, string(jsMail))
				}
			}
			count++
		}
		if count < limit {
			break
		}
		count = 0
	}

	count = 0
	maxIdx := int32(0) //最大的索引号
	for ; ; { // 加载用户 邮件-状态
		userMailStr := mgr.query(MAIL_SP_SELECT_MAIL_USER, maxIdx, limit)
		for _, s := range userMailStr {
			if len(s) < 3 {
				continue
			}
			idx, _ := strconv.Atoi(s[0])
			strUid := s[1]
			strFlag := s[2]
			mgr.redisClient.HSet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER, strFlag)
			if maxIdx < int32(idx) { //取最大值的 uid
				maxIdx = int32(idx)
			}
			count++
		}
		if count < limit {
			break
		}
		count = 0
	}
	mgr.redisClient.HSet(REDIS_MAIL, "Loaded", true) //设置为加载完成
}

// 新邮件，插入数据库
func (mgr *MailManager) NewMail(srcUID int64, decUID string, title string, content string, gift string, mailType,condition string,expiryTime int64) *Mail {
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	var decUid []int64 //目标用户列表
	decStr := strings.Split(decUID, "-")
	for _, v := range decStr {
		uId, _ := strconv.ParseInt(v, 10, 64)
		decUid = append(decUid, uId)
	}
	if expiryTime > 0 {
		expiryTime += time.Now().Unix()
	}
	mgr.maxId += 1
	_type, _ := strconv.Atoi(mailType)
	mail := &Mail{
		ID:         mgr.maxId,
		SrcUid:     srcUID,
		DecUid:     decUid,
		Title:      title,
		Content:    content,
		Gift:       gift,
		SendTime:   time.Now().Unix(),
		ExpiryTime: expiryTime,
		Type:       int32(_type),
		Condition:condition,
	}

	if len(decUid) > 0 && decUid[0] == 0 {
		strId := strconv.FormatInt(mail.ID, 10)
		mgr.redisClient.SAdd(REDIS_MAIL_GLOBAL, strId)
	} else { // 非全服的邮件
		for _, uid := range decUid {
			strUid := strconv.FormatInt(uid, 10)
			strFlag := mgr.redisClient.HGet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER)
			strFlag += strconv.FormatInt(mail.ID, 10) + "-" + strconv.Itoa(int(MAIL_FLAG_UNREAD)) + "#"
			mgr.redisClient.HSet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER, strFlag)
		}
	}
	// 写数据库
	mgr.acceptChannel(MAIL_SQL_CMD_INSERT, mail.ID, mail.SrcUid, mail.SrcName, decUID, mail.Title, mail.Content, mail.Gift, mail.SendTime, mail.ExpiryTime, mail.Type,mail.Condition)
	//写入缓存
	strMail, _ := json.Marshal(mail)
	if len(strMail) > 0 {
		strId := strconv.FormatInt(mail.ID, 10)
		mgr.redisClient.HSet(REDIS_MAIL+ strId,REDIS_FILED_MAIL, string(strMail))
	}
	return mail
}

//刷新玩家个人邮件信息【主要用于确认是否有新群邮件】
func (mgr *MailManager) RefreshUserMail(uid int64,regTime int64) []*MailFlag {
	mgr.RLock()
	defer mgr.RUnlock()

	mgr.assert()

	var globalIds []int64 //全局邮件ID
	//ids := mgr.redisClient.HGetAll( REDIS_MAIL_GLOBAL )
	ids := mgr.redisClient.SGetAllMember(REDIS_MAIL_GLOBAL)
	for _, strId := range ids {
		id, _ := strconv.ParseInt(strId, 10, 64)
		globalIds = append(globalIds, id)
	}

	mailFlags := mgr.getUserMail(uid)
	//>>>>>>>>>>>>>>>> 在看看有没有需要新增的群邮件
	for _, id := range globalIds { //全局的邮件ID
		var found bool = false
		for _, mailFlag := range mailFlags {
			if mailFlag.ID == id {
				found = true
				break
			}
		}
		if found {
			continue
		}
		mailDB := mgr.getMail(id)
		if mailDB == nil { continue }
		if mailDB.ExpiryTime > 0 {	// 这玩家还没注册就发了这邮件，或者过期了  不给他
			if mailDB.ExpiryTime < time.Now().Unix() { continue }
			if mailDB.SendTime < regTime && time.Unix(mailDB.SendTime,0).YearDay() != time.Unix(regTime,0).YearDay() {	// 当天邮件 当天注册用户可收到
				continue
			}
		}
		mailFlags = append(mailFlags, &MailFlag{ID: id, Flag: MAIL_FLAG_UNREAD})
	}
	mgr.updateUserMail(uid, mailFlags) //更新回 redis /mysql
	return mailFlags
}

//获取玩家个人邮件数据
func (mgr *MailManager) GetUserMail(uid int64) []*MailFlag {
	mgr.RLock()
	defer mgr.RUnlock()
	mgr.assert()
	return mgr.getUserMail(uid)
}

//更新玩家个人邮件操作
func (mgr *MailManager) UpdateUserMail(uid int64, mf MailFlag) {
	mgr.RLock()
	defer mgr.RUnlock()
	mailFlags := mgr.getUserMail(uid)
	var found bool = false
	for _, mailFlag := range mailFlags {
		if mailFlag.ID == mf.ID {
			mailFlag.Flag = mf.Flag
			found = true
			break
		}
	}
	if found == false { //这里表示要新增
		mailFlags = append(mailFlags, &MailFlag{ID: mf.ID, Flag: mf.Flag})
	}

	mgr.updateUserMail(uid, mailFlags) //更新回 redis /mysql
}

//更新玩家个人邮件操作批量
func (mgr *MailManager) UpdateUserMailMult(uid int64, flags []*MailFlag) {
	mgr.RLock()
	defer mgr.RUnlock()
	mgr.assert()
	mgr.updateUserMail(uid, flags) //更新回 redis /mysql
}

//获取邮件数据
func (mgr *MailManager) GetMail(id int64) *Mail {
	mgr.RLock()
	defer mgr.RUnlock()

	return mgr.getMail(id)
}

func (mgr *MailManager) getMail(id int64) *Mail {
	mgr.assert()

	strId := strconv.FormatInt(id, 10)
	strMail := mgr.redisClient.HGet(REDIS_MAIL+strId, REDIS_FILED_MAIL)
	var mail *Mail
	err := json.Unmarshal([]byte(strMail), &mail)
	if err != nil {
		log.Debug(" sol mail get >>> %+v", err)
		return nil
	}
	return mail
}

func (mgr *MailManager) query(query string, args ...interface{}) [][]string { // 查询
	mgr.check()
	stmt, err := mgr.dbBase.Prepare(query)
	if err != nil {
		log.Debug(" query failed. error %v", err)
		return nil
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		log.Debug("query failed. %v", err)
		return nil
	}
	stmt.Close()
	defer rows.Close()
	return mgr.getRowStr(rows)
}

func (mgr *MailManager) exec(cmd string, args ...interface{}) { //执行sql 语句
	mgr.check()
	stmt, err := mgr.dbBase.Prepare(cmd)
	if err != nil {
		log.Debug(" update mail database error %v", err)
		return
	}
	_, err = stmt.Exec(args...)
	stmt.Close()
	if err != nil {
		log.Debug(" update mail database error %v", err)
	}
}

func (mgr *MailManager) getRowStr(rows *sql.Rows) [][]string { // 返回一个数据行的字符数组
	if rows == nil {
		return nil
	}
	cols, err := rows.Columns()
	rawResult := make([][]byte, len(cols))
	var result [][]string
	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}
	for rows.Next() {
		err = rows.Scan(dest...)
		colRet := make([]string, len(cols))
		for i, raw := range rawResult {
			if raw == nil {
				colRet[i] = ""
			} else {
				colRet[i] = string(raw)
			}
		}
		result = append(result, colRet)
	}
	_ = err
	return result
}

func (mgr *MailManager) parseMail(arrStr []string) *Mail { // 解析邮件字符串
	mail := &Mail{}
	if len(arrStr) < 9 {
		return nil
	}
	mail.ID, _ = strconv.ParseInt(arrStr[0], 10, 64)     //注册时间
	mail.SrcUid, _ = strconv.ParseInt(arrStr[1], 10, 64) //发送ID
	mail.SrcName = arrStr[2]                             //发送者的昵称

	mail.DecUid = []int64{} //清空下数组
	arr := strings.Split(arrStr[3], ",")
	if arr != nil || len(arr) >= 0 {
		for _, v := range arr {
			id, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				continue
			}
			mail.DecUid = append(mail.DecUid, id)
		}
	}
	mail.Title = arrStr[4]                                   //标题
	mail.Content = arrStr[5]                                 //内容
	mail.Gift = arrStr[6]                                    //附件
	mail.SendTime, _ = strconv.ParseInt(arrStr[7], 10, 64)   //时间戳
	mail.ExpiryTime, _ = strconv.ParseInt(arrStr[8], 10, 64) //过期时间
	mail.Condition = arrStr[9]	// 邮件条件
	return mail
}

// 校验，防止db为空
func (mgr *MailManager) check() {
	if mgr.dbBase == nil {
		mgr.dbBase, _ = sql.Open("mysql", mgr.dbBaseInfo)
	}
}

func (mgr *MailManager) acceptChannel(cmdType int32, args ...interface{}) { // 插入通道
	if mgr.isChanClose || mgr.channel == nil {
		mgr.openChannel()
	}

	cmd := &MailSqlCmd{
		Type: cmdType,
	}
	for _, arg := range args {
		cmd.Args = append(cmd.Args, arg)
	}
	select {
	case mgr.channel <- cmd:
		//case <-timeout: // 如果上面阻塞了，超时后会走这里
	default:
		log.Debug("sql message channel full")
	}
}

//开启通道
func (mgr *MailManager) openChannel() {
	if mgr.channel != nil {
		return
	}
	mgr.isChanClose = false
	mgr.channel = make(chan interface{}, 1000)
	go func() {
		for {
			v, ok := <-mgr.channel
			if !ok {
				mgr.channel = nil
				break
			}
			cmd, _ := v.(*MailSqlCmd)
			switch cmd.Type {
			case MAIL_SQL_CMD_INSERT:
				mgr.exec(MAIL_SP_INSERT_MAIL, cmd.Args[0], cmd.Args[1], cmd.Args[2], cmd.Args[3], cmd.Args[4], cmd.Args[5], cmd.Args[6], cmd.Args[7], cmd.Args[8], cmd.Args[9], cmd.Args[10])
				break
			case MAIL_SQL_CMD_UPDATE:
				mgr.exec(MAIL_SP_UPDATE_MAIL, cmd.Args[0], cmd.Args[1])
				break
			}
		}
		mgr.closeChannel()
	}()
}

//关闭通道
func (mgr *MailManager) closeChannel() {
	if mgr.isChanClose || mgr.channel == nil {
		return
	}
	close(mgr.channel)
	mgr.isChanClose = true
}

//获取玩家个人邮件数据
func (mgr *MailManager) getUserMail(uid int64) []*MailFlag {
	mgr.assert()
	mailFlags := []*MailFlag{}
	strUid := strconv.FormatInt(uid, 10)
	strFlag := mgr.redisClient.HGet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER)
	arrStr := strings.Split(strFlag, "#")
	for _, str := range arrStr {
		if str == "" {
			continue
		}
		mailFlag := &MailFlag{}
		mailFlag.parse(str)
		mailFlags = append(mailFlags, mailFlag) //写入数组
	}
	return mailFlags
}

//更新玩家个人邮件数据
func (mgr *MailManager) updateUserMail(uid int64, mailFlags []*MailFlag) {
	mgr.assert()
	strFlag := ""
	for _, mailFlag := range mailFlags {
		strFlag += mailFlag.form()
	}
	strUid := strconv.FormatInt(uid, 10)
	oldFlag := mgr.redisClient.HGet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER)
	mgr.redisClient.HSet(REDIS_MAIL_USER+strUid, REDIS_FILED_MAIL_USER, strFlag)
	if oldFlag != strFlag {	// 做下判断，没改动就不更新数据库了
		mgr.acceptChannel(MAIL_SQL_CMD_UPDATE, uid, strFlag) // 写数据库
	}
}

//校验
func (mgr *MailManager) assert() {
	if mgr.redisClient == nil {
		panic("mail manager ... redis is nil ")
	}

	if mgr.redisClient.HGet(REDIS_MAIL, "Loaded") != "1" {
		panic("mail manager ... loaded undone")
	}
}