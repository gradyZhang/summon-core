package sol

import (
	"sync"
	"time"
	"github.com/gradyZhang/summon-core/aliens/common/cache"
	"strconv"
	"github.com/gradyZhang/summon-core/aliens/log"
	"strings"
	"encoding/json"
	"bytes"
	"compress/zlib"
	"io"
)

const (
	REDIS_USER = "rU"//redis 的key
	REDIS_USER_DATA = "rfUData" //用户数据
	REDIS_USER_LOAD = "rfULoad" //用户数据加载标识
	REDIS_USER_SHARE = "rfUShare"	// 用户分享标识
	REDIS_SQL = "rSql" // SQL执行命令

	REDIS_ACC = "rAcc" //帐号
	REDIS_ACC_UID = "rfAccUid" //帐号 和 uid 的映射表

	REDIS_USER_ONLINE = "rOnline" //玩家是否在线的key

	REDIS_NAME = "rName" //昵称
	REDIS_NAME_UID = "rfNameUid" //用户ID
	MASTER_NAME = "mymaster"	// 主机名
)

type UserManager struct{
	sync.Mutex
	maxUid       int64            //最大的用户ID
	redisClient *cache.RedisCacheClient
	cacheUser []*User	// 缓存个
}

var userMgr *UserManager
var onceUser sync.Once

//单例
func GetInsUserMgr() *UserManager {
	onceUser.Do(func() {
		userMgr = &UserManager{}
	})
	return userMgr
}

//初始化
func (mgr* UserManager)Init( redisInfo string,redisPw string, dbIdx int ,sentinelAddrs []string){
	mgr.Lock()
	defer mgr.Unlock()
	mgr.redisClient = &cache.RedisCacheClient{
		MaxIdle:     10,
		MaxActive:   200,
		Address:     redisInfo,
		IdleTimeout: 180 * time.Second,
		DBIdx:dbIdx,
		Password:redisPw,
		MasterName:MASTER_NAME,
		SentinelAddr:sentinelAddrs,
	}
	if len(sentinelAddrs) <= 0 {
		mgr.redisClient.Start()
	} else {
		mgr.redisClient.StartSentinel()
	}
	mgr.cacheUser = []*User{}
}

//刷新Redis 【注：只允许 一台服务器进行操作】
func (mgr *UserManager)FlushDB(){
	if mgr.redisClient == nil{
		return
	}
	mgr.redisClient.FlushDB()
}

//插入用户数据
func (mgr *UserManager)InserUser( arrStr []string )*User{
	mgr.Lock()
	defer mgr.Unlock()
	user := &User{}
	user.Parse( arrStr )
	if mgr.maxUid < user.Uid { //记录最大的 uid
		mgr.maxUid = user.Uid
	}
	//mgr.redisSave( user )//存入到redis
	mgr.cacheUser = append(mgr.cacheUser,user)
	return user
}

func (mgr *UserManager) SetMaxUID(id int64) {
	mgr.maxUid = id
}

//创建用户数据
func (mgr *UserManager)CreateUser( account string )*User{
	mgr.Lock()
	defer mgr.Unlock()
	if mgr.redisClient == nil{
		return nil
	}
	if mgr.maxUid == 0{
		mgr.maxUid = 100000001
	}else{
		mgr.maxUid = mgr.maxUid + 1
	}
	user := &User{}
	user.New( mgr.maxUid, account )
	return user
}

//设置加载标识
func (mgr *UserManager)SetLoaded( loaded bool ){
	mgr.Lock()
	defer mgr.Unlock()
	if mgr.redisClient == nil{
		return
	}
	if loaded == true{//表示加载完毕
		mgr.redisClient.HSet( REDIS_USER, REDIS_USER_LOAD, "1")
	}else{
		mgr.redisClient.HSet( REDIS_USER, REDIS_USER_LOAD, "0")
	}
}

//是否已经加载完毕
func (mgr *UserManager)IsLoaded()bool{
	mgr.Lock()
	defer mgr.Unlock()
	if mgr.redisClient == nil{
		return false
	}
	str := mgr.redisClient.HGet( REDIS_USER, REDIS_USER_LOAD )
	if str == ""{
		return false
	}
	ret, _ := strconv.Atoi( str )
	if ret != 0{
		return true
	}
	return false
}

//设置用户的在线情况
func (mgr *UserManager)SetOnline( uid int64, srvId int32 ){
	mgr.Lock()
	defer mgr.Unlock()
	if mgr.redisClient == nil{
		return
	}
	strUid := strconv.FormatInt( uid, 10 )
	strSrvId := strconv.Itoa( int(srvId) )
	if srvId == 0 {	// 标识下线了
		mgr.redisClient.HDel(REDIS_USER_ONLINE, strUid)
	} else {
		mgr.redisClient.HSet(REDIS_USER_ONLINE, strUid, strSrvId )
	}
}

