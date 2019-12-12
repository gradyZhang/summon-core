package def

//消息包序列号
const (
	// sys
	ReqSysFriendApply int32 = 150
	ReqConnCenter int32 = 151	// 连接center server

	ReqCreateUser int32 = 51
	ReqMailUpdate int32 = 52
	ReqKillOff int32 = 54	// 踢人
	ReqClientLogout int32 = 55	// 客户端下线
	ReqSendMail int32 = 56	// 请求发送邮件
	ReqGift		int32 = 232 //请求礼包
	ReqHeartBeat int32 = 233	// 请求 心跳
	// =========== base req msg : 1000 ~ 1100
	ReqClientLogin int32 = 1000
	ReqSetName     int32 = 1010
	ReqPlayerInfos	int32  = 1011 // 请求用户信息
	ReqUpdateFlag int32 = 1012	// 更新红点
	ReqClientFeedback int32 = 1013	// 客户端错误反馈
	ReqSetCountry int32 = 1014	// 设置 国家名称
	ReqResetIcon int32 = 1015	// 重置url
	ReqAccBind int32 = 1016		// 游客绑定账号
	ReqGetUrl  int32 = 1017		// 获取链接

	ReqSignInfo int32 = 1050 //请求签到静态数据
	ReqSign     int32 = 1051 //请求签到操作
	ReqUpdatePanelTag int32 = 1052	// 客户端更新界面标识

	ReqUserRank int32 = 1250	// 请求 玩家排行

	ReqVipInfo int32 = 1280 //请求 vip信息
	ReqVipReceive int32 = 1281 //请求 vip礼品领取

	ReqChestOp  int32 = 2000
	ReqChestBuy int32 = 2001
	ReqChestFree int32 = 2002 //领取免费宝箱
	ReqChestUpgrade int32 = 2003	// 请求 升级开锁匠

	ReqHeroFormation int32 = 2100
	ReqHeroCreate    int32 = 2105 //请求英雄创建
	ReqHeroLvUp      int32 = 2110 //请求英雄升级
	ReqHeroSkillLvUp int32 = 2111 //请求英雄技能审计
	ReqFragFusion    int32 = 2112 // 请求 碎片熔炼
	ReqFusionRule	int32 = 2113	// 请求 熔炼规则
	ReqHeroTalentUp	int32 = 2114	// 请求 英雄天赋升级
	ReqHeroTalentAwaken int32 = 2115// 请求 英雄天赋觉醒
	ReqHeroFormationChange int32 = 2116// 请求 当前阵容替换
	ReqHeroQualityUpdate int32 = 2117	// 请求 英雄合成标识
	ReqHeroTalentReset int32 = 2118	// 请求 天赋重置

	ReqReelFormation int32 = 2150 //请求设置卷轴阵容
	ReqRecFamilyExp int32 = 2151 // 请求 加家族经验

	ReqShopInfo int32 = 2500 //请求商品信息
	ReqShopBuy int32 = 2501 //请求商品购买
	ReqVitBuy int32 = 2502	// 请求 体力购买
	ReqExchangeItem int32 = 2503	// 请求 兑换
	ReqMysticStore int32 = 2504	// 请求 神秘商店
	ReqBuyMysticStore int32 = 2505	// 请求 神秘商店购买
	ReqLeaveMysticStore int32 = 2506	// 请求 离开神秘商店
	ReqAltar int32 = 2507	// 请求 祭坛
	ReqBuyAltar int32 = 2508	// 请求 购买祭坛
	ReqScoreShopUp int32 = 2509	// 请求 积分商店更新
	ReqScoreShopBuy int32 = 2510// 请求 积分商店购买

	ReqChatInfo int32 = 2551;//请求聊天历史数据
	ReqChatSend int32 = 2552;//请求聊天数据发送
	ReqBattleChatSend int32 = 2554	// 请求 竞技场聊天

	ReqMailList          int32 = 5000 //请求邮件数据
	ReqMailRead			 int32 = 5002 // 请求 邮件读取
	ReqMailReceive		 int32 = 5003 // 请求 邮件领取
	ReqMailDelete		 int32 = 5004 // 请求 邮件删除

	ReqInvite 			 int32 = 5100  // 请求 邀请码
	ReqDataActivity		 int32 = 5101  // 请求活动静态数据
	ReqActRewardRec		 int32 = 5102  // 请求 活动奖励领取
	ReqDraw				 int32 = 5104  // 请求 抽卡
	ReqRecDraw			 int32 = 5105  // 请求 领取抽卡积分奖励
	ReqPrizeWheel		 int32 = 5107  // 请求 大转盘
	ReqContinueChargeRec int32 = 5111  // 请求 连续充值奖励领取
	ReqChargeSpineShowData int32 = 5112 // 请求 累充spine表现数据
	ReqChapterInfo       int32 = 6000 //请求关卡静态信息
	ReqChapterNodePickUp int32 = 6005 //请求关卡奖励拾取

	ReqChapterBattleSet int32 = 6010 //请求设置章节战斗数据
	ReqChapterBattleReset int32 = 6011 //客户端请求关卡战斗重置

	ReqChapterNodeState int32 = 6020 //客户端请求设置关卡状态
	ReqChapterState     int32 = 6021 //客户端请求章节通关
	ReqChapterMopUp		int32 = 6022 //客户端请求副本稍等
	ReqAdventureChoose int32 = 6024	// 请求 奇遇事件

	ReqTaskReceive        int32 = 6300 //客户端请求任务领取
	ReqGetDailyTask       int32 = 6301
	ReqCliRewardRece	  int32 = 6303
	ReqTaskScoreReceive   int32 = 6304
	ReqAchievementData    int32 = 6350 //客户端请求成就数据
	ReqAchievementReceive int32 = 6351 //客户端请求成就领取
	ReqSetAchieve		  int32 = 6352 // 请求设置成就

	ReqEquipLvUp   int32 = 6500 //客户端请求装备强化
	ReqEquipWear   int32 = 6501 //客户端请求装备装备
	ReqEquipDel    int32 = 6502 //客户端请求装备删除
	ReqEquipUnWear int32 = 6503 //客户端请求脱下装备
	ReqEquipSell   int32 = 6504 //客户端请求装备出售
	ReqEquipEvolution int32 = 6505// 请求 装备进化
	ReqClickEquip  int32 = 6506	// 请求 点击新装备
	ReqEquipLock	int32 = 6507	// 请求 上下锁
	ReqEquipSort 	int32 = 6508	// 请求 排序
	ReqEquipLvAttr  int32 = 6509	// 请求 装备等级数据

	ReqTalentInfo int32 = 6600 //客户端请求天赋信息
	ReqTalentLvUp int32 = 6601 //客户端请求天赋升级
	ReqChangeCareer int32 = 6602	// 请求 转职
	ReqTalentReset int32 = 6603	// 请求 重置天赋

	//好友相关
	ReqFriendInfo		int32 = 6700//客户端请求好友数据
	ReqFriendApply		int32 = 6701//客户端请求加为好友
	ReqFriendApplyRet	int32 = 6702//客户端好友申请反馈（同意 / 拒绝）
	ReqFriendFind		int32 = 6703//客户端请求好友查找
	ReqFriendDel		int32 = 6704//客户端请求好友删除
	ReqFriendGiftSend	int32 = 6705//客户端请求礼品赠送
	ReqFriendGiftRecv	int32 = 6706//客户端请求礼品领取
	ReqChestHelp 		int32 = 6707// 请求 宝箱助力
	ReqClickShare		int32 = 6708// 请求 点击分享
	ReqShare int32 = 6709	// 请求 分享
	ReqFriendRecomm int32 = 6710	// 请求 好友推荐
	ReqPlayShareGame int32 = 6711	// 请求 玩分享小游戏
	ReqGetFBFriends  int32 = 6712	// 请求 获取FB好友数据
	ReqSetChannelFriends int32 = 6713	// 请求 记录好友数据
	ReqWatchAdv int32 = 6714	// 请求 看广告
	ReqFriendInviteInfo int32 = 6715	// 请求 好友邀请数据
	ReqRecvInvite int32 = 6716	// 请求 领取邀请奖励

	//世界boss
	ReqBossInfo int32 = 8001; //客户端请求世界boss 信息
	ReqHitBoss int32 = 8002	// 请求 攻击boss
	ReqGetBossRank int32 = 8003	// 请求 获取世界boss排行
	ReqRefreshCD int32 = 8004	// 请求 刷新攻打boss冷却时间
	ReqBossInfoUpdate int32 = 8005	// 请求 更新boss信息
	ReqBossRewardInfo int32 = 8006	// 请求 boss奖励信息

	//竞技场
	ReqArenaInfo  int32 = 8501; //客户端请求竞技场数据
	ReqArenaMatch int32 = 8502; //客户端请求竞技场匹配
	ReqArenaRound int32 = 8503; //响应 客户端请求回合结束
	ReqArenaStart int32 = 8504; //响应 客户端战斗开始
	ReqArenaExit  int32 = 8505; //响应 客户端离开竞技场
	ReqHonorShop int32 = 8506	// 请求 荣耀商店
	ReqArenaGenMonster int32 = 8507; // 请求 请求竞技场
	ReqArenaRoundEnd int32 = 8508// 请求 战斗结束（客户端表现结束）
	ReqRecArenaReward int32 = 8509	// 请求 竞技场奖励领取
	ReqArenaRank int32 = 8510	// 请求 竞技场排行
	ReqHonorBuy int32 = 8511	// 请求 荣耀商店购买
	ReqArenaReconn int32 = 8512	// 请求 竞技场断线重连
	ReqArenaReconnStart int32 = 8513// 请求 竞技场断线重连开始
	ReqRecvArenaBox int32 = 8514	// 请求 竞技场宝箱领取
	ReqArenaOpenTime int32 = 8515	// 请求 竞技场开放时间
	// 矿战
	ReqGetMineInfo      int32 = 8701 // 请求 矿山信息
	ReqSaveMineHeroes   int32 = 8702 // 请求 保存守矿阵容
	ReqRefreshMineEnemy int32 = 8703 // 请求 刷新矿战对手
	ReqBattleResult     int32 = 8704 // 请求 上传矿战结果
	ReqMineUpgrade      int32 = 8705 // 请求 矿山升级
	ReqPlunderMine      int32 = 8706 // 请求 掠夺矿场
	ReqMineCollect      int32 = 8707 // 请求 矿山收集
	ReqMineBattleInfoOp int32 = 8708 // 请求 矿山情报操作
	ReqMining			int32 = 8709	// 请求 矿山开采

	// 小游戏
	ReqMiniGameInfo	int32 = 8751	// 请求 获取小游戏数据
	ReqMiniGameEnter int32 = 8752	// 请求 小游戏进入
	ReqMiniGameResult int32 = 8753	// 请求 上报小游戏结果
	
	// 支付
	ReqPayOrder        int32 = 8801; // 请求 订单支付
	ReqPayOrderCheck   int32 = 8802; // 请求 订单校验
	ReqRecChargeReward int32 = 8803  // 请求 领取充值奖励
	ReqChargeRewardInfo int32 = 8804	// 请求 充值奖励信息
	ReqPayShopShow int32 = 8805 // 请求 支付商品展示
)

