package sol

import (
	"strconv"
	"aliens/common/character"
	"time"
	"encoding/json"
	"strings"
)

type User struct {
	//玩家数据
	Id         int32   `json:"id"`
	Uid        int64   `json:"a"` //用户唯一ID
	Account    string  `json:"b"` //用户帐号
	BindAcc    string  `json:"c"` // 绑定账号
	Login      int64   `json:"d"` //登录时间
	Logout     int64   `json:"e"` //登出时间
	Register   int64   `json:"f"` //注册时间
	Country    int32   `json:"g"` // 国家
	State      int32   `json:"h"` //帐号状态
	Name       string  `json:"i"` //昵称
	Sex        int32   `json:"j"` //性别
	Icon       int32   `json:"k"` //icon 的 id
	IconUrl    string  `json:"l"` //icon 的 url 的地址
	Career     int32   `json:"m"` //职业
	Level      int32   `json:"n"` //等级
	Exp        int32   `json:"o"` //经验值
	Gold       int32   `json:"p"` //金币
	Diamond    int32   `json:"q"` //钻石
	PayDiamond int32   `json:"r"` //付费的钻石【主要作用于Vip】
	Cash       float32 `json:"s"` //充值的总额
	BoxScore   int32   `json:"t"` // 宝箱积分
	Vit        int32   `json:"u"` //体力
	VitTime    int64   `json:"v"` //体力下一次的恢复时间
	LastLogin  int32   `json:"w"` //连续登录天数
	SumLogin   int32   `json:"x"` //累计登录天数
	//Guide int32 //新手引导
	Sign     int32 `json:"y"` //签到数据
	SignTime int64 `json:"z"` //签到时间

	Chests       string `json:"a1"` //宝箱数据
	Heroes       string `json:"a2"` //英雄数据
	Formation    string `json:"a3"` //阵容数据
	FormationIdx int32  `json:"a4"` //当前阵容的索引号

	ReelFormation string `json:"a5"` //卷轴的阵容数据
	Reels         string `json:"a6"` //卷轴数据

	Battle          string `json:"a7"` //战斗数据
	Chapter         string `json:"a8"` //章节数据
	ChapterNodePass string `json:"a9"` // 章节关卡通关记录（只针对最新一章节）
	BossBattle      string `json:"a0"` // boss战斗数据

	Achievement string `json:"b1"` //成就数据

	TaskDailyTime int64  `json:"b2"` //每日任务的刷新时间
	TaskDaily     string `json:"b3"` //每日任务的数据
	TaskMain      string `json:"b4"` //主线任务的数据

	Equip    string `json:"b5"` //玩家的装备数据
	EquipBag string `json:"b6"` //玩家的装备背包数据

	TalentPoint int32  `json:"b7"` //玩家的天赋点数
	Talent      string `json:"b8"` //玩家的天赋数据
	// 矿山
	MineLv          int32  `json:"b9"` // 矿山等级
	Badge           int32  `json:"b0"` // 徽章
	PlunderNum      int32  `json:"c1"` // 掠夺次数
	PlunderTime     int64  `json:"c2"` // 下次掠夺+1时间
	MineBattleInfo  string `json:"c3"` // 矿战情报
	MiningEndTime   int64  `json:"c4"` // 开采结束时间
	MiningCD        int64  `json:"c5"`
	MineArmy        string `json:"c6"` // 玩家矿战数据
	MineExtraReward string `json:"c7"` // 矿山额外奖励次数
	// 竞技场
	ArenaSID          int32             `json:"c8"` // 赛季
	ArenaScore        int32             `json:"c9"` // 竞技场积分
	Honor             int32             `json:"c0"` // 荣耀点
	DailyHonor        int32             `json:"d1"` // 每日获得的荣耀点
	HonorItemBuyTimes string            `json:"d2"` // 荣耀商店购买次数 key:商品id value:剩余次数
	DayTimes          int32             `json:"d3"` // 竞技场每日进行场次
	ShopRefreshTime   int64             `json:"d4"` // 荣耀商店刷新时间
	GameNum           int32             `json:"d5"` // 竞技场场次
	GameWin           int32             `json:"d6"` // 竞技场胜场
	Streak            int32             `json:"d7"` // 连胜，连败纪录（负数连败）
	InArenaGame       bool              `json:"d8"` // 是否在竞技场中	（断线重连用）
	Mails             string            `json:"d9"` // 邮件数据（id,state#...)
	DailyRefineNum    int32             `json:"d0"` // 每日指定熔炼次数
	Star              int32             `json:"e1"` // 星星
	ClientTag         map[int32][]int32 `json:"e2"` // 客户端标识

	ChestFree    string `json:"e3"` //免费宝箱数据
	ChestCount   int32  `json:"e4"` // 章节宝箱获取的计数器
	GrowthValue  int32  `json:"e5"` // 成长值
	VitBuyTimes  int32  `json:"e6"` // 体力购买次数
	VipGift      string `json:"e7"` //vip 已经领取礼包的数据 （月卡过期时间#领取奖励日期#已领取充值奖励）
	MonCardState string `json:"e8"` // 月卡奖励状态
	//MonCardExpiryTime int64 // 月卡过期时间
	//RecMonCardDay int // 领取奖励日期
	ChargeRewards   string `json:"e9"` // 已领取的累计充值奖励
	ChangeNameTimes int32  `json:"e0"` // 改名次数
	//社交
	SocialFriends        string          `json:"f1"` //好友数据
	SocialApplies        string          `json:"f2"` //申请列表
	SocialDeletes        string          `json:"f3"` //准备删除列表
	//SocialChannelFriends map[int64]int32 `json:"f4"` // 渠道好友数据
	SocialInvitee		 string			 `json:"f4"`	// 邀请码好友
	SocialGiftSendCUid   []int64         `json:"f5"` // 赠送过的渠道好友
	InviteCode           int64           `json:"f6"` // 填的邀请码
	// 世界boss
	HitBossTime       int64 `json:"f7"`
	HitBossTimes      int32 `json:"f8"`
	HitBossShareTimes int32 `json:"f9"`
	// 小游戏
	MiniGame string `json:"f0"` // 小游戏副本

	Items    string `json:"g1"` // 道具
	Activity string `json:"g2"` // 活动数据
	// 分享
	ShareDB           string          `json:"g3"` // 分享数据
	DailyLinkCount    int32           `json:"g4"`
	DailyChestShare   int32           `json:"g5"`
	RecRewardId       int32           `json:"g6"`
	DailyChapterShare int32           `json:"g7"` // 关卡分享次数
	DailyVitShare     int32           `json:"g8"` // 体力分享次数
	PermanentAttr     map[int32]int32 `json:"g9"` // 主角永久属性加成

	SpendRewards string `json:"g0"`	// 累计消费奖励id
	SpendDiamond int32 `json:"k1"`	// 累计消费值
	PayShopId int32	`json:"k2"`	// 支付商品id
}

