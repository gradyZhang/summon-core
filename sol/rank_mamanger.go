package sol

import (
	"sync"
	"aliens/common/cache"
	"time"
	"strconv"
	"sort"
	"strings"
)
const(
	REDIS_RANK_LEVEL = "rRankLv_" //等级排行
	REDIS_RANK_LV_TEMP = "trRankLv_"	// 等级排行临时存储
	REDIS_RANK_GROWTHVALUE = "rRankGrowth"	// 成长值排行
	REDIS_RANK_ACHIEVE = "rRankAchi"	// 成就排行
	REDIS_RANK_BOSS = "rRankBoss"	// 世界boss排行
	REDIS_RANK_ARENA = "rRankArena"	// 竞技场排行
)

const(
	MAX_RANK_DEFAULT = 500	// 默认上限
	MAX_RANK_LEVEL = 50000 //等级排行的数量上限
	MAX_RANK_BOSS = 500// boss排行上限
	MAX_RANK_ARENA = 5000// arena排行上限
)

type RankInfo struct{//排行榜信息
	Uid int64
	Value int64
	Rank int32 //排名 -1 表示未在排名内
}

type RankManager struct{
	sync.Mutex
	redisClient *cache.RedisCacheClient
	levelRank map[string]RankObjList	// 等级排行榜
	growthValueRank RankObjList	// 成长值排行
}

var rankMgr *RankManager
var onceRank sync.Once

//单例
func GetInsRankMgr() *RankManager {
	onceRank.Do(func() {
		rankMgr = &RankManager{}
	})
	return rankMgr
}

//初始化
func ( mgr *RankManager )Init(  redisInfo string,redisPw string, dbIdx int,sentinelAddrs []string ){
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
	mgr.levelRank = make(map[string]RankObjList)
	mgr.growthValueRank = RankObjList{}
}

//刷新Redis 【注：只允许 一台服务器进行操作】
func (mgr *RankManager)FlushDB(){
	if mgr.redisClient == nil{
		return
	}
	mgr.redisClient.FlushDB()
}

//获取排行榜数据
func (mgr *RankManager)GetRank( rankKey string, page int32, pageNum int32 )[]RankInfo{
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	if pageNum > 50 {	// 做下限制，一页最多50个数据
		pageNum = 50
	}
	// 要获取的排行榜范围
	begin := page*pageNum
	end := ( page+3 )*pageNum //默认多给三页

	var infos []RankInfo
	ranks := mgr.redisClient.ZRevRangeWithScore(rankKey,begin,(end-1) )
	for _, r := range ranks {
		begin += 1//默认从 0 开始
		if rankKey == REDIS_RANK_LEVEL{
			if begin >= MAX_RANK_LEVEL {	// 不在排行内
				break
			}
		}
		info := RankInfo{}
		info.Uid, _ = strconv.ParseInt( r.Member, 10, 64 )
		info.Value = r.Score
		info.Rank = begin

		infos = append( infos, info )
	}
	return infos
}

//获取自己的排行榜数据
func (mgr *RankManager)GetUserRank( rankKey string, uid int64 )*RankInfo{
	mgr.Lock()
	defer mgr.Unlock()

	mgr.assert()

	uIDStr := strconv.FormatInt(uid,10)
	rank := mgr.redisClient.ZRevRank(rankKey,uIDStr)

	if rankKey == REDIS_RANK_LEVEL{
		if rank >= int(MAX_RANK_LEVEL) {	// 不在排行内
			rank = -1
		}
	} else if rankKey == REDIS_RANK_BOSS {
		if rank >= int(MAX_RANK_BOSS) {	// 不在排行内
			rank = -1
		}
	}

	value := mgr.redisClient.ZScore(rankKey, uIDStr )
	return &RankInfo{ Uid:uid, Rank:int32(rank+1), Value:value}
}

// 获取指定排名区间排行数据
func (mgr *RankManager) GetRankByRange(rankKey string,begin int32,end int32) []RankInfo {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.assert()
	var infos []RankInfo
	ranks := mgr.redisClient.ZRevRangeWithScore(rankKey,begin,end)
	for _, r := range ranks {
		begin += 1//默认从 0 开始
		info := RankInfo{}
		info.Uid, _ = strconv.ParseInt( r.Member, 10, 64 )
		info.Value = r.Score
		info.Rank = begin

		infos = append( infos, info )
	}
	return infos
}

// 获取排行榜长度
func (mgr *RankManager) GetRankLen(rankKey string) int {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.assert()
	ranks := mgr.redisClient.ZRevRangeWithScore(rankKey,0,-1)
	return len(ranks)
}