// 获取当前在线总人数
func (mgr *UserManager) GetOnlineNum() int32 {
	mgr.Lock()
	defer mgr.Unlock()

	if mgr.redisClient == nil {
		return 0
	}
	ret := mgr.redisClient.HLen(REDIS_USER_ONLINE)
	return int32(ret)
}

//设置某个服务器的用户 全部下线
func (mgr *UserManager)SetOffline( srvId int32 ){
	mgr.Lock()
	defer mgr.Unlock()
	if mgr.redisClient == nil{
		return
	}
	strs := mgr.redisClient.HGetAll( REDIS_USER_ONLINE )
	strUid := ""
	for i, v := range strs{
		if i%2 == 0{
			strUid = v
			continue
		}
		tmpSrvId, _ := strconv.Atoi( v )
		if int32(tmpSrvId) == srvId{//只有在同一个服务器的情况下去出了力
			mgr.redisClient.HDel( REDIS_USER_ONLINE ,strUid)
		}
	}
}

//用户是否在线
func (mgr *UserManager)IsOnline( uid int64 )int32{
	mgr.Lock()
	defer mgr.Unlock()
	if mgr.redisClient == nil{
		return 0
	}
	strUid := strconv.FormatInt( uid, 10 )
	strSrvId := mgr.redisClient.HGet(REDIS_USER_ONLINE,strUid)
	if strSrvId == ""{
		return 0
	}
	strId, _ := strconv.Atoi( strSrvId )
	return int32(strId)
}

//设置帐号关联映射
func (mgr *UserManager)SetAccUid( account string, uid int64 ){
	//mgr.Lock()
	//defer mgr.Unlock()
	//if mgr.redisClient == nil{
	//	return
	//}
	//strUid := strconv.FormatInt( uid, 10 )
	//mgr.redisClient.HSet(REDIS_ACC+account, REDIS_ACC_UID, strUid)
}

// 获取用户id -- BY account
func (mgr *UserManager) GetUidByAcc( account string) int64{
	mgr.Lock()
	defer mgr.Unlock()
	if mgr.redisClient == nil{
		return 0
	}
	strUid := mgr.redisClient.HGet( REDIS_ACC+account, REDIS_ACC_UID )
	if strUid == ""{
		return 0
	}
	uid, _ := strconv.ParseInt( strUid, 10, 64)
	return uid
}

// 获取用户id
func (mgr *UserManager) GetUserID(account string) string {
	mgr.Lock()
	defer mgr.Unlock()

	if mgr.redisClient == nil{
		return ""
	}
	strUid := mgr.redisClient.HGet( REDIS_ACC+account, REDIS_ACC_UID )
	return strUid
}

// 获取用户数据 -- BY account
func (mgr *UserManager)GetUserByAcc( account string )*User{
	mgr.Lock()
	defer mgr.Unlock()

	if mgr.redisClient == nil{
		return nil
	}
	strUid := mgr.redisClient.HGet( REDIS_ACC+account, REDIS_ACC_UID )
	if strUid == ""{
		return nil
	}
	return mgr.redisGet( strUid )
}

// 获取用户数据 -- BY uid
func (mgr *UserManager)GetUserByUid( uid int64 )*User{
	mgr.Lock()
	defer mgr.Unlock()

	return mgr.getUserByUid(uid)
}

func (mgr *UserManager) getUserByUid(uid int64) *User {
	if mgr.redisClient == nil {
		return nil
	}
	strUid := strconv.FormatInt( uid, 10 )
	return mgr.redisGet( strUid )
}


// 获取用户数据 -- BY 昵称
func (mgr *UserManager)GetUserByName( name string )[]*User{
	mgr.Lock()
	defer mgr.Unlock()
	var users []*User
	userMap := make(map[int64]*User)	// 确保不会有重复的user
	if mgr.redisClient == nil{
		return users
	}
	// 模糊查找
	values := mgr.redisClient.Scan(REDIS_NAME +name+"*",10)	// 获取key
	values = append(values,REDIS_NAME + name)	// 确保精确的名字能找到

	for _,s := range values {
		if len(userMap) >= 10 {
			break
		}
		strUids := mgr.redisClient.HGet(s, REDIS_NAME_UID )//获取用户ID
		if strUids != "" {
			arrUid := strings.Split( strUids, "#")
			for _,s := range arrUid {
				if len(userMap) >= 10 {
					break
				}
				if s == "" {
					continue
				}
				user := mgr.redisGet( s )
				if user == nil{
					continue
				}
				userMap[user.Uid] = user
			}
		}
	}
	for _,v := range userMap {
		users = append(users,v)
	}
	return users
}

