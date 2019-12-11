package sol

import (
	"sync"
	"strconv"
	"github.com/gradyZhang/summon-core/aliens/common/cache"
	"time"
	"encoding/json"
	"github.com/gradyZhang/summon-core/aliens/log"
)

const (
	REDIS_FRIERND = "rF"//好友 redis 的key
	REDIS_FRIEND_DATA = "rfData" //好友
	REDIS_FRIEND_APPLY = "rfApply" //申请数据
	REDIS_FRIEND_GIFT_SEND = "rfGiftSend" //送礼品的次数
	REDIS_FRIEND_GIFT_RECV = "rfGiftRecv" //收礼品的礼品
	REDIS_FRIEND_GIFT_SEND_UID = "rfSendUid" //送过礼品的用户
	REDIS_FRIEND_DATA_CHAN = "rfDataC"	// 渠道好友数据

	REDIS_FRIEND_DATA_INVITE = "rFDataI"	// 邀请码好友
)
const(
	Gift_State_None int32 = iota	// 不可领
	Gift_State_Recv 	// 可领状态
)

type Friend struct{
	Uid int64
	//GiftSend bool //是否赠送了礼品
	GiftState int32 //是否存在礼品( 0x00 无, n:存在n个)
	FriValue int32 	// 好友值
}

type FriendManager struct {
	sync.Mutex
	redisClient *cache.RedisCacheClient //redis 句柄
}

var friendMgr *FriendManager
var onceFriend sync.Once
func GetInsFriendMgr() *FriendManager {
	onceFriend.Do(func() {
		friendMgr = &FriendManager{}
	})
	return friendMgr
}


func ( mgr *FriendManager )Init( redisInfo string,redisPw string, dbIdx int,sentinelAddrs []string ){
	mgr.redisClient = &cache.RedisCacheClient{
		MaxIdle:     10,
		MaxActive:   200,
		Address:     redisInfo,
		IdleTimeout: 180 * time.Second,
		DBIdx: dbIdx,
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

func ( mgr *FriendManager )FlushDB(){//清空数据
	if mgr.redisClient == nil{
		return
	}
	mgr.redisClient.FlushDB()
}

func ( mgr *FriendManager )GetData( uid int64 )(int32, int32, []*Friend, []int64, []int64, []*Friend){//获取数据
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	strUid := strconv.FormatInt(uid,10)

	strSendNum := mgr.redisClient.HGet( REDIS_FRIERND+strUid, REDIS_FRIEND_GIFT_SEND ) //礼品赠送次数
	strRecvNum := mgr.redisClient.HGet( REDIS_FRIERND+strUid, REDIS_FRIEND_GIFT_RECV ) //礼品领取次数
	sendUidStr := mgr.redisClient.HGet( REDIS_FRIERND+strUid, REDIS_FRIEND_GIFT_SEND_UID)// 已经送过礼品的好友
	strFriend := mgr.redisClient.HGet( REDIS_FRIERND+strUid, REDIS_FRIEND_DATA )
	strApply := mgr.redisClient.HGet( REDIS_FRIERND+strUid, REDIS_FRIEND_APPLY )
	strInvitee := mgr.redisClient.HGet( REDIS_FRIERND+strUid, REDIS_FRIEND_DATA_INVITE)	// 邀请码好友

	var friends,invitee []*Friend

	if strInvitee != "" {
		err := json.Unmarshal( []byte(strInvitee), &invitee )
		if err != nil{
			log.Debug("uid (%v) parse invitee err >> %+v ", uid, err )
		}
	}

	if strFriend != ""{
		err := json.Unmarshal( []byte(strFriend), &friends )
		if err != nil{
			log.Debug("uid (%v) parse friend err >> %+v ", uid, err )
		}
	}
	var applys []int64

	if strApply != ""{
		err := json.Unmarshal( []byte(strApply), &applys )
		if err != nil{
			log.Debug("uid (%v) parse apply err >> %+v ", uid, err )
		}
	}
	var sendUids []int64
	if sendUidStr != "" {
		err := json.Unmarshal([]byte(sendUidStr), &sendUids)
		if err != nil {
			log.Debug("uid (%v) parse send gift uid err >> %+v ", uid, err)
		}
	}

	var sendNum int
	if strSendNum == ""{
		sendNum = 0
	}else{
		sendNum, _ = strconv.Atoi( strSendNum )
	}

	var recvNum int
	if strRecvNum == ""{
		recvNum = 0
	}else{
		recvNum, _ = strconv.Atoi( strRecvNum )
	}
	return int32(sendNum), int32(recvNum), friends, applys,sendUids,invitee
}

func ( mgr *FriendManager )WriteData( uid int64, sendNum,recvNum int32, friends,invitee []*Friend, applys,sendUid []int64)(string,string, string){//写入缓存数据
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	strUid := strconv.FormatInt(uid,10)

	strSendNum := strconv.Itoa( int(sendNum) ) //赠送礼品的次数
	strRecvNum := strconv.Itoa( int(recvNum) ) //领取礼品的次数
	strFriend, _ := json.Marshal( friends )
	strApply, _ := json.Marshal( applys )
	strInvitee,_ := json.Marshal(invitee)	// 邀请码好友
	sendUidStr,_ := json.Marshal(sendUid)	// 送过礼物的用户
	//cFriends,_ := json.Marshal(channelFriends)	// 渠道好友数据
	redisCmd := [][]string{}
	//redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_FRIERND+strUid, REDIS_FRIEND_DATA_CHAN,string(cFriends)})
	redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_FRIERND+strUid, REDIS_FRIEND_GIFT_SEND_UID,string(sendUidStr)})// 送过礼品的用户ID
	redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_FRIERND+strUid, REDIS_FRIEND_GIFT_SEND, strSendNum }) //赠送次数
	redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_FRIERND+strUid, REDIS_FRIEND_GIFT_RECV, strRecvNum }) //领取次数
	redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_FRIERND+strUid, REDIS_FRIEND_DATA, string(strFriend) }) //好友数据
	redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_FRIERND+strUid, REDIS_FRIEND_APPLY, string(strApply) }) //申请数据
	redisCmd = append(redisCmd,[]string{cache.OP_H_SET,REDIS_FRIERND+strUid, REDIS_FRIEND_DATA_INVITE,string(strInvitee)})// 邀请码好友
	mgr.redisClient.Send(redisCmd)
	return string(strFriend), string(strApply),string(strInvitee)
}

