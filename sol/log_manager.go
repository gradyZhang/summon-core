package sol

import (
"sync"
"os"
"time"
"def"
"aliens/common/tools"
"strings"
"io"
"encoding/json"
"bufio"
)

// 日志类型
const (
	LogTypeNull     int32 = iota
	LogTypeLogin     //登录日志
	LogTypeGold      //金币日志
	LogTypeDiamond   // 钻石
	LogTypeFeedback  // 客户端反馈
	LogTypeSQLCmd    // 更新异常
	LogTypeChapter   // 章节通关
	LogTypeUserAct   // 活跃玩家
	LogTypeCharge    // 玩家充值
	LogTypeRefine    // 熔炼记录
	LogTypeArena	// 竞技场日志
	LogTypeGMOp		// GM操作日志
	LogTypeLogout	// 登出日志
	LogTypeCard		// 怪物碎片日志
	LogTypeShopBuy	// 商店购买日志
	LogTypeArenaShop// 竞技场商店
	LogTypeClientDot // 客户端打点
)

type LogDB struct {
	LogName string
	Data    []byte
}

// 日志管理
type LogManager struct {
	sync.RWMutex
	channel chan interface{}
}

var logMgr *LogManager
var onceLog sync.Once

func GetInsLogMgr() *LogManager {
	onceLog.Do(func() {
		logMgr = &LogManager{}
		logMgr.init()
	})
	return logMgr
}

func (l *LogManager) init() {
	_, err := os.Stat("log/"); //判断日志文件夹是否存在
	if err != nil { //日志文件夹不存在哦,那就创建一个新的日志 文件夹
		os.MkdirAll("./log", os.ModePerm)
	}
	l.openChan()
}

func (l *LogManager) openChan() {
	if l.channel == nil {
		l.channel = make(chan interface{}, 10)
	}
	go func() {
		for {
			log, ok := <-l.channel
			if !ok {
				break
			}
			logDB := log.(LogDB)
			date := time.Now().Format("20060102")
			fileName := "log/" + logDB.LogName + "-" + date + ".json"
			file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
			if err != nil {
				continue
			}
			n, _ := file.Seek(0, os.SEEK_END)
			js := append(logDB.Data, '\n')
			file.WriteAt(js, n)
			file.Close()
		}
		close(l.channel)
	}()
}