// 保存用户数据
func (mgr *UserManager)SaveUser( user *User ){
	mgr.Lock()
	defer mgr.Unlock()
	mgr.redisSave( user )
}

// 保存用户数据
func (mgr *UserManager)redisSave( user *User ){
	if mgr.redisClient == nil{
		return
	}
	if user == nil {
		return
	}
	redisCmd := [][]string{}
	strUid := strconv.FormatInt( user.Uid, 10 )
	//mgr.redisClient.HSet(REDIS_USER+strUid,REDIS_USER_SHARE,user.ShareDB)	// 分享数据
	redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_USER+strUid,REDIS_USER_SHARE,user.ShareDB})
	user.ShareDB = ""	// 置空了，让数据少点
	strUser, err := json.Marshal(user)
	if err != nil {
		log.Info("json err : %v", err)
		return
	}

//把js 数据压缩一轮进入 redis
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write([]byte(strUser))
	w.Close()
	redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_USER + strUid, REDIS_USER_DATA, in.String()})
	redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_ACC + user.Account, REDIS_ACC_UID, string(strUid)})
	if len(user.BindAcc) > 2 {	// 有绑定的账号
		redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_ACC + user.BindAcc, REDIS_ACC_UID, string(strUid)})
	}
	//mgr.redisClient.HSet( REDIS_USER + strUid, REDIS_USER_DATA, in.String() )//保存用户数据
	//mgr.redisClient.HSet( REDIS_ACC + user.Account, REDIS_ACC_UID, string(strUid) )// 存个acc - uid 映射
	ret := mgr.redisClient.HGet( REDIS_NAME + user.Name, REDIS_NAME_UID )//由于昵称是允许重复的, 所以存 redis 的时候需要注意下...
	if ret == ""{
		//mgr.redisClient.HSet( REDIS_NAME + user.Name, REDIS_NAME_UID, strUid+"#" ) //存个 昵称 - uid 映射
		redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_NAME + user.Name, REDIS_NAME_UID, strUid+"#"})
	}else{
		ret += strUid + "#"
		//mgr.redisClient.HSet( REDIS_NAME + user.Name, REDIS_NAME_UID, ret ) //存个 昵称 - uid 映射
		redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_NAME + user.Name, REDIS_NAME_UID, ret })
	}
	mgr.redisClient.Send(redisCmd)
}

//获取用户数据
func (mgr *UserManager)redisGet( strUid string )*User{
	strUser := mgr.redisClient.HGet( REDIS_USER + strUid, REDIS_USER_DATA )//保存用户数据
	if strUser == ""{
		return nil
	}
//获取到数据以后先进行一轮解压
	var tmp bytes.Buffer
	tmp.Write( []byte(strUser) )
	var out bytes.Buffer
	r, e := zlib.NewReader( &tmp )
	if e != nil{
		return nil
	}
	io.Copy(&out, r)

	var user *User
	err := json.Unmarshal([]byte(out.String()), &user)
	if err != nil {//哎哟出错了哟
		log.Info("redis get user err >> %+v", err.Error() )
	}
	user.ShareDB = mgr.redisClient.HGet(REDIS_USER+strUid,REDIS_USER_SHARE)	// 用户分享的数据
	return user
}

