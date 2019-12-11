package def
const DEFAULT_COUNTRY int32 = 101 // 默认国家
const GRID_ID_BASE int32 = 100000 // 沙盘上id的基础倍数
const CHEST_MAX_NUM int32 = 4     //宝箱的最大数量
const RankTop int32 = 50          // 排行榜排名数

const MailStorageLimit = 100 // 邮件存储上限
const MailVisibleLimit = 50  // 邮件客户端显示上限

const HeroTalentResetSpend int32 = 100 // 英雄觉醒重置花费
const MineBattleInfoLen int = 10       // 矿山情报长度
// 小游戏相关
const GameRecvTime = 1 * 60 * 60 // 小游戏恢复时长
const GameMaxTimes = 5           // 小游戏次数上限
const ChangeNameFreeTimes int32 = 1	 // 免费改名次数
const (
	MiniGameTypeNone  int32 = iota + 1000
	MiniGameTypeGold   // 金币游戏
	MiniGameAdventure  // 奇遇
	MiniGameBigTable   // 大沙盘
)

const (
	RewardTypeNull    int32 = iota // value --> 0
	RewardTypeGold                 //金币 1
	RewardTypeDiamond              //钻石 2
	RewardTypeHero                 //英雄家族 3
	RewardTypeEquip                //装备 4
	RewardTypeReel                 //卷轴 5
	RewardTypeChest                // 宝箱 6
	RewardTypeVit                  // 体力 7
	RewardRoleArr                  // 主角属性 8
	RewardRandomClip               // 随机家族碎片 9
	RewardItem                     // 活动道具 10
	RewardRandomReel               // 随机卷轴 11
	RewardMaxReel                  // 随机卷轴，根据最高品质 12
	RewardMaxClip                  // 随机家族碎片，根据最高品质 13
	RewardRoleExp                  // 主角经验 14
	RewardMonCard                  // 月卡 15
	RewardSpecial                  // 特殊奖励（奇遇 神秘商店 神殿） 16
	RewardAct					   // 活动商品 17
	RewardRandomEquip			   // 随机装备 新 18
	RewardDrawCard				   // 抽卡 19
	RewardTaskScore				   // 每日任务积分 20
)

// 货币类型
const (
	CurTypeGold    int32 = iota + 1 // 金币
	CurTypeDiamond                  // 钻石
	CurTypeRMB                      // 人民币
)

const ( //成就类型
	AchievementMarkNull            int32 = iota //value --> 0
	AchievementMarkPlayerLv                     //玩家等级 1
	AchievementMarkGoldSum                      //玩家累计拥有的金 2
	AchievementMarkHeroFamily                   //玩家拥有的家族 3
	AchievementMarkChapter                      //玩家已通關章節 4
	AchievementMarkMoksterKill                  //玩家已消灭的怪物 5
	AchievementMarkChestOpen                    //玩家开启宝箱的数量 6
	AchievementMarkMineLv                       //玩家矿的等级 7
	AchievementMarkMinePlunder                  //玩家矿战掠夺次数 8
	AchievementMarkHeroFarmilyLv                //玩家所有家族的等级总和 9
	AchievementHeroTalentUnlock                 // 解锁家族觉醒 10
	AchievementGet                              // 占坑
	AchievementMarkFinishDailyTask              // 累计完成每日任务 12
	AchievementMarkGetReel                      // 累计获得卷轴数 13
	AchievementMarkGetReelOrange                // 获取橙色卷轴 14
	AchievementMarkGetReelPurple                // 获取紫色卷轴 15
	AchievementMarkGetReelBlue                  // 获取蓝色卷轴 16
	AchievementMarkUseReelOrange                // 使用橙色卷轴 17
	AchievementMarkUseReelPurple                // 使用紫色卷轴 18
	AchievementMarkUseReelBlue                  // 使用蓝色卷轴 19
	AchievementMarkGetDiamond                   // 累计获得钻石数量 20
	AchievementMarkComOrange                    // 合成橙色怪 21
	AchievementMarkComPurple                    // 合成紫色怪 22
	AchievementMarkComBule                      // 合成蓝色怪 23
	AchievementMarkHeroSkillLvUp                //玩家怪物家族技能升级 24
	AchieveOpenShop                             // 打开商店 25
	AchieveBuyShop                              // 购买商品 26
	AchieveArenaBattle                          // 进行竞技场战斗 27
	AchieveRefine                               // 熔炼 28
	AchieveWearEquip                            // 穿普通装备 29
	AchieveWearFashion                          // 穿戴任意品质套装 30
	AchieveOpenArenaShop                        // 打开竞技场商店 31
	AchieveBuyArenaShop                         // 购买竞技场商店 32
	AchieveOpenChapterChest                     // 开启通关宝箱 33
	AchieveShopChestA                           // 开启商城 星辉宝箱（蓝色那个）34
	AchieveShopChestS                           // 开启商城 月灵宝箱（紫色那个）35
	AchieveShopChestSS                          // 开启商城 日瑶宝箱（橙色那个）36
	AchieveChangeCareer                         // 转职 37
	AchieveFormation3B                          // 召唤3个B级家族 38
	AchieveFormation3A                          // 召唤3个A级家族 39
	AchieveFormation1S                          // 召唤1个S级家族 40
	AchieveFormation1SS                         // 召唤1个SS级家族 41
	AchieveHitBoss                              // 打世界boss 42
	AchieveChangeIcon                           // 更换头像 43
	AchieveChapterCE                            // 关卡中战力到达 44
	AchieveBackGame                             // 回到游戏 45
	AchieveUseHeroSpell                         // 使用怪物技能 46
	AchieveTableCombo                           // 沙盘多消 47
	AchieveInvite                               // 邀请好友 48
	AchieveShareFeed                            // 分享feed 49
	AchieveReadAds                              // 看广告 50
	AchieveSpendDiamond							// 累计花费钻石 51
	AchieveChangeName							// 改名成就 52
)

