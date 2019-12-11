package def

type ErrMsg struct {
	ErrID    int32
	UserID   int64
	SocketID int32
}

type UserMail struct {
	MailID []int64
	UserID int64
}

type GetUserMailArg struct {
	UserID  int64
	MailIDs []int64
}

type PushMailArg struct {
	MailID  int64
	UserIDs []int64
}

type ItemBase struct {
	Type int32
	ID   int32
	Num  int32
}

type LogLogin struct {
	Time   string `json:"t"` // 日志记录时间
	UserID int64  `json:"uid"`
	//Name    string //用户昵称
	Lv       int32 `json:"lv"`       //等级
	Gold     int32 `json:"gold"`     //金币
	Diamond  int32 `json:"diamond"`  //钻石
	RegTime  int64 `json:"reg"`      // 注册时间
	Vit      int32 `json:"vit"`      // 体力
	Platform int32 `json:"platform"` // 登陆平台
}

type LogLogout struct {
	Time      string `json:"t"`     // 日志记录时间
	LoginTime int64  `json:"login"` // 登录时间
	UserID    int64  `json:"uid"`
	//Name    string //用户昵称
	Lv      int32 `json:"lv"`      //等级
	Gold    int32 `json:"gold"`    //金币
	Diamond int32 `json:"diamond"` //钻石
	RegTime int64 `json:"reg"`     // 注册时间
	Vit     int32 `json:"vit"`     // 体力
}

type LogGold struct {
	//货币日志
	Time   string `json:"t"`   //记录时间
	UserID int64  `json:"uid"` //用户ID
	//Name          string `json:""`//用户昵称
	OpType        int32 `json:"type"` //操作类型
	CurrencyType  int32 // 货币类型 1：金币 2：钻石
	CurrencyTotal int32 `json:"total"` //货币总额
	Value         int32 `json:"value"` //变动数额
}

type LogDiamond struct {
	Time   string `json:"t"`   //记录时间
	UserID int64  `json:"uid"` //用户ID
	//Name          string `json:""`//用户昵称
	OpType int32 `json:"type"` //操作类型
	//CurrencyType  int32  // 货币类型 1：金币 2：钻石
	CurrencyTotal int32 `json:"total"` //货币总额
	Value         int32 `json:"value"` //变动数额
}

type LogOnline struct {
	Time      string `json:"t"`
	OnlineNum int32  `json:"num"` // 在线用户数
}

// sql cmd failed log
type LogSqlCmd struct {
	Time   string `json:"t"`    //记录时间
	UserID int64  `json:"uid"`  //用户ID
	SqlCmd int32  `json:"cmd"`  // sql 命令
	Args   string `json:"args"` // sql 参数
}

// 客户端反馈
type LogClientFeedback struct {
	Time   string `json:"t"`
	UserID int64  `json:"uid"`
	Type   int32  `json:"type"`
	Msg    string `json:"msg"`
	Count  int32  `json:"count"`
}

//  玩家通关章节记录
type LogChapter struct {
	Time      string `json:"t"`
	UserID    int64  `json:"uid"`
	ChapterID int32  `json:"cId"`
	UseReel   bool   `json:"useReel"` // 是否使用卷轴
	RegTime   int64  `json:"regTime"` // 注册时间
}

type LogCharge struct {
	Time       string  `json:"t"`
	UserID     int64   `json:"uid"`
	Amount     float32 `json:"amount"` // 充值金额
	ShopID     int32   `json:"shop"`
	Platform   int32   `json:"platform"` // 平台
	RegTime    int64   `json:"regTime"`
	Lv         int32   `json:"lv"`
	MaxChapter int32   `json:"maxCid"`
	Country    int32   `json:"country"`
	ChargeSum  int32   `json:"chargeSum"`
	Formation  string  `json:"formation"` // 上阵家族
	Equip      string  `json:"equip"`     // 身上装备
}