func (u *User) Parse(arrStr []string) {
	u.Uid, _ = strconv.ParseInt(arrStr[0], 10, 64)      //用户id
	u.Account = arrStr[1]                               //用帐号
	u.Name = arrStr[2]                                  //昵称
	u.Sex = character.StringToInt32(arrStr[3])          //性别
	u.Login, _ = strconv.ParseInt(arrStr[4], 10, 64)    //登录时间
	u.Logout, _ = strconv.ParseInt(arrStr[5], 10, 64)   //登出时间
	u.Register, _ = strconv.ParseInt(arrStr[6], 10, 64) //注册时间
	u.State = character.StringToInt32(arrStr[7])        //状态
	u.Level = character.StringToInt32(arrStr[8])        //等级
	u.Exp = character.StringToInt32(arrStr[9])          //经验值
	u.Career = character.StringToInt32(arrStr[10])      //职业

	u.Icon = character.StringToInt32(arrStr[11])       //头像id
	u.IconUrl = arrStr[12]                             //头像地址
	u.Gold = character.StringToInt32(arrStr[13])       //金币
	u.Diamond = character.StringToInt32(arrStr[14])    //钻石
	u.PayDiamond = character.StringToInt32(arrStr[15]) //充值的钻石数量
	u.Cash = character.StringToFloat32(arrStr[16])     //充值金额

	u.Vit = character.StringToInt32(arrStr[17])         //体力值
	u.VitTime, _ = strconv.ParseInt(arrStr[18], 10, 64) //下一次恢复体力的时间

	u.Chests = arrStr[19] //宝箱数据

	u.LastLogin = character.StringToInt32(arrStr[20]) //连续登录天数
	u.SumLogin = character.StringToInt32(arrStr[21])  //累计登录天数

	u.ChangeNameTimes = character.StringToInt32(arrStr[22]) //新手引导
	u.Sign = character.StringToInt32(arrStr[23])            //签到数据
	u.SignTime, _ = strconv.ParseInt(arrStr[24], 10, 64)    //签到时间
	u.ChestFree = arrStr[25]                                //免费宝箱数据
	u.GrowthValue = character.StringToInt32(arrStr[26])     //成长值
	u.VipGift = arrStr[27]                                  //vip礼包的领取数据
	u.Heroes = arrStr[28]                                   //英雄数据
	u.Formation = arrStr[29]                                //阵容数据
	u.FormationIdx = character.StringToInt32(arrStr[30])    //当前阵容的索引号

	u.Battle = arrStr[31]        //战斗相关数据
	u.Chapter = arrStr[32]       //章节数据
	u.ReelFormation = arrStr[33] //卷轴阵容
	u.Reels = arrStr[34]         //卷轴数据

	u.Achievement = arrStr[35] //成就标签数据

	u.TaskDailyTime, _ = strconv.ParseInt(arrStr[36], 10, 64) //每日任务的刷新时间
	u.TaskDaily = arrStr[37]                                  //每日任务的数据
	u.TaskMain = arrStr[38]                                   //主线任务的数据

	u.Equip = arrStr[39]    //装备数据
	u.EquipBag = arrStr[40] //装备背包

	u.TalentPoint = character.StringToInt32(arrStr[41]) //天赋点数
	u.Talent = arrStr[42]                               //天赋数据
	// 39 - 44 为 矿山数据
	u.MineLv = character.StringToInt32(arrStr[43])
	u.Badge = character.StringToInt32(arrStr[44])
	u.PlunderNum = character.StringToInt32(arrStr[45])
	u.MiningEndTime = character.StringToInt64(arrStr[46])
	u.MineBattleInfo = arrStr[47]
	u.PlunderTime = character.StringToInt64(arrStr[48])
	//
	u.SocialFriends = arrStr[49]                        //好友数据
	u.SocialApplies = arrStr[50]                        //申请数据
	u.SocialDeletes = arrStr[51]                        //准备删除数据
	u.VitBuyTimes = character.StringToInt32(arrStr[52]) // 每日购买体力次数
	u.Id = character.StringToInt32(arrStr[53])          // 自增id
	u.ChestCount = character.StringToInt32(arrStr[54])  // 获取宝箱数量的计数器
	//u.ShareDB = arrStr[54]
	json.Unmarshal([]byte(arrStr[55]), &u.ClientTag) // 客户端标识
	u.DailyRefineNum = character.StringToInt32(arrStr[56])
	// 分享数据
	u.DailyChestShare = character.StringToInt32(arrStr[57])
	u.DailyLinkCount = character.StringToInt32(arrStr[58])
	u.RecRewardId = character.StringToInt32(arrStr[59])
	u.ShareDB = arrStr[60]
	u.Country = character.StringToInt32(arrStr[61])
	u.ArenaSID = character.StringToInt32(arrStr[62])
	u.ArenaScore = character.StringToInt32(arrStr[63])
	u.GameNum = character.StringToInt32(arrStr[64])
	u.GameWin = character.StringToInt32(arrStr[65])
	u.Honor = character.StringToInt32(arrStr[66])
	u.Streak = character.StringToInt32(arrStr[67])
	u.DailyHonor = character.StringToInt32(arrStr[68])
	u.MineArmy = arrStr[69]
	u.ShopRefreshTime = character.StringToInt64(arrStr[70])   // 竞技场荣誉商店
	u.HonorItemBuyTimes = arrStr[71]                          // 荣誉商品购买次数
	u.HitBossTimes = character.StringToInt32(arrStr[72])      // 攻击boss次数
	u.HitBossShareTimes = character.StringToInt32(arrStr[73]) // 攻击boss分享次数
	u.HitBossTime = character.StringToInt64(arrStr[74])       // 攻击那个时间段boss
	u.ChargeRewards = arrStr[75]                               // 周卡领取状态
	u.MineExtraReward = arrStr[76]                            // 矿山额外奖励
	u.DailyChapterShare = character.StringToInt32(arrStr[77]) // 每日关卡步数分享上限
	u.DailyVitShare = character.StringToInt32(arrStr[78])     // 每日体力分享上限
	u.BindAcc = arrStr[79]                                    // 绑定账号
	u.Star = character.StringToInt32(arrStr[80])              // 竞技场星星数
	u.Activity = arrStr[81]                                   // 活动数据
	u.InviteCode = character.StringToInt64(arrStr[82])        // 邀请码
	u.ChapterNodePass = arrStr[83]
	u.BoxScore = character.StringToInt32(arrStr[84])
	u.Items = arrStr[85]
	u.SocialInvitee = arrStr[86]
}