const ( //每日任务类型
	DailyTaskMarkNull        int32 = iota
	DailyTaskMarkChapterPass  //章节通关
)

const ( //章节状态
	ChapterStateUnPassed        int32 = iota //为通关
	ChapterStatePassed                       //已通关
	ChapterStatePassedImperfect              // 通关但不可扫荡
)
const ( //装备的部位
	EquipPartNull         int32 = iota
	EquipPartWeapon        //武器 1
	EquipPartHead          //头部 2
	EquipPartClothes       //衣服 3
	EquipPartPants         //裤子 4
	EquipPartShoes         //鞋子 5
	EquipPartWrister       //护腕 6
	EquipPartCloak         //披风 7
	EquipPartFationWeapon  //时装武器 8
	EquipPartFation        //时装9 整体

	EquipPartMax = EquipPartFation //标底最大值
)

const (
	EquipQualityNull   int32 = iota //装备品质
	EquipQualityWhite               //白色 1
	EquipQualityGreen               //绿色 2
	EquipQualityBlue                //蓝色 3
	EquipQualityPurple              //紫色 4
	EquipQualityOrange              //橙色 5
	EquipQualityRed                 //红色 6
)

const (
	HeroQualityNull   int32 = iota //英雄品质
	HeroQualityWhite               //白色 1
	HeroQualityGreen               //绿色 2
	HeroQualityBlue                //蓝色 3
	HeroQualityPurple              //紫色 4
	HeroQualityOrange              //橙色 5
	HeroQualityRed                 //红色 6
)

const ( //属性枚举
	AttrNull    int32 = iota
	AttrHp       //主角血量上限【固定值】 1
	AttrDamage   //主角伤害【固定值】 2
	AttrPhyAtk   //主角物理强度【固定值】 3
	AttrMagAtk   //主角魔法强度【固定值】 4
	AttrPhyDef   //主角物理防御【固定值】 5
	AttrMagDef   //主角魔法防御【固定值】 6
	AttrCritAtk  // 主角暴击【固定值】 7
	AttrStep     // 主角副本步数【固定值】 8
	AttrSave     // 主角保留【固定值】9

	Reserve                     // 预留 10
	AttrHpPCT                   // 主角血量【百分比】11
	AttrDamagePCT               // 主角伤害【百分比】12
	AttrPhyAtkPCT               // 主角物理强度【百分比】13
	AttrMagAtkPCT               // 主角魔法强度【百分比】14
	AttrPhyDefPCT               // 主角物理防御【百分比】15
	AttrMagDefPCT               // 主角魔法防御【百分比】16

	AttrPhyAtkDef	// 主角物强物防 17
	AttrMagAtkDef	// 主角魔强魔防 18
	AttrDoubleAtk	// 主角双攻 19
	AttrDoubleAtkPCT// 主角双攻千分比 20
	AttrDoubleDef	// 主角双防 21
	AttrDoubleDefPCT// 主角双防千分比 22
	AttrReserve		// 预留 23
	AttrAtkSpeed	// 主角攻速千分比 24

	AttrLeader	= 32	// 领导力 32
	AttrMonsterHp = 101 //增加怪物家族血量上限
)