// 批量保存
func (mgr *UserManager) redisSaveBatch(users []*User) {
	if mgr.redisClient == nil{
		return
	}
	c := mgr.redisClient.GetConn()
	nameUid := make(map[string][]int64)
	for _,user := range users {
		strUid := strconv.FormatInt( user.Uid, 10 )
		//c.Send("HSET",REDIS_USER+strUid,REDIS_USER_SHARE,user.ShareDB)
		shareStr := user.ShareDB
		user.ShareDB = ""	// 让数据少点
		strUser, err := json.Marshal(user)
		if err != nil {
			log.Info("json err : %v", err)
			return
		}
		if _,ok := nameUid[user.Name];!ok {
			nameUid[user.Name] = []int64{}
		}
		nameUid[user.Name] = append(nameUid[user.Name],user.Uid)
		//把js 数据压缩一轮进入 redis
		var w *zlib.Writer
		var in bytes.Buffer
		w = zlib.NewWriter(&in)
		w.Write([]byte(strUser))
		w.Close()
		c.Send("HMSET",REDIS_USER + strUid, REDIS_USER_DATA, in.String(),REDIS_USER_SHARE,shareStr)//保存用户数据
		c.Send("HSET",REDIS_ACC + user.Account, REDIS_ACC_UID, string(strUid))// 存个acc - uid 映射
		if len(user.BindAcc) > 2 {	// 有绑定的账号
			c.Send("HSET",REDIS_ACC + user.BindAcc, REDIS_ACC_UID, string(strUid))// 存个acc - uid 映射
		}
		//ret := ""//mgr.redisClient.HGet( REDIS_NAME + user.Name, REDIS_NAME_UID )//由于昵称是允许重复的, 所以存 redis 的时候需要注意下...
		//if ret == ""{
		//	c.Send( "HSET",REDIS_NAME + user.Name, REDIS_NAME_UID, strUid+"#" ) //存个 昵称 - uid 映射
		//}else{
		//	ret += strUid + "#"
		//	c.Send( "HSET", REDIS_NAME + user.Name, REDIS_NAME_UID, ret ) //存个 昵称 - uid 映射
		//}
		//c.Send("HSET",REDIS_USER_ONLINE, strUid, 0x00 )
		c.Send("SELECT",1)
		c.Send("HMSET",REDIS_FRIERND+strUid, REDIS_FRIEND_DATA, user.SocialFriends, REDIS_FRIEND_APPLY,user.SocialApplies,REDIS_FRIEND_DATA_INVITE,user.SocialInvitee)//好友数据
		//c.Send("HSET",REDIS_FRIERND+strUid, REDIS_FRIEND_APPLY,user.SocialApplies)//申请数据
		c.Send("SELECT",3)
	}
	for name,uids := range nameUid {
		str := ""
		for _,id := range uids {
			str += strconv.FormatInt( id, 10 )
			str += "#"
		}
		c.Send("HSET", REDIS_NAME+name, REDIS_NAME_UID, str) //存个 昵称 - uid 映射
	}
	c.Flush()
	c.Receive()
	c.Close()
}

func (mgr *UserManager) ReadUserEnd() {
	mgr.Lock()
	defer mgr.Unlock()

	if len(mgr.cacheUser) > 0 {
		mgr.redisSaveBatch(mgr.cacheUser)
		mgr.cacheUser = []*User{}
	}
}

// 记录其他用户点击分享链接
func (mgr *UserManager) RecordClickShare(timestamp int64,uid int64,clickUid int64) {
	mgr.Lock()
	defer mgr.Unlock()

	if mgr.redisClient == nil {
		return
	}
	strUid := strconv.FormatInt( uid, 10 )
	str := mgr.redisClient.HGet(REDIS_USER+strUid,REDIS_USER_SHARE)
	var share *UserShare
	err := json.Unmarshal([]byte(str),&share)
	if err != nil {
		log.Debug("== err %v",err)
		return
	}
	if share.DailyLinkCount > 20 {	// 超过上限了
		return
	}
	uIDs,ok := share.Shares[timestamp]	// 没有这个分享链接
	if !ok {
		return
	}
	for _,id := range uIDs {
		if id == clickUid {	// 已经点过了
			return
		}
	}
	uIDs = append(uIDs,clickUid)	// 加上这个用户标识
	share.Shares[timestamp] = uIDs
	share.DailyLinkCount += 1
	bytes,_ := json.Marshal(share)
	mgr.redisClient.HSet(REDIS_USER+strUid,REDIS_USER_SHARE,string(bytes))	// 写回去
}

// 记录SQL执行命令
func (mgr *UserManager) RecordSqlCMD(uid int64,cmd string,cmdObj string) {
	mgr.Lock()
	defer mgr.Unlock()

	if mgr.redisClient == nil {
		return
	}
	idStr := strconv.FormatInt( uid, 10 )
	mgr.redisClient.HSet(REDIS_SQL+idStr,cmd,cmdObj)
}

// 删除key
func (mgr *UserManager) RemoveKey(key string,feild string,uid int64) {
	mgr.Lock()
	defer mgr.Unlock()
	if mgr.redisClient == nil {
		return
	}
	idStr := strconv.FormatInt( uid, 10 )
	mgr.redisClient.HDel(REDIS_SQL+idStr,feild)
}

// 记录玩家支付商品
func (mgr *UserManager) RecordUserPayShop(uid int64,shopId int32) {
	mgr.Lock()
	defer mgr.Unlock()

	userDB := mgr.getUserByUid(uid)
	if userDB != nil {
		userDB.PayShopId = shopId
		mgr.redisSave( userDB )
	}
}