func (mgr *FriendManager)InsertData( uid int64, strFriend string, strApply string ){//插入数据
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	strUid := strconv.FormatInt(uid,10)
	mgr.redisClient.HSet(REDIS_FRIERND+strUid, REDIS_FRIEND_DATA, string(strFriend) ) //好友数据
	mgr.redisClient.HSet(REDIS_FRIERND+strUid, REDIS_FRIEND_APPLY, string(strApply) ) //申请数据
	mgr.redisClient.HSet(REDIS_FRIERND+strUid, REDIS_FRIEND_DATA_INVITE, "" ) //邀请码好友数据
}
// 插入 邀请码好友
func (mgr *FriendManager)AddInvite( srcUid int64, decUid int64 ,max int) {
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()

	strUid := strconv.FormatInt(srcUid, 10)
	strFriend := mgr.redisClient.HGet(REDIS_FRIERND+strUid, REDIS_FRIEND_DATA_INVITE)

	var friends []*Friend

	err := json.Unmarshal( []byte(strFriend), &friends)
	if err != nil{
		log.Debug("uid (%v) add friend err >> %+v ", srcUid, err )
		return
	}
	if len(friends) >= max {	// 好友数超上限了
		return
	}
	for _, f := range friends{
		if f.Uid == decUid{//已经是好友了 所以就不用反复添加了
			return
		}
	}
	friend := &Friend{
		Uid:decUid,
	}

	friends = append( friends, friend )
	byteFriend, _ := json.Marshal( friends )
	mgr.redisClient.HSet(REDIS_FRIERND+strUid, REDIS_FRIEND_DATA_INVITE, string( byteFriend ) )
}
func (mgr *FriendManager)AddApplyUid( srcUid int64, decUid int64 ,applieMax int) { //插入申请列表
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	strUid := strconv.FormatInt(srcUid, 10)
	strApply := mgr.redisClient.HGet(REDIS_FRIERND+strUid, REDIS_FRIEND_APPLY)

	var applys []int64
	err := json.Unmarshal([]byte(strApply), &applys)
	if err != nil{
		log.Debug("uid (%v) add apply err >> %+v ", srcUid, err )
		return
	}
	for _, uid := range applys{//遍历看下有没有重复的
		if uid == decUid{//重复了 就不需要了
			return
		}
	}
	if len(applys) >= applieMax {	// 超上限
		return
	}
	applys = append( applys, decUid )//插入列表内
	byteApply, _ := json.Marshal( applys )
	mgr.redisClient.HSet(REDIS_FRIERND+strUid, REDIS_FRIEND_APPLY, string( byteApply ) ) //申请数据
}