// 记录登录日志
func (l *LogManager) Write(opType int32, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	t := time.Now().Format("2006-01-02 15:04:05")
	log := LogDB{}
	switch opType {
	case LogTypeDiamond:
		log.LogName = "diamond"
		data := &def.LogDiamond{Time: t, UserID: args[0].(int64), CurrencyTotal: args[2].(int32), Value: args[3].(int32), OpType: args[4].(int32)}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeGold:
		log.LogName = "gold"
		data := &def.LogGold{Time: t, UserID: args[0].(int64), CurrencyTotal: args[2].(int32), Value: args[3].(int32), OpType: args[4].(int32)}
		//strOp := strconv.Itoa(int(args[4].(int32)))
		//switch  args[4].(int32) {
		//case def.LogOpReward:
		//	strOp = "reward"
		//	break
		//case def.LogOpHeroCreate:
		//	strOp = "hero create"
		//	break
		//case def.LogOpHeroLvUp:
		//	strOp = "hero lvl up"
		//	break
		//case def.LogOpChest: //宝箱
		//	strOp = "chest open"
		//	break
		//case def.LogOpChestBuy: //商店宝箱购买
		//	strOp = "chest buy"
		//	break
		//}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeLogin:
		log.LogName = "login"
		data := &def.LogLogin{Time: t, UserID: args[0].(int64), //Name: args[1].(string), 
			Lv: args[2].(int32), Gold: args[3].(int32), Diamond: args[4].(int32), RegTime: args[5].(int64),Vit:args[6].(int32),Platform:args[7].(int32)}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeLogout:
		log.LogName = "logout"
		data := &def.LogLogout{Time: t, UserID: args[0].(int64), //Name: args[1].(string), 
			Lv: args[2].(int32), Gold: args[3].(int32), Diamond: args[4].(int32), RegTime: args[5].(int64),Vit:args[6].(int32),LoginTime:args[7].(int64)}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeChapter:
		log.LogName = "chapter"
		data := &def.LogChapter{
			Time:      t,
			UserID:    args[0].(int64),
			ChapterID: args[1].(int32),
			UseReel:   args[2].(bool),
			RegTime:   args[3].(int64),
		}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeFeedback:
		log.LogName = "feedback"
		data := &def.LogClientFeedback{
			Time:   t,
			UserID: args[0].(int64),
			Type:   args[1].(int32),
			Msg:    args[2].(string),
			Count:  args[3].(int32),
			//args[4].(time.Time)
		}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeSQLCmd:
		log.LogName = "sql_failed"
		data := &def.LogSqlCmd{
			Time:   t,
			UserID: args[0].(int64),
			SqlCmd: args[1].(int32),
			Args:   args[2].(string),
		}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeCharge:
		log.LogName = "charge"
		data := &def.LogCharge{
			Time:     t,
			UserID:   args[0].(int64),
			Amount:   args[1].(float32),
			ShopID:   args[2].(int32),
			Platform: args[3].(int32),
			RegTime:  args[4].(int64),
			Lv :args[5].(int32),
			MaxChapter :args[6].(int32),
			Country:args[7].(int32),
			ChargeSum:args[8].(int32),
			Formation:args[9].(string),
			Equip:args[10].(string),
		}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeUserAct:
		log.LogName = "user_active"
		data := &def.LogAct{Time: t, Act: args[0].(int)}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeRefine:
		log.LogName = "refine"
		data := &def.LogRefine{
			Time:      t,
			UserID:    args[0].(int64),
			RefineGet: args[1].(string),
			UseCard:   args[2].(string),
			//Cost:      args[3].(int32),
		}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeArena:
		log.LogName = "arena"
		data := &def.LogArena{
			Time:t,
			BeginTime:args[0].(int64),
			EndTime:args[1].(int64),
			RoundRet:args[2].([]int32),
			Victor:args[3].(int64),
			PlayerRed:args[4].(*def.LogArenaPlayer),
			PlayerBlue:args[5].(*def.LogArenaPlayer),
		}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeGMOp:
		log.LogName = "gm_op"
		data := &def.LogGMOp{
			Time:t,
			Addr:args[0].(string),
			User:args[1].(string),
			OpType:args[2].(string),
			Msg:args[3].(string),
		}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeCard:
		log.LogName = "card"
		data := &def.LogCard{
			Time:t,
			UserID:    args[0].(int64),
			FamilyID:args[1].(int32),
			OpType:args[2].(int32),
			Value:args[3].(int32),
			Total:args[4].(int32),
		}
		log.Data = []byte(tools.MarshalObj(data))
	case LogTypeShopBuy:
		log.LogName = "shop_buy"
		data := &def.LogShopBuy{
			Time:t,
			UserID:args[0].(int64),
			ShopID:args[1].(int32),
			Price:args[2].(float32),
			CurType:args[3].(int32),
		}
		log.Data = []byte(tools.MarshalObj(data))
	default:
		return
	}
	l.channel <- log
}

func (l *LogManager) AddRegisterLog(uid int64,account string,channelID int32,platform int32) {
	l.Lock()
	defer l.Unlock()
	t := time.Now().Format("2006-01-02 15:04:05")
	data := &def.LogRegUser{ Time:t,UserID:uid,Account:account,ChannelID:channelID,Platform:platform}
	log := LogDB{LogName:"reg_user",Data:[]byte(tools.MarshalObj(data))}
	l.channel <- log
}

func (l *LogManager) AddArenaChestLog(uid int64,chestID int32,star int32) {
	l.Lock()
	defer l.Unlock()
	t := time.Now().Format("2006-01-02 15:04:05")
	data := &def.LogArenaChest{ Time:t,UserID:uid,ChestID:chestID,Star:star}
	log := LogDB{LogName:"arena_chest",Data:[]byte(tools.MarshalObj(data))}
	l.channel <- log
}

