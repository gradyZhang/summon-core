// Code generated by protoc-gen-go.
// source: mine.proto
// DO NOT EDIT!

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// ==================== 矿战 相关 协议
//
type Mine_Player_Info_ struct {
	UserID           *int32      `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	Nickname         *string     `protobuf:"bytes,2,opt,name=Nickname" json:"Nickname,omitempty"`
	Role             *Role_Info_ `protobuf:"bytes,3,opt,name=Role" json:"Role,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Mine_Player_Info_) Reset()                    { *m = Mine_Player_Info_{} }
func (m *Mine_Player_Info_) String() string            { return proto.CompactTextString(m) }
func (*Mine_Player_Info_) ProtoMessage()               {}
func (*Mine_Player_Info_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *Mine_Player_Info_) GetUserID() int32 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Mine_Player_Info_) GetNickname() string {
	if m != nil && m.Nickname != nil {
		return *m.Nickname
	}
	return ""
}

func (m *Mine_Player_Info_) GetRole() *Role_Info_ {
	if m != nil {
		return m.Role
	}
	return nil
}

// 请求 矿山信息
type Req_Get_MineInfo_ struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_Get_MineInfo_) Reset()                    { *m = Req_Get_MineInfo_{} }
func (m *Req_Get_MineInfo_) String() string            { return proto.CompactTextString(m) }
func (*Req_Get_MineInfo_) ProtoMessage()               {}
func (*Req_Get_MineInfo_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *Req_Get_MineInfo_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

// 响应 矿山信息
type Resp_Get_MineInfo_ struct {
	UserID           *int64             `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	CurLv            *int32             `protobuf:"varint,2,opt,name=CurLv" json:"CurLv,omitempty"`
	CurOutput        *int32             `protobuf:"varint,3,opt,name=CurOutput" json:"CurOutput,omitempty"`
	BeginTime        *int64             `protobuf:"varint,4,opt,name=BeginTime" json:"BeginTime,omitempty"`
	EndTime          *int64             `protobuf:"varint,5,opt,name=EndTime" json:"EndTime,omitempty"`
	PlunderNum       *int32             `protobuf:"varint,6,opt,name=PlunderNum" json:"PlunderNum,omitempty"`
	NextPlunderTime  *int64             `protobuf:"varint,7,opt,name=NextPlunderTime" json:"NextPlunderTime,omitempty"`
	NextLvOutput     *int32             `protobuf:"varint,8,opt,name=NextLvOutput" json:"NextLvOutput,omitempty"`
	UpgradeBadge     *int32             `protobuf:"varint,9,opt,name=UpgradeBadge" json:"UpgradeBadge,omitempty"`
	Heroes           []*BattleHeroInfo_ `protobuf:"bytes,10,rep,name=Heroes" json:"Heroes,omitempty"`
	TableInfo        *TableInfo_        `protobuf:"bytes,11,opt,name=TableInfo" json:"TableInfo,omitempty"`
	Step             *int32             `protobuf:"varint,12,opt,name=Step" json:"Step,omitempty"`
	MaxPlunderNum    *int32             `protobuf:"varint,13,opt,name=MaxPlunderNum" json:"MaxPlunderNum,omitempty"`
	CollectTime      *int64             `protobuf:"varint,14,opt,name=CollectTime" json:"CollectTime,omitempty"`
	MiningCDTime     *int64             `protobuf:"varint,15,opt,name=MiningCDTime" json:"MiningCDTime,omitempty"`
	RefreshPaice     *int32             `protobuf:"varint,16,opt,name=RefreshPaice" json:"RefreshPaice,omitempty"`
	BeRob            *int32             `protobuf:"varint,17,opt,name=BeRob" json:"BeRob,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Resp_Get_MineInfo_) Reset()                    { *m = Resp_Get_MineInfo_{} }
func (m *Resp_Get_MineInfo_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Get_MineInfo_) ProtoMessage()               {}
func (*Resp_Get_MineInfo_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *Resp_Get_MineInfo_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetCurLv() int32 {
	if m != nil && m.CurLv != nil {
		return *m.CurLv
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetCurOutput() int32 {
	if m != nil && m.CurOutput != nil {
		return *m.CurOutput
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetBeginTime() int64 {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetPlunderNum() int32 {
	if m != nil && m.PlunderNum != nil {
		return *m.PlunderNum
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetNextPlunderTime() int64 {
	if m != nil && m.NextPlunderTime != nil {
		return *m.NextPlunderTime
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetNextLvOutput() int32 {
	if m != nil && m.NextLvOutput != nil {
		return *m.NextLvOutput
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetUpgradeBadge() int32 {
	if m != nil && m.UpgradeBadge != nil {
		return *m.UpgradeBadge
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetHeroes() []*BattleHeroInfo_ {
	if m != nil {
		return m.Heroes
	}
	return nil
}

func (m *Resp_Get_MineInfo_) GetTableInfo() *TableInfo_ {
	if m != nil {
		return m.TableInfo
	}
	return nil
}

func (m *Resp_Get_MineInfo_) GetStep() int32 {
	if m != nil && m.Step != nil {
		return *m.Step
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetMaxPlunderNum() int32 {
	if m != nil && m.MaxPlunderNum != nil {
		return *m.MaxPlunderNum
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetCollectTime() int64 {
	if m != nil && m.CollectTime != nil {
		return *m.CollectTime
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetMiningCDTime() int64 {
	if m != nil && m.MiningCDTime != nil {
		return *m.MiningCDTime
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetRefreshPaice() int32 {
	if m != nil && m.RefreshPaice != nil {
		return *m.RefreshPaice
	}
	return 0
}

func (m *Resp_Get_MineInfo_) GetBeRob() int32 {
	if m != nil && m.BeRob != nil {
		return *m.BeRob
	}
	return 0
}

// 请求 矿山升级
type Req_Mine_Upgrade_ struct {
	TargetLv         *int32 `protobuf:"varint,1,opt,name=TargetLv" json:"TargetLv,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_Mine_Upgrade_) Reset()                    { *m = Req_Mine_Upgrade_{} }
func (m *Req_Mine_Upgrade_) String() string            { return proto.CompactTextString(m) }
func (*Req_Mine_Upgrade_) ProtoMessage()               {}
func (*Req_Mine_Upgrade_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *Req_Mine_Upgrade_) GetTargetLv() int32 {
	if m != nil && m.TargetLv != nil {
		return *m.TargetLv
	}
	return 0
}

// 响应 矿山升级
type Resp_Mine_Upgrade_ struct {
	CurLv            *int32 `protobuf:"varint,1,opt,name=CurLv" json:"CurLv,omitempty"`
	CurOutput        *int32 `protobuf:"varint,2,opt,name=CurOutput" json:"CurOutput,omitempty"`
	NextLvOutput     *int32 `protobuf:"varint,3,opt,name=NextLvOutput" json:"NextLvOutput,omitempty"`
	UpgradeBadge     *int32 `protobuf:"varint,4,opt,name=UpgradeBadge" json:"UpgradeBadge,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Resp_Mine_Upgrade_) Reset()                    { *m = Resp_Mine_Upgrade_{} }
func (m *Resp_Mine_Upgrade_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Mine_Upgrade_) ProtoMessage()               {}
func (*Resp_Mine_Upgrade_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{4} }

func (m *Resp_Mine_Upgrade_) GetCurLv() int32 {
	if m != nil && m.CurLv != nil {
		return *m.CurLv
	}
	return 0
}

func (m *Resp_Mine_Upgrade_) GetCurOutput() int32 {
	if m != nil && m.CurOutput != nil {
		return *m.CurOutput
	}
	return 0
}

func (m *Resp_Mine_Upgrade_) GetNextLvOutput() int32 {
	if m != nil && m.NextLvOutput != nil {
		return *m.NextLvOutput
	}
	return 0
}

func (m *Resp_Mine_Upgrade_) GetUpgradeBadge() int32 {
	if m != nil && m.UpgradeBadge != nil {
		return *m.UpgradeBadge
	}
	return 0
}

// 请求 保存矿山英雄（开采）
type Req_Save_MineHeroes_ struct {
	UserID           *int64             `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	Heroes           []*BattleHeroInfo_ `protobuf:"bytes,2,rep,name=Heroes" json:"Heroes,omitempty"`
	Player           *Mine_Player_Info_ `protobuf:"bytes,3,opt,name=Player" json:"Player,omitempty"`
	MineLv           *int32             `protobuf:"varint,4,opt,name=MineLv" json:"MineLv,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Req_Save_MineHeroes_) Reset()                    { *m = Req_Save_MineHeroes_{} }
func (m *Req_Save_MineHeroes_) String() string            { return proto.CompactTextString(m) }
func (*Req_Save_MineHeroes_) ProtoMessage()               {}
func (*Req_Save_MineHeroes_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{5} }

func (m *Req_Save_MineHeroes_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Req_Save_MineHeroes_) GetHeroes() []*BattleHeroInfo_ {
	if m != nil {
		return m.Heroes
	}
	return nil
}

func (m *Req_Save_MineHeroes_) GetPlayer() *Mine_Player_Info_ {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *Req_Save_MineHeroes_) GetMineLv() int32 {
	if m != nil && m.MineLv != nil {
		return *m.MineLv
	}
	return 0
}

// 响应 保存矿山英雄（开采）
type Resp_Save_MineHeroes_ struct {
	UserID           *int64             `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	BeginTime        *int64             `protobuf:"varint,2,opt,name=BeginTime" json:"BeginTime,omitempty"`
	EndTime          *int64             `protobuf:"varint,3,opt,name=EndTime" json:"EndTime,omitempty"`
	CollectTime      *int64             `protobuf:"varint,4,opt,name=CollectTime" json:"CollectTime,omitempty"`
	Heroes           []*BattleHeroInfo_ `protobuf:"bytes,5,rep,name=Heroes" json:"Heroes,omitempty"`
	MiningCDTime     *int64             `protobuf:"varint,6,opt,name=MiningCDTime" json:"MiningCDTime,omitempty"`
	Rewards          []*Reward_Info_    `protobuf:"bytes,7,rep,name=Rewards" json:"Rewards,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Resp_Save_MineHeroes_) Reset()                    { *m = Resp_Save_MineHeroes_{} }
func (m *Resp_Save_MineHeroes_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Save_MineHeroes_) ProtoMessage()               {}
func (*Resp_Save_MineHeroes_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{6} }

func (m *Resp_Save_MineHeroes_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Save_MineHeroes_) GetBeginTime() int64 {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return 0
}

func (m *Resp_Save_MineHeroes_) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *Resp_Save_MineHeroes_) GetCollectTime() int64 {
	if m != nil && m.CollectTime != nil {
		return *m.CollectTime
	}
	return 0
}

func (m *Resp_Save_MineHeroes_) GetHeroes() []*BattleHeroInfo_ {
	if m != nil {
		return m.Heroes
	}
	return nil
}

func (m *Resp_Save_MineHeroes_) GetMiningCDTime() int64 {
	if m != nil && m.MiningCDTime != nil {
		return *m.MiningCDTime
	}
	return 0
}

func (m *Resp_Save_MineHeroes_) GetRewards() []*Reward_Info_ {
	if m != nil {
		return m.Rewards
	}
	return nil
}

// 请求 刷新矿战对手
type Req_Refresh_MineEnemy_ struct {
	UserID           *int64  `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	MineLv           *int32  `protobuf:"varint,3,opt,name=MineLv" json:"MineLv,omitempty"`
	EnemyUid         []int64 `protobuf:"varint,4,rep,name=EnemyUid" json:"EnemyUid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Req_Refresh_MineEnemy_) Reset()                    { *m = Req_Refresh_MineEnemy_{} }
func (m *Req_Refresh_MineEnemy_) String() string            { return proto.CompactTextString(m) }
func (*Req_Refresh_MineEnemy_) ProtoMessage()               {}
func (*Req_Refresh_MineEnemy_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{7} }

func (m *Req_Refresh_MineEnemy_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Req_Refresh_MineEnemy_) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *Req_Refresh_MineEnemy_) GetMineLv() int32 {
	if m != nil && m.MineLv != nil {
		return *m.MineLv
	}
	return 0
}

func (m *Req_Refresh_MineEnemy_) GetEnemyUid() []int64 {
	if m != nil {
		return m.EnemyUid
	}
	return nil
}

// 响应 刷新矿战对手
type Resp_Refresh_MineEnemy_ struct {
	UserID           *int64             `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	Name             *string            `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	MineLv           *int32             `protobuf:"varint,3,opt,name=MineLv" json:"MineLv,omitempty"`
	Output           *int32             `protobuf:"varint,4,opt,name=Output" json:"Output,omitempty"`
	GetGold          *int32             `protobuf:"varint,5,opt,name=GetGold" json:"GetGold,omitempty"`
	GetExp           *int32             `protobuf:"varint,6,opt,name=GetExp" json:"GetExp,omitempty"`
	Heroes           []*BattleHeroInfo_ `protobuf:"bytes,7,rep,name=Heroes" json:"Heroes,omitempty"`
	Role             *Role_Info_        `protobuf:"bytes,8,opt,name=Role" json:"Role,omitempty"`
	DefUserID        *int64             `protobuf:"varint,9,opt,name=DefUserID" json:"DefUserID,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Resp_Refresh_MineEnemy_) Reset()                    { *m = Resp_Refresh_MineEnemy_{} }
func (m *Resp_Refresh_MineEnemy_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Refresh_MineEnemy_) ProtoMessage()               {}
func (*Resp_Refresh_MineEnemy_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{8} }

func (m *Resp_Refresh_MineEnemy_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Refresh_MineEnemy_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Resp_Refresh_MineEnemy_) GetMineLv() int32 {
	if m != nil && m.MineLv != nil {
		return *m.MineLv
	}
	return 0
}

func (m *Resp_Refresh_MineEnemy_) GetOutput() int32 {
	if m != nil && m.Output != nil {
		return *m.Output
	}
	return 0
}

func (m *Resp_Refresh_MineEnemy_) GetGetGold() int32 {
	if m != nil && m.GetGold != nil {
		return *m.GetGold
	}
	return 0
}

func (m *Resp_Refresh_MineEnemy_) GetGetExp() int32 {
	if m != nil && m.GetExp != nil {
		return *m.GetExp
	}
	return 0
}

func (m *Resp_Refresh_MineEnemy_) GetHeroes() []*BattleHeroInfo_ {
	if m != nil {
		return m.Heroes
	}
	return nil
}

func (m *Resp_Refresh_MineEnemy_) GetRole() *Role_Info_ {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *Resp_Refresh_MineEnemy_) GetDefUserID() int64 {
	if m != nil && m.DefUserID != nil {
		return *m.DefUserID
	}
	return 0
}

// 请求 矿山收集
type Req_Mine_Collect_ struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_Mine_Collect_) Reset()                    { *m = Req_Mine_Collect_{} }
func (m *Req_Mine_Collect_) String() string            { return proto.CompactTextString(m) }
func (*Req_Mine_Collect_) ProtoMessage()               {}
func (*Req_Mine_Collect_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{9} }

func (m *Req_Mine_Collect_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

// 响应 矿山收集
type Resp_Mine_Collect_ struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	Gold             *int32 `protobuf:"varint,2,opt,name=Gold" json:"Gold,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Resp_Mine_Collect_) Reset()                    { *m = Resp_Mine_Collect_{} }
func (m *Resp_Mine_Collect_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Mine_Collect_) ProtoMessage()               {}
func (*Resp_Mine_Collect_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{10} }

func (m *Resp_Mine_Collect_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Mine_Collect_) GetGold() int32 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

// 请求 掠夺战结果
type Req_Battle_Result_ struct {
	UserID           *int64  `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	Result           *int32  `protobuf:"varint,2,opt,name=Result" json:"Result,omitempty"`
	MineLv           *int32  `protobuf:"varint,3,opt,name=MineLv" json:"MineLv,omitempty"`
	GetGold          *int32  `protobuf:"varint,4,opt,name=GetGold" json:"GetGold,omitempty"`
	GetExp           *int32  `protobuf:"varint,5,opt,name=GetExp" json:"GetExp,omitempty"`
	DefUserID        *int64  `protobuf:"varint,6,opt,name=DefUserID" json:"DefUserID,omitempty"`
	Name             *string `protobuf:"bytes,7,opt,name=Name" json:"Name,omitempty"`
	InfoID           *int32  `protobuf:"varint,8,opt,name=InfoID" json:"InfoID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Req_Battle_Result_) Reset()                    { *m = Req_Battle_Result_{} }
func (m *Req_Battle_Result_) String() string            { return proto.CompactTextString(m) }
func (*Req_Battle_Result_) ProtoMessage()               {}
func (*Req_Battle_Result_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{11} }

func (m *Req_Battle_Result_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Req_Battle_Result_) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *Req_Battle_Result_) GetMineLv() int32 {
	if m != nil && m.MineLv != nil {
		return *m.MineLv
	}
	return 0
}

func (m *Req_Battle_Result_) GetGetGold() int32 {
	if m != nil && m.GetGold != nil {
		return *m.GetGold
	}
	return 0
}

func (m *Req_Battle_Result_) GetGetExp() int32 {
	if m != nil && m.GetExp != nil {
		return *m.GetExp
	}
	return 0
}

func (m *Req_Battle_Result_) GetDefUserID() int64 {
	if m != nil && m.DefUserID != nil {
		return *m.DefUserID
	}
	return 0
}

func (m *Req_Battle_Result_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Req_Battle_Result_) GetInfoID() int32 {
	if m != nil && m.InfoID != nil {
		return *m.InfoID
	}
	return 0
}

// 响应 掠夺战结果
type Resp_Battle_Result_ struct {
	UserID           *int64          `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	GetGold          *int32          `protobuf:"varint,2,opt,name=GetGold" json:"GetGold,omitempty"`
	GetBadge         *int32          `protobuf:"varint,3,opt,name=GetBadge" json:"GetBadge,omitempty"`
	GetExp           *int32          `protobuf:"varint,4,opt,name=GetExp" json:"GetExp,omitempty"`
	MineLv           *int32          `protobuf:"varint,5,opt,name=MineLv" json:"MineLv,omitempty"`
	Rewards          []*Reward_Info_ `protobuf:"bytes,6,rep,name=Rewards" json:"Rewards,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Resp_Battle_Result_) Reset()                    { *m = Resp_Battle_Result_{} }
func (m *Resp_Battle_Result_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Battle_Result_) ProtoMessage()               {}
func (*Resp_Battle_Result_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{12} }

func (m *Resp_Battle_Result_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Battle_Result_) GetGetGold() int32 {
	if m != nil && m.GetGold != nil {
		return *m.GetGold
	}
	return 0
}

func (m *Resp_Battle_Result_) GetGetBadge() int32 {
	if m != nil && m.GetBadge != nil {
		return *m.GetBadge
	}
	return 0
}

func (m *Resp_Battle_Result_) GetGetExp() int32 {
	if m != nil && m.GetExp != nil {
		return *m.GetExp
	}
	return 0
}

func (m *Resp_Battle_Result_) GetMineLv() int32 {
	if m != nil && m.MineLv != nil {
		return *m.MineLv
	}
	return 0
}

func (m *Resp_Battle_Result_) GetRewards() []*Reward_Info_ {
	if m != nil {
		return m.Rewards
	}
	return nil
}

// 响应 掠夺
type Resp_Plunder_Mine_ struct {
	PlunderNum       *int32 `protobuf:"varint,1,opt,name=PlunderNum" json:"PlunderNum,omitempty"`
	NextPlunderTime  *int64 `protobuf:"varint,2,opt,name=NextPlunderTime" json:"NextPlunderTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Resp_Plunder_Mine_) Reset()                    { *m = Resp_Plunder_Mine_{} }
func (m *Resp_Plunder_Mine_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Plunder_Mine_) ProtoMessage()               {}
func (*Resp_Plunder_Mine_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{13} }

func (m *Resp_Plunder_Mine_) GetPlunderNum() int32 {
	if m != nil && m.PlunderNum != nil {
		return *m.PlunderNum
	}
	return 0
}

func (m *Resp_Plunder_Mine_) GetNextPlunderTime() int64 {
	if m != nil && m.NextPlunderTime != nil {
		return *m.NextPlunderTime
	}
	return 0
}

// 请求 矿山情报操作
type Req_MineInfo_Op_ struct {
	UserID           *int64 `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	InfoID           *int32 `protobuf:"varint,2,opt,name=InfoID" json:"InfoID,omitempty"`
	EnemyUid         *int64 `protobuf:"varint,3,opt,name=EnemyUid" json:"EnemyUid,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Req_MineInfo_Op_) Reset()                    { *m = Req_MineInfo_Op_{} }
func (m *Req_MineInfo_Op_) String() string            { return proto.CompactTextString(m) }
func (*Req_MineInfo_Op_) ProtoMessage()               {}
func (*Req_MineInfo_Op_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{14} }

func (m *Req_MineInfo_Op_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Req_MineInfo_Op_) GetInfoID() int32 {
	if m != nil && m.InfoID != nil {
		return *m.InfoID
	}
	return 0
}

func (m *Req_MineInfo_Op_) GetEnemyUid() int64 {
	if m != nil && m.EnemyUid != nil {
		return *m.EnemyUid
	}
	return 0
}

// 响应 反击
type Resp_MineInfo_Op_ struct {
	UserID           *int64             `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	InfoType         *int32             `protobuf:"varint,2,opt,name=InfoType" json:"InfoType,omitempty"`
	GetGold          *int32             `protobuf:"varint,3,opt,name=GetGold" json:"GetGold,omitempty"`
	GetExp           *int32             `protobuf:"varint,4,opt,name=GetExp" json:"GetExp,omitempty"`
	GetBadge         *int32             `protobuf:"varint,5,opt,name=GetBadge" json:"GetBadge,omitempty"`
	Name             *string            `protobuf:"bytes,6,opt,name=Name" json:"Name,omitempty"`
	MineLv           *int32             `protobuf:"varint,7,opt,name=MineLv" json:"MineLv,omitempty"`
	Output           *int32             `protobuf:"varint,8,opt,name=Output" json:"Output,omitempty"`
	Heroes           []*BattleHeroInfo_ `protobuf:"bytes,9,rep,name=Heroes" json:"Heroes,omitempty"`
	Role             *Role_Info_        `protobuf:"bytes,10,opt,name=Role" json:"Role,omitempty"`
	DefUserID        *int64             `protobuf:"varint,11,opt,name=DefUserID" json:"DefUserID,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Resp_MineInfo_Op_) Reset()                    { *m = Resp_MineInfo_Op_{} }
func (m *Resp_MineInfo_Op_) String() string            { return proto.CompactTextString(m) }
func (*Resp_MineInfo_Op_) ProtoMessage()               {}
func (*Resp_MineInfo_Op_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{15} }

func (m *Resp_MineInfo_Op_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_MineInfo_Op_) GetInfoType() int32 {
	if m != nil && m.InfoType != nil {
		return *m.InfoType
	}
	return 0
}

func (m *Resp_MineInfo_Op_) GetGetGold() int32 {
	if m != nil && m.GetGold != nil {
		return *m.GetGold
	}
	return 0
}

func (m *Resp_MineInfo_Op_) GetGetExp() int32 {
	if m != nil && m.GetExp != nil {
		return *m.GetExp
	}
	return 0
}

func (m *Resp_MineInfo_Op_) GetGetBadge() int32 {
	if m != nil && m.GetBadge != nil {
		return *m.GetBadge
	}
	return 0
}

func (m *Resp_MineInfo_Op_) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Resp_MineInfo_Op_) GetMineLv() int32 {
	if m != nil && m.MineLv != nil {
		return *m.MineLv
	}
	return 0
}

func (m *Resp_MineInfo_Op_) GetOutput() int32 {
	if m != nil && m.Output != nil {
		return *m.Output
	}
	return 0
}

func (m *Resp_MineInfo_Op_) GetHeroes() []*BattleHeroInfo_ {
	if m != nil {
		return m.Heroes
	}
	return nil
}

func (m *Resp_MineInfo_Op_) GetRole() *Role_Info_ {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *Resp_MineInfo_Op_) GetDefUserID() int64 {
	if m != nil && m.DefUserID != nil {
		return *m.DefUserID
	}
	return 0
}

// 推送矿山情报
type Resp_Mine_Info_ struct {
	UserID           *int64   `protobuf:"varint,1,opt,name=UserID" json:"UserID,omitempty"`
	InfoType         []int32  `protobuf:"varint,2,rep,name=InfoType" json:"InfoType,omitempty"`
	AtkName          []string `protobuf:"bytes,3,rep,name=AtkName" json:"AtkName,omitempty"`
	AtkUserID        []int64  `protobuf:"varint,4,rep,name=AtkUserID" json:"AtkUserID,omitempty"`
	Gold             []int32  `protobuf:"varint,5,rep,name=Gold" json:"Gold,omitempty"`
	Badge            []int32  `protobuf:"varint,6,rep,name=Badge" json:"Badge,omitempty"`
	Status           []int32  `protobuf:"varint,7,rep,name=Status" json:"Status,omitempty"`
	InfoID           []int32  `protobuf:"varint,8,rep,name=InfoID" json:"InfoID,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Resp_Mine_Info_) Reset()                    { *m = Resp_Mine_Info_{} }
func (m *Resp_Mine_Info_) String() string            { return proto.CompactTextString(m) }
func (*Resp_Mine_Info_) ProtoMessage()               {}
func (*Resp_Mine_Info_) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{16} }

func (m *Resp_Mine_Info_) GetUserID() int64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

func (m *Resp_Mine_Info_) GetInfoType() []int32 {
	if m != nil {
		return m.InfoType
	}
	return nil
}

func (m *Resp_Mine_Info_) GetAtkName() []string {
	if m != nil {
		return m.AtkName
	}
	return nil
}

func (m *Resp_Mine_Info_) GetAtkUserID() []int64 {
	if m != nil {
		return m.AtkUserID
	}
	return nil
}

func (m *Resp_Mine_Info_) GetGold() []int32 {
	if m != nil {
		return m.Gold
	}
	return nil
}

func (m *Resp_Mine_Info_) GetBadge() []int32 {
	if m != nil {
		return m.Badge
	}
	return nil
}

func (m *Resp_Mine_Info_) GetStatus() []int32 {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Resp_Mine_Info_) GetInfoID() []int32 {
	if m != nil {
		return m.InfoID
	}
	return nil
}

func init() {
	proto.RegisterType((*Mine_Player_Info_)(nil), "protocol.Mine_Player_Info_")
	proto.RegisterType((*Req_Get_MineInfo_)(nil), "protocol.Req_Get_MineInfo_")
	proto.RegisterType((*Resp_Get_MineInfo_)(nil), "protocol.Resp_Get_MineInfo_")
	proto.RegisterType((*Req_Mine_Upgrade_)(nil), "protocol.Req_Mine_Upgrade_")
	proto.RegisterType((*Resp_Mine_Upgrade_)(nil), "protocol.Resp_Mine_Upgrade_")
	proto.RegisterType((*Req_Save_MineHeroes_)(nil), "protocol.Req_Save_MineHeroes_")
	proto.RegisterType((*Resp_Save_MineHeroes_)(nil), "protocol.Resp_Save_MineHeroes_")
	proto.RegisterType((*Req_Refresh_MineEnemy_)(nil), "protocol.Req_Refresh_MineEnemy_")
	proto.RegisterType((*Resp_Refresh_MineEnemy_)(nil), "protocol.Resp_Refresh_MineEnemy_")
	proto.RegisterType((*Req_Mine_Collect_)(nil), "protocol.Req_Mine_Collect_")
	proto.RegisterType((*Resp_Mine_Collect_)(nil), "protocol.Resp_Mine_Collect_")
	proto.RegisterType((*Req_Battle_Result_)(nil), "protocol.Req_Battle_Result_")
	proto.RegisterType((*Resp_Battle_Result_)(nil), "protocol.Resp_Battle_Result_")
	proto.RegisterType((*Resp_Plunder_Mine_)(nil), "protocol.Resp_Plunder_Mine_")
	proto.RegisterType((*Req_MineInfo_Op_)(nil), "protocol.Req_MineInfo_Op_")
	proto.RegisterType((*Resp_MineInfo_Op_)(nil), "protocol.Resp_MineInfo_Op_")
	proto.RegisterType((*Resp_Mine_Info_)(nil), "protocol.Resp_Mine_Info_")
}

func init() { proto.RegisterFile("mine.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 824 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0x95, 0xe3, 0xd8, 0x49, 0x26, 0x40, 0x88, 0xf9, 0xf3, 0xf7, 0xf5, 0x06, 0x19, 0x55, 0xb4,
	0xaa, 0xc4, 0x05, 0x6f, 0x40, 0x08, 0xa2, 0x48, 0x10, 0x50, 0x08, 0x52, 0x7b, 0x65, 0x39, 0xc9,
	0x90, 0x46, 0x38, 0xb6, 0x71, 0x6c, 0x0a, 0xef, 0x50, 0x55, 0xaa, 0x54, 0xf5, 0xa1, 0xfa, 0x08,
	0x7d, 0x92, 0x5e, 0x76, 0x77, 0x76, 0x8d, 0x1d, 0xc7, 0x21, 0x5c, 0x59, 0x3b, 0xbb, 0x3b, 0x3b,
	0xe7, 0xcc, 0x39, 0x63, 0x80, 0xc9, 0xd8, 0xc3, 0x83, 0x20, 0xf4, 0x23, 0xdf, 0xa8, 0xd2, 0x67,
	0xe0, 0xbb, 0xff, 0x43, 0xdf, 0x99, 0xca, 0xa8, 0xf5, 0x19, 0x9a, 0x17, 0xec, 0x8c, 0x7d, 0xe5,
	0x3a, 0x4f, 0x18, 0xda, 0x67, 0xde, 0xad, 0x6f, 0x1b, 0x6b, 0xa0, 0xdf, 0x4c, 0x31, 0x3c, 0x6b,
	0x9b, 0xca, 0xae, 0xf2, 0x4e, 0x33, 0xd6, 0xa1, 0xda, 0x19, 0x0f, 0xee, 0x3c, 0x67, 0x82, 0x66,
	0x89, 0x45, 0x6a, 0x86, 0x05, 0xe5, 0xae, 0xef, 0xa2, 0xa9, 0xb2, 0x55, 0xfd, 0x70, 0xf3, 0x20,
	0xc9, 0x7d, 0xc0, 0xa3, 0x22, 0x8b, 0xb5, 0x07, 0xcd, 0x2e, 0xde, 0xdb, 0xa7, 0x18, 0xd9, 0xfc,
	0x89, 0xa2, 0xd4, 0xaa, 0xf5, 0x4d, 0x05, 0xa3, 0x8b, 0xd3, 0xe0, 0xe5, 0x63, 0xc6, 0x2a, 0x68,
	0xc7, 0x71, 0x78, 0xfe, 0x40, 0xcf, 0x6b, 0x46, 0x13, 0x6a, 0x6c, 0x79, 0x19, 0x47, 0x41, 0x1c,
	0x51, 0x0d, 0x14, 0x6a, 0xe1, 0x68, 0xec, 0xf5, 0xc6, 0xac, 0xc8, 0x32, 0x5d, 0x6a, 0x40, 0xe5,
	0xc4, 0x1b, 0x52, 0x40, 0xa3, 0x80, 0x01, 0x70, 0xe5, 0xc6, 0xde, 0x10, 0xc3, 0x4e, 0x3c, 0x31,
	0x75, 0xba, 0xb7, 0x03, 0x8d, 0x0e, 0x3e, 0x46, 0x32, 0x4e, 0x87, 0x2b, 0x74, 0x78, 0x13, 0x56,
	0xf8, 0xc6, 0xf9, 0x83, 0x7c, 0xa6, 0x4a, 0xc7, 0x59, 0xf4, 0x26, 0x18, 0x85, 0xce, 0x10, 0x5b,
	0xce, 0x70, 0x84, 0x66, 0x8d, 0xa2, 0xef, 0x41, 0xff, 0x88, 0xa1, 0x8f, 0x53, 0x13, 0x76, 0x55,
	0x46, 0xc8, 0x7f, 0x29, 0x21, 0x2d, 0x27, 0x8a, 0x5c, 0xe4, 0xbb, 0x02, 0xd9, 0x3e, 0xd4, 0x7a,
	0x4e, 0xdf, 0x25, 0x9c, 0x66, 0x3d, 0x4f, 0xdf, 0xf3, 0x96, 0x6d, 0xac, 0x40, 0xf9, 0x3a, 0xc2,
	0xc0, 0x5c, 0xa1, 0x17, 0xb6, 0x60, 0xf5, 0xc2, 0x79, 0xcc, 0x54, 0xbf, 0x4a, 0xe1, 0x0d, 0xa8,
	0x1f, 0xfb, 0xae, 0x8b, 0x83, 0x88, 0x2a, 0x5f, 0x4b, 0x2a, 0x67, 0x4c, 0x8e, 0xbd, 0xd1, 0x71,
	0x9b, 0xa2, 0x8d, 0x24, 0xda, 0xc5, 0xdb, 0x10, 0xa7, 0x5f, 0xae, 0x9c, 0xf1, 0x00, 0xcd, 0x75,
	0x4a, 0xc0, 0x88, 0x6d, 0x61, 0xd7, 0xef, 0x9b, 0x4d, 0xbe, 0xb4, 0xde, 0x8a, 0x9e, 0x91, 0x24,
	0x24, 0x4e, 0x9b, 0xb7, 0xbf, 0xe7, 0x84, 0x23, 0x64, 0x5c, 0x08, 0x41, 0x58, 0x7d, 0xd9, 0xb4,
	0xd9, 0x73, 0xcf, 0x4d, 0x52, 0xe6, 0x9b, 0x54, 0x4a, 0xd8, 0x9b, 0xe1, 0x54, 0x2d, 0xe4, 0xb4,
	0x4c, 0x6f, 0x7c, 0x57, 0x60, 0x93, 0xd7, 0x72, 0xed, 0x3c, 0x20, 0x3d, 0x24, 0x18, 0x9e, 0xd7,
	0x46, 0x4a, 0x7e, 0x69, 0x19, 0xf9, 0x1f, 0x40, 0x17, 0x42, 0x97, 0xc2, 0x7d, 0x93, 0x1e, 0x2d,
	0x74, 0x01, 0x0f, 0x32, 0x3c, 0xa2, 0xa0, 0xdf, 0x0a, 0x6c, 0x11, 0xea, 0xa5, 0x15, 0xcd, 0x68,
	0xb1, 0x94, 0xd7, 0xa2, 0x4a, 0x81, 0x5c, 0xe7, 0xca, 0x39, 0x28, 0xda, 0x32, 0x28, 0xf9, 0x26,
	0xeb, 0x94, 0x60, 0x1f, 0x2a, 0x5d, 0xfc, 0xea, 0x84, 0xc3, 0x29, 0x53, 0x31, 0xcf, 0xb0, 0x9d,
	0xb1, 0x26, 0x6d, 0x48, 0x73, 0x7e, 0x82, 0x6d, 0x4e, 0xae, 0x54, 0x04, 0xa1, 0x39, 0xf1, 0x70,
	0xf2, 0x34, 0x0f, 0x86, 0xe9, 0xb0, 0xf7, 0x14, 0xa0, 0xec, 0x60, 0x4a, 0x8a, 0x9a, 0x8c, 0x06,
	0xba, 0x77, 0x33, 0x1e, 0x32, 0x0c, 0x2a, 0x73, 0xf4, 0x1f, 0x05, 0x76, 0x88, 0xa6, 0xd7, 0xe5,
	0xee, 0xa4, 0x43, 0x25, 0x9f, 0x9b, 0xad, 0xa5, 0x4e, 0xa8, 0x01, 0x9c, 0x43, 0x36, 0x25, 0x4e,
	0x7d, 0x77, 0x48, 0x7e, 0xa6, 0x03, 0x2c, 0x70, 0xf2, 0x18, 0x48, 0x2f, 0xa7, 0xf4, 0x55, 0x96,
	0xd1, 0x97, 0x0c, 0xb0, 0xea, 0xe2, 0x01, 0xc6, 0xdb, 0xd8, 0xc6, 0x5b, 0x59, 0x70, 0x8d, 0xc6,
	0xd5, 0x5e, 0xc6, 0x1f, 0xb2, 0x7d, 0xf3, 0x33, 0xed, 0x30, 0xeb, 0x8e, 0x45, 0xa7, 0x38, 0x76,
	0x82, 0x42, 0xbc, 0x5a, 0xbf, 0x14, 0x7e, 0xe9, 0xde, 0x16, 0x75, 0x32, 0xee, 0xa6, 0xb1, 0x5b,
	0x70, 0x89, 0xad, 0xc5, 0xd6, 0x82, 0x76, 0x64, 0x28, 0x2a, 0xe7, 0x28, 0xd2, 0x12, 0x53, 0xa6,
	0x98, 0xf4, 0x99, 0x26, 0x54, 0x92, 0x26, 0x70, 0xf4, 0x6c, 0x97, 0x06, 0x9e, 0xf5, 0x43, 0x81,
	0x0d, 0x42, 0xb3, 0xa4, 0xb2, 0xcc, 0xcb, 0xa5, 0x44, 0x19, 0x2c, 0x20, 0x1c, 0xad, 0xe6, 0x6a,
	0x29, 0xe7, 0x8a, 0x17, 0xb5, 0x65, 0xc4, 0xab, 0xbf, 0x28, 0xde, 0x23, 0x49, 0xb0, 0x1c, 0x87,
	0x82, 0xe8, 0xdc, 0x74, 0x57, 0x16, 0x4d, 0x77, 0xf2, 0xa3, 0xd5, 0x86, 0xf5, 0xa4, 0x91, 0x94,
	0xf3, 0x32, 0x28, 0x24, 0x5b, 0x52, 0x51, 0x9a, 0xd3, 0x3a, 0x99, 0xd8, 0xfa, 0xab, 0x70, 0x3d,
	0xc8, 0x56, 0x2f, 0xcc, 0xc3, 0xee, 0xf1, 0xbd, 0x8c, 0x8b, 0x32, 0x64, 0x15, 0x53, 0x93, 0x25,
	0x4f, 0x90, 0x93, 0x74, 0x49, 0xcf, 0x59, 0xa5, 0x92, 0xb3, 0x4a, 0x35, 0xe7, 0x84, 0xda, 0x6b,
	0x9d, 0x00, 0xaf, 0x75, 0x42, 0x9d, 0xa0, 0xff, 0x54, 0xa0, 0x91, 0xaa, 0xbc, 0xf8, 0xaf, 0x3d,
	0x0b, 0x5c, 0x15, 0xc0, 0x8f, 0xa2, 0x3b, 0x02, 0xa2, 0xb2, 0x40, 0x8d, 0x67, 0x66, 0x01, 0x79,
	0x8b, 0x06, 0xc8, 0xb3, 0x31, 0x34, 0xba, 0xc1, 0x7f, 0x50, 0x44, 0x83, 0x4e, 0x4b, 0xf6, 0xc4,
	0x75, 0xe4, 0x44, 0xb1, 0xb0, 0xb8, 0x36, 0x23, 0x57, 0xb6, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xcc, 0x7f, 0x13, 0x0c, 0xf2, 0x08, 0x00, 0x00,
}