func (u *User) New(uid int64, account string) {
	u.Uid = uid
	u.Account = account            //帐号
	u.Register = time.Now().Unix() //注册时间
	u.Name = ""
	u.State = 0 //状态
	u.Level = 1 //等级
	u.Exp = 0   //经验值
	u.Career = 1
	u.Icon = 0      //头像id
	u.IconUrl = ""  //头像地址
	u.Gold = 10000  //金币
	u.Diamond = 100 //钻石
	u.Sex = 1       // 默认男
	// todo 190729 sdw版本 初始化为中国
	u.Country = 101
	u.Vit = 100   //体力值
	u.VitTime = 0 //下一次体力的恢复时间

	u.Sign = 0            //签到数据
	u.SignTime = 0        //签到时间
	u.DailyRefineNum = 50 // 每日指定熔炼次数

	//创建新用户时候 默认给他4个宝箱
	u.Chests = "1|1,0,-1#2,0,-1#3,0,-1#4,0,-1#" //"1|1,1,-1#2,2,-1#3,3,-1#4,4,-1#"

	u.FormationIdx = 0
	u.GrowthValue = 1

	//创建新用户时候 默认给他4个怪物
	u.Heroes = ""
	u.Formation = ""

	//卷轴
	u.ReelFormation = "0#0#0#0#0#"
	u.Reels = ""
	u.ChangeNameTimes = 0
}

func (u *User) GetMaxChapter() int32 {
	chapterMax := int32(0)
	arrStr := strings.Split(u.Chapter, "#")
	for _, strChapter := range arrStr {
		if strChapter == "" {
			continue
		}
		arrStrSub := strings.Split(strChapter, ",")
		if len(arrStrSub) < 2 {
			continue
		}
		id := character.StringToInt32(arrStrSub[0])
		state := character.StringToInt32(arrStrSub[1])
		if state > 0 && id > chapterMax {
			chapterMax = id
		}
	}
	return chapterMax
}