func (l *LogManager) AddEquipLog(uid int64,opType,subType int32,equipIDs,equipBaseIDs []int32,upEquipID,upEquipBaseID int32) {
	l.Lock()
	defer l.Unlock()
	t := time.Now().Format("2006-01-02 15:04:05")
	data := &def.LogEquip{ Time:t,UserID:uid,Type:opType,SubType:subType,EquipID:equipIDs,EquipBaseID:equipBaseIDs,UpEquipBaseID:upEquipBaseID,UpEquipID:upEquipID,Count:int32(len(equipIDs))}
	log := LogDB{LogName:"equip",Data:[]byte(tools.MarshalObj(data))}
	l.channel <- log
}

func (l *LogManager) AddReelLog(uid int64,opType int32,ids,nums []int32,stageID int32) {
	l.Lock()
	defer l.Unlock()
	t := time.Now().Format("2006-01-02 15:04:05")
	data := &def.LogReel{ Time:t,UserID:uid,Type:opType,ReelIDs:ids,Num:nums,StageID:stageID}
	log := LogDB{LogName:"reel",Data:[]byte(tools.MarshalObj(data))}
	l.channel <- log
}

func (l *LogManager) AddHitBossLog(uid, damage,bossTime int64,num int32) {
	l.Lock()
	defer l.Unlock()
	t := time.Now().Format("2006-01-02 15:04:05")
	data := &def.LogHitBoss{ Time:t,UserID:uid,Damage:damage,BossTime:time.Unix(bossTime,0).Format("060102150405"),HitNum:num}
	log := LogDB{LogName:"hit_boss",Data:[]byte(tools.MarshalObj(data))}
	l.channel <- log
}

func (l *LogManager) AddShareLog(uid int64,shareT string) {
	l.Lock()
	defer l.Unlock()
	t := time.Now().Format("2006-01-02 15:04:05")
	data := &def.LogShare{ Time:t,UserID:uid,ShareType:shareT}
	log := LogDB{LogName:"share",Data:[]byte(tools.MarshalObj(data))}
	l.channel <- log
}
//
func (l *LogManager) AddArenaShopLog(uid int64,goods string,price, times,score int32) {
	l.Lock()
	defer l.Unlock()
	t := time.Now().Format("2006-01-02 15:04:05")
	data := &def.LogArenaShop{ Time:t,UserID:uid,Goods:goods,Price:price,Times:times,Score:score}
	log := LogDB{LogName:"a_shop",Data:[]byte(tools.MarshalObj(data))}
	l.channel <- log
}
// 客户端打点
func (l *LogManager) AddClientDot(uid int64,key,arg1,arg2 int32) {
	l.Lock()
	defer l.Unlock()
	t := time.Now().Format("2006-01-02 15:04:05")
	data := &def.LogClientDot{ Time:t,UserID:uid,Key:key,Arg1:arg1,Arg2:arg2}
	log := LogDB{LogName:"client_dot",Data:[]byte(tools.MarshalObj(data))}
	l.channel <- log
}

func (l *LogManager) RecordPlayerExp(m map[int64]int32) {
	l.Lock()
	defer l.Unlock()
	t := time.Now().Format("2006-01-02 15:04:05")

	fileName := "log/exp.json"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	js := ""
	for uid,exp := range m {
		if uid == 0 || exp == 0 { continue }
		data := &def.LogExp{ Time:t,UserID:uid,Exp:exp}
		js += (tools.MarshalObj(data) + "\n")
	}
	file.Write([]byte(js))
	file.Close()
}

func (l *LogManager) ReadPlayerExp() map[int64]int32 {
	l.Lock()
	defer l.Unlock()
	fileName := "log/exp.json"
	f,err := os.Open(fileName)
	if err != nil {
		return nil
	}
	result := make(map[int64]int32)
	reader := bufio.NewReader(f)
	for {
		line,err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		l := &def.LogExp{}
		json.Unmarshal([]byte(line),&l)
		if l.UserID != 0 && l.Exp != 0 {
			result[l.UserID] = l.Exp
		}
		if err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	f.Close()
	return result
}