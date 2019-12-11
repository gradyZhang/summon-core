package def

// 错误码

// 基础类型 1-20
const (
	ERR_GOLD_NOT_ENOUGH                       int32 = iota + 1 // 金币不足 1
	ERR_VIT_NOT_ENOUGH                                         // 体力不足 2
	ERR_DIAMOND_NOT_ENOUGH                                     // 钻石不足 3
	ERR_HEROCLIP_NOT_ENOUGH                                    // 怪物家族碎片不足 4
	ERR_EQUIP_NOT_EXIST                                        // 装备不存在 5
	ERR_HERO_KEEP_OUT                                          // 怪物保留数量超出 6
	ERR_TALENT_NOT_EXIST                                       // 天赋不存在 7
	ERR_CHAPTER_NOT_EXIST                                      // 章节不存在 8
	ERR_CHAPTER_KEY_NOT_ENOUGH                                 // 章节钥匙不足 9
	ERR_CHAPTER_STAGE_NOT_EXIST                                // 章节关卡不存在 10
	ERR_CHAPTER_STATE_PASSED                                   // 章节关卡已通关 11
	ERR_MINE_NOT_FOUND                                         // 矿山信息未找到 12
	ERR_BADGE_NOT_ENOUGH                                       // 徽章不足 13
	ERR_REEL_NOT_EXIST                                         // 卷轴不存在 14
	ERR_REEL_NOT_ENOUGH                                        // 卷轴数量不足 15
	ERR_ACHIEVE_NOT_EXIST                                      // 成就不存在 16
	ERR_ACHIEVE_NOT_ENOUGH                                     // 成就条件不满足 17
	ERR_KICK_SAME_ACC                                          // 同号在线踢人 18
	ERR_CHEST_MAX_UNLOCK                                       // 宝箱同时解锁已达上限 19
	ERR_EQUIP_WEARS                                            // 装备穿着呢 20
	ERR_MAIL_NULL                                              // 邮件未找到 21
	ERR_MAIL_NOT_RECEIVE                                       // 有未领取邮件 22
	ERR_CHAPTER_STATE_UNPASSED                                 // 章节关卡未通关 23
	ERR_MAIL_RECEIVED                                          // 邮件已领取 24
	ERR_CANNOT_RECEIVE                                         // 不可领取 25
	ERR_TASK_NOT_FOUND                                         // 任务未找到26
	ERR_VIP_NOT_EXIST                                          // vip 不存在 27
	ERR_VIP_RECEIVED                                           // vip 礼包 已领取 28
	ERR_CHANGE_CAREER_FAILED                                   // 转职失败 29
	ERR_TALENT_UPGRADE_FAILED                                  // 天赋升级失败 30
	ERR_TALENTPOINT_NOT_ENOUGH                                 // 天赋点不足 31
	ERR_TALENT_SKILL_NOT_EXIST                                 // 天赋技能未找到 32
	ERR_TALENT_SKILL_LV_OUTOFF                                 // 天赋技能等级越界 33
	ERR_PRE_CAREER_USE_TALENTPOINT_NOT_ENOUGH                  // 上个职业使用的天赋点不够 34
	ERR_FRIEND_NOT_FOUND                                       // 好友数据未找到 35
	ERR_SIGNIN_REPEAT                                          // 已签到 36
	ERR_PICKREWARD_PICKED                                      // 已领取关卡掉落奖励 37
	ERR_PICKREWARD_NOTFOUND                                    // 关卡掉落奖励未找到 38
	ERR_ADVENTURE_NOTFOUND                                     // 奇遇事件未找到 39
	ERR_CAREER_FAILURE_BUY                                     // 职业不符，不能购买 40
	ERR_CAREER_FAILURE_WEAR                                    // 职业不符，不能穿 41
	ERR_CHAT_CENTENT_MAX                                       // 聊天内容 超出上限 42
	ERR_VIT_BUY_TIMES_NOTENOUGH                                // 体力购买次数不足 43
	ERR_GIFT_SEND_NUM_FULL                                     // 好友赠送次数已满 44
	ERR_GIFT_SENDED                                            // 好友已赠送过 45
	ERR_RECV_NUM_NOTENOUGH                                     // 领取好友赠送礼物次数不足 46
	ERR_SHOP_NOTFOUND                                          // 商品未找到 47
	ERR_TIMES_NOT_ENOUGH                                       // 材料不足 48
	ERR_CHEST_NOT_EXIST                                        // 宝箱不存在 49
)