func ( mgr *FriendManager )AddFriendUid( srcUid int64, decUid int64,maxFriend int ){//增加好友ID
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()

	strUid := strconv.FormatInt(srcUid, 10)
	strFriend := mgr.redisClient.HGet(REDIS_FRIERND+strUid, REDIS_FRIEND_DATA)

	var friends []*Friend

	err := json.Unmarshal( []byte(strFriend), &friends)
	if err != nil{
		log.Debug("uid (%v) add friend err >> %+v ", srcUid, err )
		return
	}
	if len(friends) >= maxFriend {	// 好友数超上限了
		return
	}
	for _, f := range friends{
		if f.Uid == decUid{//已经是好友了 所以就不用反复添加了
			return
		}
	}
	friend := &Friend{
		Uid:decUid,
	}

	friends = append( friends, friend )
	byteFriend, _ := json.Marshal( friends )
	mgr.redisClient.HSet(REDIS_FRIERND+strUid, REDIS_FRIEND_DATA, string( byteFriend ) ) //申请数据
}

func ( mgr *FriendManager )DelFriendUid( srcUid int64, decUid int64 ){//移除好友ID
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()

	strUid := strconv.FormatInt(srcUid, 10)
	strFriend := mgr.redisClient.HGet(REDIS_FRIERND+strUid, REDIS_FRIEND_DATA)

	var friends []*Friend

	err := json.Unmarshal( []byte(strFriend), &friends)
	if err != nil{
		log.Debug("uid (%v) del friend err >> %+v ", srcUid, err )
		return
	}

	idx := int(-1)
	for i, f := range friends{
		if f.Uid == decUid{//找到这个好友了
			idx = i
			break
		}
	}
	if idx >= 0{//哎哟找到了哟
		friends = append( friends[:idx], friends[idx+1:]...)
	}
	byteFriend, _ := json.Marshal( friends )
	mgr.redisClient.HSet(REDIS_FRIERND+strUid, REDIS_FRIEND_DATA, string( byteFriend ) ) //更新
}

func ( mgr *FriendManager )GiftSend( srcUid int64, decUid int64 ) { //礼品赠送
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()

	strDecUid := strconv.FormatInt(decUid, 10)
	strFriend := mgr.redisClient.HGet(REDIS_FRIERND+strDecUid, REDIS_FRIEND_DATA)
	var friends []*Friend
	err := json.Unmarshal( []byte(strFriend), &friends)
	if err != nil{
		return
	}
	isUpdate := false
	for _, f := range friends{
		if f.Uid == srcUid{
			if f.GiftState != 0x00 {//就不用反复设置了
				return
			}
			f.GiftState = 0x01 //设置为有东西领取状态
			isUpdate = true
			break
		}
	}
	if isUpdate {	// 有改动，更新
		byteFriend, _ := json.Marshal( friends )
		mgr.redisClient.HSet(REDIS_FRIERND+strDecUid, REDIS_FRIEND_DATA, string( byteFriend ) ) //更新
	} else {	// 没改动，是渠道好友
		chanFriendsStr := mgr.redisClient.HGet(REDIS_FRIERND+strDecUid,REDIS_FRIEND_DATA_CHAN)	// 渠道好友
		chanFriends := make(map[int64]int32)
		err := json.Unmarshal([]byte(chanFriendsStr),&chanFriends)
		if err != nil {
			return
		}
		v,ok := chanFriends[srcUid]
		if ok && v == Gift_State_Recv {	// 存在 有领取的状态
			return
		} else {
			chanFriends[srcUid] = Gift_State_Recv
		}
		js,_ := json.Marshal(chanFriends)
		mgr.redisClient.HSet(REDIS_FRIERND+strDecUid,REDIS_FRIEND_DATA_CHAN,string(js))	// 更新回去
	}

}

func (mgr *FriendManager)assert(){
	if mgr.redisClient == nil{
		panic("friend mamanger redis is nil")
	}
}