const (
	// 服务器间通讯协议
	RespCreateUser int32 = 51
	RespMailUpdate int32 = 52
	RespNewMailPush int32 = 53
	RespKillOff int32 = 54

	RespIssueCode				int32 = 10000
	RespPlayerUpdateBase		int32 = 10001
	RespPlayerUpdateHero		int32 = 10002
	RespPlayerUpdateTask		int32 = 10003
	RespPlayerUpdateEquip		int32 = 10004
	RespPlayerUpdateReel		int32 = 10005
	RespPlayerMailMail			int32 = 10006
	RespPlayerUpdateAchieve		int32 = 10007
	RespPlayerUpdateArena		int32 = 10008
	RespPlayerUpdateChestFree	int32 = 10009
	RespPlayerUpdateFriend		int32 = 10010
	RespPlayerUpdateVip			int32 = 10011
	RespPlayerUpdateItem 		int32 = 10012
	RespPlayerUpdateActivity    int32 = 10013

	RespPlayerLvUp				int32 = 10050 //主角升級消息包

	RespGift					int32 = 1232 //响应领取礼包
	RespHeartBeat			    int32 = 1233	// 响应 心跳
	// =========== base resp msg : 11000 ~ 11100
	RespClientLogin int32 = 11000
	RespSetName     int32 = 11010 //回复设置昵称
	RespPlayerInfos	int32  = 11011 // 响应用户信息
	RespSetCountry	int32 = 11014	// 响应设置国家信息
	RespAccBinding  int32 = 11016	// 响应游客账号绑定
	RespGetUrl		int32 = 11017	// 响应获取url
	RespSignInfo int32 = 11050 //请求签到静态数据
	RespSign     int32 = 11051 //请求签到操作

	RespPlayerDataFlag        int32 = 11500
	RespPlayerDataBase        int32 = 11501
	RespPlayerDataHero        int32 = 11502
	RespPlayerDataBattle      int32 = 11503
	RespPlayerDataAchievement int32 = 11504
	RespPlayerDataTask        int32 = 11505
	RespPlayerDataEquip       int32 = 11506
	RespPlayerDataActivity	  int32 = 11507
	RespPlayerDataItem		  int32 = 11508

	RespUserRank int32 = 11250	// 响应 玩家排行

	RespVipInfo int32 = 11280 //响应 vip信息
	RespVipReceive int32 = 11281 //响应 vip礼品领取

	RespChestOp  int32 = 12000 //响应 请求宝箱操作
	RespChestBuy int32 = 12001 //响应 请求宝箱购买
	RespChestFree int32 = 12002 //响应 领取免费宝箱
	RespChestUpgrade int32 = 12003// 响应 升级开锁匠

	RespHeroFormation int32 = 12100 //响应 英雄设置阵容
	RespHeroCreate    int32 = 12105 //响应 英雄创建
	RespHeroLvUp      int32 = 12110 //响应 英雄升级
	RespHeroSkillLvUp int32 = 12111 //响应 英雄技能升级
	RespFragRefine    int32 = 12112 // 响应 碎片熔炼
	RespRefineRule	  int32 = 12113 // 响应 熔炼规则
	RespHeroTalentUp	int32 = 12114	// 响应 英雄天赋升级
	RespReelFormation int32 = 12150 //响应 设置卷轴阵容
	RespRecFamilyExp int32 = 12151 // 响应 加家族经验
	//商店相关
	RespShopInfo	int32 = 12500 //响应商店信息
	RespShopBuy 	int32 = 12501//请求商店购买
	RespVitBuy 		int32 = 12502	// 响应 体力购买
	RespExchangeItem int32 = 12503	// 响应 兑换
	RespMysticStore int32 = 12504	// 响应 神秘商店
	RespBuyMysticStore int32 = 12505	// 响应 神秘商店购买
	RespAltar int32 = 12507	// 响应 祭坛
	RespBuyAltar int32 = 12508	// 响应 购买祭坛
	RespScoreShopUp int32 = 12509	// 响应 积分商店更新
	RespScoreShopBuy int32 = 12510// 响应 积分商店购买
	RespUpdatePlayerDrawInfo int32 = 12511// 更新 玩家抽卡数据

	//聊天相关
	RespMarquee		int32 = 12550//响应系统通知
	RespChatInfo	int32 = 12551//响应聊天历史数据
	RespChatSend	int32 = 12552//响应聊天数据发送
	RespChatPriSend int32 = 12553//响应私聊聊数据
	RespBattleChatSend int32 = 12554// 响应 竞技场聊天
	RespUpdatePlayerChatInfo int32 = 12555// 更新 玩家聊天信息

	RespMailList int32 = 15000 //回复邮件数据
	PushMail     int32 = 15001 // 推送新邮件
	RespMailRead int32 = 15002	// 响应 邮件读取
	RespMailReceive int32 = 15003	// 响应 邮件领取
	RespMailDelete int32 = 15004	// 响应 邮件删除
	RespInvite int32 = 15100		// 响应 邀请码
	RespDataActivity int32 = 15101	// 响应 活动数据
	RespActRewardRec int32 = 15102	// 响应 活动奖励领取
	RespDraw		 int32 = 15104	// 响应 限时抽卡
	RespRecDraw		 int32 = 15105	// 响应 领取抽卡积分奖励
	RespGrowthCapital int32 = 15106	// 响应 成长基金
	RespPrizeWheel	 int32 = 15107	// 响应 大转盘
	RespDataActivityUpdate int32 = 15108	// 响应 
	RespUpdatePlayerPrivileges int32 = 15109	// 更新玩家特权
	RespContinueChargeRec int32 = 15111;	// 连续充值奖励领取
	RespChargeSpineShowData int32 = 15112 // 累充spine表现数据

	RespChapterInfo       int32 = 16000 //回复关卡静态数据
	RespChapterNodePickUp int32 = 16005 //回复章节关卡拾取
	RespChapterBattleSet int32 = 16010	// 响应 关卡数据设置
	RespChapterNodeState	int32 = 16020 //响应客户端请求设置关卡状态
	RespChapterState		int32 = 16021 //响应客户端请求章节通关
	RespChapterMopUp		int32 = 16022 //响应客户端章节扫荡
	PushTableUpdate			int32 = 16023	// 推送 沙盘数据更新（替换阵容后的push）
	
	RespTaskReceive int32 = 16300 //客户端请求任务领取
	RespGetDailyTask int32 = 16301	// 响应 获取每日任务
	RespDailyTaskUpdate int32 = 16302	// 更新每日任务进度
	RespCliRewardRece int32 = 16303
	RespTaskScoreRewardReceive int32 = 16304
	RespAdventureChoose int32 = 16024	// 请求 奇遇事件

	RespAchievementData    int32 = 16350 //响应 成就数据
	RespAchievementReceive int32 = 16351 //客户端请求成就领取

	RespEquipLvUp int32 = 16500 //响应客户端装备强化

	RespEquipWear   int32 = 16501 //响应 客户端请求穿装备
	RespEquipDel    int32 = 16502 //响应 客户端请求装备丢弃
	RespEquipUnWear int32 = 16503 //响应 客户端请求脱下装备
	RespEquipSell   int32 = 16504 //响应 客户端请求装备出售
	RespEquipEvolution int32 = 16505 // 响应 装备进化
	RespEquipLvAttr int32 = 16509	// 响应 装备等级数据

	RespTalentInfo int32 = 16600 //客户端请求天赋信息
	RespTalentLvUp int32 = 16601 //客户端请求天赋升级
	RespChangeCareer int32 = 16602	// 响应 转职
	RespTalentReset int32 = 16603	// 响应 重置天赋

//好友相关
	RespFriendInfo		int32 = 16700//客户端请求好友数据
	RespFriendApply		int32 = 16701//客户端请求加为好友
	RespFriendApplyRet	int32 = 16702//客户端响应好友
	RespFriendFind		int32 = 16703//客户端响应好友查找
	RespFriendDel		int32 = 16704//客户端请求好友删除
	RespFriendGiftSend	int32 = 16705//客户端请求礼品发送
	RespFriendGiftRecv  int32 = 16706// 响应 礼品领取
	RespChestHelp		int32 = 16707// 响应 宝箱助力
	RespClickShare 		int32 = 16708// 响应 点击分享
	RespShare  int32 = 16709	// 响应分享
	RespFriendRecomm int32 = 16710	// 响应 好友推荐
	RespPlayShareGame int32 = 16711	// 响应 玩分享小游戏
	RespGetFBFriends  int32 = 16712	// 响应 获取FB好友数据
	RespWatchAdv int32 = 16714	// 相应 观看广告
	RespFriendInvite int32 = 16715	// 请求 好友邀请数据
	RespRecvInvite int32 = 16716	// 请求 领取邀请奖励
//世界boss
	RespBossInfo int32 = 18001; //客户端请求世界boss 信息
	RespHitBoss int32 = 18002	// 响应 攻击boss
	RespGetBossRank int32 = 18003	// 响应 获取世界boss排行
	RespRefreshCD int32 = 18004	// 响应 刷新攻打boss冷却时间
	RespBossInfoUpdate int32 = 18005	// 响应 更新boss信息
	RespBossRewardInfo int32 = 18006	// 响应 boss奖励信息
	RespBossPlayerInfo int32 = 18007	// 响应 boss玩家信息
	//竞技场
	RespArenaInfo  int32 = 18501; //客户端请求竞技场数据
	RespArenaMatch int32 = 18502; //客户端请求竞技场匹配
	RespArenaRound int32 = 18503; //响应 客户端请求回合结束
	RespArenaStart int32 = 18504; //响应 客户端战斗开始
	RespArenaEnd   int32 = 18505; //响应 客户端战斗结束
	RespHonorShop int32 = 18506	// 响应 荣耀商店
	FwArenaGenMonster int32 = 18507;// 转发 怪物升级
	RespArenaRoundEnd int32 = 18508	// 响应 战斗结束（客户端表现结束）
	RespRecArenaReward int32 = 18509// 响应 竞技场奖励领取
	RespArenaRank int32 = 18510		// 响应 竞技场排行
	RespHonorBuy int32 = 18511		// 响应 荣耀商店购买
	RespArenaReconn int32 = 18512	// 响应 竞技场断线重连
	RespRecvArenaBox int32 = 18514	// 响应 竞技场宝箱领取
	RespArenaOpenTime int32 = 18515	// 响应 竞技场开放时间
	// 矿战
	RespGetMineInfo      int32 = 18701 // 响应 矿山信息
	RespSaveMineHeroes   int32 = 18702 // 响应 保存守矿阵容
	RespRefreshMineEnemy int32 = 18703 // 响应 刷新矿战对手
	RespBattleResult     int32 = 18704 // 响应 矿战结果
	RespMineUpgrade      int32 = 18705 // 响应 矿山升级
	RespPlunderMine      int32 = 18706 // 响应 矿山掠夺
	RespMineCollect      int32 = 18707 // 响应 矿山收集
	RespMineBattleInfoOp int32 = 18708 // 响应 矿山情报操作
	RespMining			int32 = 18709	// 响应 矿山开采
	RespMineInfo		int32 = 18710	// 响应 矿山情报

	// 小游戏
	RespMiniGameInfo	int32 = 18751	// 响应 获取小游戏数据
	RespMiniGameEnter int32 = 18752	// 响应 小游戏进入
	RespMiniGameResult int32 = 18753	// 响应 上报小游戏结果
	RespMiniGameUpdate int32 = 18754	// 更新 小游戏数据
	
	// 支付
	RespPayOrder int32 = 18801;	// 响应 订单支付
	RespPayOrderCheck int32 = 18802;	// 响应 订单校验
	RespRecChargeReward int32 = 18803	// 响应 领取充值奖励
	RespChargeRewardInfo int32 = 18804	// 响应 充值奖励信息
)