const ( //日志操作
	LogOpNull             int32 = iota
	LogOpReward            // 奖励	1
	LogOpEquipSell         // 装备出售	2
	LogOpHeroCreate        // 怪物合成	3
	LogOpHeroLvUp          // 怪物升级	4
	LogOpHeroSkill         // 怪物技能升级	5
	LogOpHeroRefine        // 怪物熔炼	6
	LogOpHeroTalentReset   // 怪物天赋重置7
	LogOpShopBuy           // 商店购买	8
	LogOpMineBattle        // 矿战	9
	LogOpChest             // 宝箱	10
	LogOpChestBuy          // 商店宝箱购买 11
	LogOpVipGift           // vip礼包	12
	LogRoleTalent          // 主角天赋相关（重置天赋）	13
	LogOpBossRefreshCD     // 世界boss刷新CD时间	14
	LogOpMineRefreshEnemy  // 刷新矿战对手	15
	LogOpVitBuy            // 体力购买	16
	LogOpMonCard           // 月卡	17
	LogOpArenaShop         // 竞技场商店	18
	LogOpArena             // 竞技场	19
	LogOpChangeName        // 改名	20
	LogOpAchieve           // 成就	21
	LogOpArenaSeason       // 竞技场赛季结算 22
	LogOpBossHit           // 每次攻击boss 23
	LogOpAventure          // 奇遇 24
	LogOpChapter           // 关卡中 25
	LogOpMail              // 邮件 26
	LogMineReward          // 矿战奖励 27
	LogOpShare             // 分享奖励 28
	LogOpMyticShop         // 神秘商店 29
	LogOpSign              // 签到 30
	LogOpTask              // 任务 31
	LogOpWatchAdv          // 看广告奖励 32
	LogOpArenaBox          // 竞技场宝箱 33
	LogOpGiftCode		   // 礼包码 34
	LogOpCharge			   // 充值 35
	LogOpActPrize		   // 活动 大转盘 36
	LogOpChat 			   // 聊天 37
	LogOpNewFirstCharge	   // 新首充奖励 38
)

// 装备日志类型
const (
	EquipOpTypeGet  int32 = iota + 1 // 获得
	EquipOpTypeUp                    // 强化消耗
	EquipOpTypeSell                  // 出售
)

// 家族碎片日志类型
const (
	CardOpGetRefine int32 = iota + 101
	//CardOpGetRefine = "gRefine"	// 熔炼获取
	CardOpUseRefine 
	//CardOpUseRefine = "uRefine"	// 熔炼消耗
	CardOpUseUpgrade 
	//CardOpUseUpgrade = "uUp"	// 卡牌升级消耗
	CardOpUseCreate  // 创建消耗
)

// 成败
const (
	SUCCESS int32 = iota
	FAILED
	TOKEN_INVALID
)

//
const (
	FLASE int32 = iota
	TRUE  
)

// 每日获得荣耀点上限
const DailyHonorMax int32 = 100

const PlunderMineExp int32 = 10  // 夺矿获得经验
const PlunderMineBadge int32 = 3 // 夺矿获得徽章

// 主角职业
const (
	JobNone    int32 = iota // 0
	JobAdvent               // 冒险者 1
	JobWarrior              // 勇士	2
	JobFighter              // 战士 3
	JobKnight               // 游侠 4

	JobScholar  // 学者 5
	JobMage     // 法师 6
	JobDoctor   // 博士 7
)

// 英雄天赋 状态
const (
	HeroTalentNone   int32 = iota // 未解锁
	HeroTalentUnLock              // 解锁
)

// redis key
const (
//BlackList = "blacklist" // 黑名单
)

// redis db索引
const (
	REDIS_DB_MAIN   = iota //主 redis
	REDIS_DB_FRIEND        //好友 redis
	REDIS_DB_RANK          //排行榜 redis
	REDIS_DB_USER          // 用户数据
	REDIS_DB_MAIL          // 邮件数据
)

const (
	LoginRetSuc         int32 = iota //登录——成功 0x00
	LoginRetSame                     //登录——同号 0x01
	LoginRetSeal                     //登录——封号 0x02
	LoginRetHotUpdate                // 热更 3
	LoginRetSevMaintain              // 服务器维护中 4
	LoginRetErr         int32 = 0xFF //登录——异常错误
)

// 任务大类型
const (
	TaskTypeMain  int32 = iota //	主线任务
	TaskTypeDaily              // 每日任务
)