type LogRefine struct {
	Time      string `json:"t"`
	UserID    int64  `json:"uid"`
	RefineGet string `json:"getCard"` // 熔炼出的怪物 key:家族id value:数量
	//Cost      int32	`json:"cost"`
	UseCard string `json:"useCard"`
}

type LogAct struct {
	Time string `json:"t"`
	Act  int    `json:"act"`
}

type LogArena struct {
	Time       string `json:"t"`
	BeginTime  int64  `json:"begin"` // 竞技场开始时间
	EndTime    int64  `json:"end"`   // 结束时间
	PlayerRed  *LogArenaPlayer
	PlayerBlue *LogArenaPlayer
	RoundRet   []int32
	Victor     int64 // 胜方ID
}

type LogArenaPlayer struct {
	ID int64
	//Name      string
	BefScore  int32
	AftScore  int32
	BefGold   int32
	AftGold   int32
	AftRank   int32
	RegTime   int64
	Formation []*LogMonster
}

type LogMonster struct {
	FamilyID int32
	SkillLv  []int32
	TalentLv []int32
	Lv       int32
}

type LogGMOp struct {
	Time   string
	Addr   string
	User   string
	OpType string
	Msg    string
}

type LogCard struct {
	Time     string `json:"t"`
	UserID   int64  `json:"uid"`
	FamilyID int32  `json:"fId"`    // 家族ID
	OpType   int32  `json:"opType"` //操作类型
	Value    int32  `json:"value"`  // 改变值
	Total    int32  `json:"total"`  // 总
}

type LogShopBuy struct {
	Time    string  `json:"t"`
	UserID  int64   `json:"uid"`
	ShopID  int32   `json:"shopId"`
	Price   float32 `json:"price"`
	CurType int32   `json:"curType"` // 货币类型
}

type LogRegUser struct {
	Time      string `json:"t"`
	UserID    int64  `json:"uid"`
	Account   string `json:"user"`
	ChannelID int32  `json:"channel"`
	Platform  int32  `json:"platform"`
}

type LogArenaChest struct {
	Time    string `json:"t"`
	UserID  int64  `json:"uid"`
	ChestID int32  `json:"chest"`
	Star    int32  `json:"star"`
}

type LogEquip struct {
	Time          string  `json:"t"`
	UserID        int64   `json:"uid"`
	Type          int32   `json:"type"`    // 类型 1：获得 2：强化 3:出售
	SubType       int32   `json:"subType"` // 获得途径
	EquipID       []int32 `json:"equip"`   //
	EquipBaseID   []int32 `json:"equipBase"`
	Count         int32   `json:"count"`         // 计数
	UpEquipID     int32   `json:"upEquipId"`     // 强化的装备
	UpEquipBaseID int32   `json:"upEquipBaseId"` // 强化的装备
}

type LogReel struct {
	Time    string  `json:"t"`
	UserID  int64   `json:"uid"`
	Type    int32   `json:"type"`    // 类型 1：获取 2：使用
	ReelIDs []int32 `json:"reel"`    // 使用卷轴ID
	Num     []int32 `json:"num"`     // 对应数量
	StageID int32   `json:"stageId"` // 使用的关卡
}

type LogHitBoss struct {
	Time     string `json:"t"`
	UserID   int64  `json:"uid"`
	Damage   int64  `json:"damage"`
	BossTime string `json:"bossTime"`
	HitNum   int32  `json:"hitNum"`
}

type LogShare struct {
	Time      string
	UserID    int64
	ShareType string
}

type LogArenaShop struct {
	Time   string
	UserID int64
	Goods  string
	Price  int32
	Times  int32 // 剩余购买次数
	Score  int32 // 竞技场积分
}

type LogExp struct {
	Time   string
	UserID int64
	Exp    int32
}

type LogClientDot struct {
	Time   string
	UserID int64
	Key    int32
	Arg1   int32
	Arg2   int32
}

type ArgsFcm struct {
	Token    []string
	Language []int32
}