// 怪物 50 - 100
const (
	ERR_HERO_NULL       int32 = iota + 50 // 英雄为空
	ERR_CLIP_NOT_ENOUGH                   // 碎片不足 51
	ERR_REFINE_LIMIT                      // 每日指定熔炼次数不足 52
	ERR_DESIGN_FAILED                     // 指定熔炼失败 53
	ERR_REFINE_FAILED                     // 熔炼失败 54

	ERR_HERO_TALENT_UNLOCK       // 天赋未解锁 55
	ERR_HERO_TALENT_MAXLV        // 天赋达到最大 56
	ERR_FORMATION_OUTOFRANGE     // 阵容索引越界 57
	ERR_HERO_SKILL_MAXLV         // 被动技能等级达上限 58
	ERR_LEADER_NOT_ENOUGH        // 领导力不足 59
	ERR_HERO_IN_OTHER_FORMATION  // 该英雄在其他阵容里 60
	ERR_QUALITY_NOT_SAME         // 熔炼的品质不一样 61
	ERR_HERO_PRO_NOT_ENOUGH      // 英雄熟练度不足 62
	ERR_HERO_FORMATION_FAILED    // 替换阵容失败 63
	ERR_HERO_CREATED             // 该英雄已创建 64
	ERR_SKILL_NOTFOUND           // 被动技能未找到 65
	ERR_HERO_TALENT_NOTFOUND     // 英雄天赋为找到 66
	ERR_LEADER_NOT_ENOUGH2       // 领导力不足2 67
)

// 世界 100 - 150
const (
	ERR_COLLECT_FAILED          int32 = iota + 100 // 矿场收集失败 100
	ERR_MINING_FAILED                              // 矿场开采失败 101
	ERR_MINE_FAILED                                //				102
	ERR_MINE_PLUNDER_NOT_ENOUGH                    // 掠夺次数不足 103
	ERR_MINE_LV_MAX                                // 矿山等级已达最大 104
	ERR_CHEST_MAX                                  // 宝箱已满 105
	ERR_HONOR_NOT_ENOUGH                           // 荣耀点不足 106
	ERR_BOSS_DEAD                                  // boss死了 107
	ERR_ARENA_ROOM_NOTFOUND                        // 竞技场数据未找到 108
	ERR_HONOR_BUYTIMES_NOT_ENOUGH					// 荣耀商品购买次数不足 109
	ERR_ARENA_NOT_OPEN								// 竞技场未开放 110
	ERR_BOSS_HIT_TIMES_NOT_ENOUGH					// 世界boss攻击次数不足 111
	ERR_CANNOT_HITBOSS								// 不可攻打boss 112
	ERR_CANNOT_CHAT									// 未改名不可聊天 113
	ERR_CHAT_USER_OFFLINE 							// 私聊用户不在线 114
	ERR_CHAT_CD  									// 私聊在CD时间内 115
)

// 小游戏副本
const (
	ERR_MINIGAME_TIMES_NOT_ENOUGH int32 = iota + 130 // 小游戏次数不足 130
)

const (
	ERR_LOCKSMITH_OUTOFF_RANGE int32 = iota + 200 // 开锁匠达到上限 200
	ERR_EQUIP_LV_MAX                              // 装备等级达上限 201
	ERR_NOT_FASHION                               // 不是时装 202
	ERR_PAY_OFF									  // 支付未开启 203
	ERR_CHAPTER_BATTLE_OUTOFF						// 已经有两个关卡在进行了 204
	ERR_KICKOFF 									// 系统踢人 205
	ERR_SHARE_VIT_MAX								// 体力分享次数达上限 206
	ERR_SHARE_STEP_MAX								// 关卡步数分享次数达上限 207
	ERR_USE_REEL_OUTOF								// 关卡内使用卷抽超上限 208
	ERR_CHAPTER_NOT_PASSED							// 关卡未通过 209
	ERR_ILLEGAL_CHARACTER							// 非法字符 210
	ERR_MONCARD_RECE_FAILED							// 月卡领取失败 211
	ERR_CHAPTER_BATTLE_INFO_NOTFOUND				// 章节战斗数据为找到 212
	ERR_MINE_BATTLE_INFO_NOT_FOUND					// 矿山情报未找到 213
	ERR_MINE_BATTLE_INFO_DID						// 矿山情报已操作 214
	ERR_CANNOT_BINDING								// 非游客登陆不可绑定账号 215
	ERR_BINDED										// 账号已绑定 216
	ERR_ACC_BINDED_EXISTS							// 绑定账号已存在 217
	ERR_ARENA_STAR_NOTENOUGH						// 星星不足 218
	ERR_EQUIP_LOCK									// 装备已锁定 219
	ERR_ACT_NOTFOUND	// 活动未找到 220
	ERR_RECEIVED		// 已领取 221
	ERR_BUY				// 已购买 222
	ERR_LV_NOTENOUGH	// 等级未达到 223
	ERR_INVITED			// 已邀请 224
	ERR_INVITE_CODE		// 邀请码未找到 225
	ERR_NO_INVITE		// 不可邀请 226
	ERR_ACT_OVER		// 活动过期 227
	ERR_SCORESHAOP_NOT_ENOUGH	//  积分商店不足 228
	ERR_ITEM_NOTENOUGH  // 道具不足 229
	ERR_TICKET_NOTENOUGH // 抽奖券不足 230
	ERR_BUY_PRE			 // 要先购买前置 231
	ERR_NO_BUY			 // 未购买 232
	ERR_EQUIP_BAG_MAX	 // 背包满了，发邮件了 233
	ERR_DRAW_TIMES_NOTENOUGH // 抽卡次数不足
)