// 任务小类
const (
	TaskFinishNum        int32 = iota + 1 // 任务完成度 1
	TaskPassChapter                       // 通关关卡 2
	TaskSpendReel                         // 消耗卷轴 3
	TaskEquipUpgrade                      // 装备强化 4
	TaskOpenChest                         // 开宝箱 5
	TaskMinePlunder                       // 矿山争夺 6
	TaskMining                            // 矿山开采 7
	TaskSpendGold                         // 花费金币 8
	TaskSpendDiamond                      // 花费钻石 9
	TaskAreanPvp                          // 竞技场pvp 10
	TaskHeroUpgrade                       // 家族升级 11
	TaskRefine                            // 熔炼 12
	TaskShopBuy                           // 商店购买 13
	TaskTreasureMap                       // 藏宝图 14
	TaskBoss                              // 世界boss 15
	TaskTreasureIsle                      // 宝藏岛 16
	TaskPuzzle                            // 解密 17
	TaskGiveGift                          // 赠送好友 18
	TaskHeroSkillUpgrade                  // 英雄技能等级升级 19
	TaskArenaShopBuy                      // 荣誉商店购买 20
	TaskLogin                             // 今日登陆 21
	TaskWatchAdv						  // 看广告 22

	TaskChargeCount // 充值任意金额次数 23
	TaskNormalDraw	// 普通抽卡次数 24
	TaskWorldChat	// 世界聊天次数 25
	TaskVitBuy		// 体力购买 26
)

// 奖励状态
const (
	None     int32 = iota // 未完成
	Recv                  // 可领取
	Received              // 已领取
)

// 世界boss
const (
	BossBeginTime       = 9            // boss 开始时间 每天21点
	BossAtkTime   int64 = 60 * 60 * 23 // 持续23小时
	BossHitCDTime int64 = 60           // boss攻击间隔时间
)

// 排行榜类型
const (
	RankGrowthValue int32 = iota // 成长值排行 0
	RankLevel                    // 等级排行 1
	RankAchievement              // 成就排行 2
)

// 红点标识
const (
	Flag          int32 = iota
	FlagSignIn     // 签到 1
	FlagMail       // 邮件 2
	FlagDailyTask  // 每日任务 3
	FlagAchi       // 成就 4
	FlagVIP        // vip 5
	FlagMonCard    // 月卡 6
)

// 平台
const (
	PlatformH5     int32 = iota //h5 游客  浏览器
	PlatformGoogle 
	PlatformIos    
	PlatformFb      // facebook
	PlatformHiho	// hiho
	PlatformWX		// 5 wechat
	PlatformSDW		// 6 闪电玩
	PlatformXiaoDu	// 7 小度
)

// 订单状态
const (
	OrderStateSuccess int8 = iota // 成功
	OrderStateFailed              // 失败
	OrderStateCancel              // 取消订单
	OrderStateUnpay               // 未支付
)

// 月卡领取状态
const (
	MonCardNone     int32 = iota // 不可领 0
	MonCardRecv                  // 可领取 1
	MonCardReceived              // 已领取 2
	MonCardExpire                // 已过期 3
)

// 矿战情报状态
const (
	MineBattleNone  int32 = iota
	MineBattleCanDo  // 矿战情报未处理
	MineBattleDid    // 矿山情报已处理
)

// 邮件条件
const (
	MailCondNone       int32 = iota
	MailCondMaxChapter  // 最大通关章节
)

// 活动
const (
	ActLvGift      int32 = iota + 1 // 1 等级礼包
	ActDinner                       // 2 午晚餐
	ActFirstCharge                  // 3 首充
	ActMonCard                      // 4 月卡
	ActFirstGift                    // 5 新手特惠礼包
	ActDailyGift                    // 6 每日礼包
	ActLimitDraw                    // 7 限时抽卡
	ActCharge                       // 8 累计
	ActGrowthCapital                // 9 成长基金
	Act4                            // 10节日兑换
	ActNewSign                      // 11新手签到
	ActConsum						// 12累计消费
	ActLimitGift					// 13限时礼包
	ActNewFirstCharge1 				// 14新首充1
	ActNewFirstCharge2 				// 15新首充2
	ActNewFirstCharge3 				// 16新首充3
	ActPrizeWheel					// 17大转盘
	ActContinueCharge				// 18 连续充值活动
)

// 特权
const (
	PrivilegesBuyVit	int32 = iota + 1	// 1 额外购买体力
	PrivilegesStep	// 2 关卡步数增加
	PrivilegesDraw	// 3 每日免费高级抽
)

// 抽卡类型
const (
	DrawNormal int32 = iota + 1
	DrawHigh
)