// 更行排行榜
func (mgr *RankManager) Update(rankKey string, uId int64, value interface{}) {
	mgr.Lock()
	defer mgr.Unlock()

	if mgr.redisClient == nil{
		return
	}
	uIDStr := strconv.FormatInt(uId,10)
	mgr.redisClient.ZAdd(rankKey, value, uIDStr)

	ranks := mgr.redisClient.ZRevRangeWithScore(rankKey, 0, -1)	//获取该key 下面的所有数据

	rankMax := MAX_RANK_DEFAULT
	if strings.HasPrefix(rankKey,REDIS_RANK_LEVEL) {
		rankMax = MAX_RANK_LEVEL
	} else if rankKey == REDIS_RANK_BOSS {
		rankMax = MAX_RANK_BOSS
	} else if rankKey == REDIS_RANK_ARENA {
		rankMax = MAX_RANK_ARENA
	}
	if len(ranks) > int(rankMax * 2) {//默认存两倍大小
		// 超过了，移除最后一名
		mgr.redisClient.ZRemRangeByRank(rankKey,0,0)
	}
}

// 删除玩家等级排行榜
func (mgr *RankManager) DelLvRankKey() {
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	ret := mgr.redisClient.Keys(REDIS_RANK_LEVEL+"*")
	c := mgr.redisClient.GetConn()
	for _, key := range ret {
		// 拷贝一份排行榜
		ranks := mgr.redisClient.ZRevRangeWithScore(key,0,200)
		for _,r := range ranks {
			c.Send("ZADD","t"+key,r.Score,r.Member)
		}
	}
	c.Flush()
	c.Receive()
	c.Close()
	for _,key := range ret {
		mgr.redisClient.DelData(key)
	}
}

// 移除临时玩家等级排行
func (mgr *RankManager) DelTempLvRankKey() {
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	ret := mgr.redisClient.Keys(REDIS_RANK_LV_TEMP+"*")
	for _,k := range ret {
		mgr.redisClient.DelData(k)
	}
}

// 清空排行榜
func (mgr *RankManager) CleanRank(rankKey string) {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.assert()
	mgr.redisClient.DelData(rankKey)
}

// 移除排行行指定元素
func (mgr *RankManager) RemRankMember(rankKey string,member ...interface{}) {
	mgr.Lock()
	defer mgr.Unlock()
	mgr.assert()
	mgr.redisClient.ZRem(rankKey,member)
}

func (mgr *RankManager)assert(){
	if mgr.redisClient == nil{
		panic("rank mamanger redis is nil")
	}
}

// 批量保存
func (mgr *RankManager) SaveRank() {
	mgr.Lock()
	defer mgr.Unlock()
	if mgr.redisClient == nil{
		return
	}
	c := mgr.redisClient.GetConn()
	// 排行插入
	for k,lr := range mgr.levelRank {
		for _,r := range lr {
			strUid := strconv.FormatInt( r.UID, 10 )
			c.Send("ZADD",k,r.Value,strUid)
		}
	}
	for _,r := range mgr.growthValueRank {
		strUid := strconv.FormatInt( r.UID, 10 )
		c.Send("ZADD",REDIS_RANK_GROWTHVALUE,r.Value,strUid)
	}
	if len(mgr.levelRank) > 0 || mgr.growthValueRank.Len() > 0 {
		c.Flush()
		c.Receive()
	}
	c.Close()
	// 更完了，置空了
	mgr.levelRank = make(map[string]RankObjList)
	mgr.growthValueRank = RankObjList{}
}

// 更新排行榜
func (mgr *RankManager) UpdateRank(rankKey string,uid int64,score int32) {
	mgr.Lock()
	defer mgr.Unlock()
	var rankList RankObjList
	rankMax := MAX_RANK_DEFAULT
	if strings.Contains(rankKey,REDIS_RANK_LEVEL) {
		rankMax = MAX_RANK_LEVEL
		if _,ok := mgr.levelRank[rankKey];!ok {
			mgr.levelRank[rankKey] = RankObjList{}
		}
		rankList = mgr.levelRank[rankKey]
		score += 1	// lv * 1000 + maxChapter (最大章节是到达章节)
	} else if rankKey == REDIS_RANK_GROWTHVALUE {
		rankMax = MAX_RANK_DEFAULT
		rankList = mgr.growthValueRank
	}
	rankMax = rankMax * 2	// 存两倍的
	if rankList.Len() >= rankMax {
		if rankList[rankMax-1].Value < score {	// 比小 要挤出排行了
			// 替换
			rankList[rankMax-1].UID = uid
			rankList[rankMax-1].Value = score
		}
	} else {
		rankList = append(rankList,&RankObj{UID:uid,Value:score})
	}
	// 重新排序
	sort.Sort(rankList)
	if strings.Contains(rankKey,REDIS_RANK_LEVEL) {
		mgr.levelRank[rankKey] = rankList
	} else if rankKey == REDIS_RANK_GROWTHVALUE {
		mgr.growthValueRank = rankList
	}
}

type RankObjList []*RankObj
func (list RankObjList)Len() int{
	return len(list)
}
func (list RankObjList)Less(i,j int) bool {
	if list[i].Value > list[j].Value {
		return true
	} else if list[i].Value < list[j].Value {
		return false
	} else {
		return list[i].UID < list[j].UID
	}
}
func (list RankObjList) Swap(i, j int) {
	var temp *RankObj = list[i]
	list[i] = list[j]
	list[j] = temp
}
type RankObj struct {
	UID   int64
	Value int32
}