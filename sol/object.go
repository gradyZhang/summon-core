package sol

type UserShare struct {
	DailyChestShare int32 	// 宝箱每日分享次数
	DailyLinkCount int32 	// 每日点击链接次数
	RecRewardId int32 	// 已领取奖励的id
	Shares map[int64][]int64	// 分享的链接
}
