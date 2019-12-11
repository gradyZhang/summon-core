package sol
import (
	"sync"
	"github.com/gradyZhang/summon-core/proto_msg"
	"github.com/gogo/protobuf/proto"
	"time"
	"github.com/gradyZhang/summon-core/def"
)
const(
	Mar_Type_None int32 = iota
	Mar_Type_Win5 					// 1 5连胜
	Mar_Type_Win10					// 2 10连胜
	Mar_Type_WinN					// 3 N连胜
	Mar_Type_Shutdown				// 4 终结连胜（5，>= 10）
	Mar_Type_Chest3            		// 5 开到品质3宝箱
	Mar_Type_Chest4					// 6 开到品质4宝箱

)
const Msg_Max = 1000	// 消息长度
// 游戏内消息推送
type GamePush struct {
	ID int32 //唯一ID
	Content string //内容
	Times int32 //循环次数
	Gap int32 //时间间隔
	Trigger int64 //触发时间
	Type int32 // 消息类型
	Args []string	// 客户端消息拼接参数
}

type MarqueeManager struct{
	sync.Mutex
	AutoId int32 //一个自增ID
	GamePush []*GamePush	// 游戏内消息推送（游戏内触发）
	updateTime int64
}

var marqueeMgr *MarqueeManager
var onceMarquee sync.Once

func GetInsMarqueeMgr() *MarqueeManager {
	onceMarquee.Do(func() {
		marqueeMgr = &MarqueeManager{}
		marqueeMgr.init()
	})
	return marqueeMgr
}

func (mgr *MarqueeManager) init() {
	mgr.GamePush = []*GamePush{}
	mgr.AutoId = 0
	mgr.updateTime = time.Now().Unix()
}

//心跳
func (mgr *MarqueeManager)Update() *protocol.GS2C {
	mgr.Lock()
	defer mgr.Unlock()
	now := time.Now().Unix()
	if now - mgr.updateTime < 5 {	// 5秒更新一次
		return nil
	}
	mgr.updateTime = now
	if len( mgr.GamePush ) <= 0{
		return nil
	}
	resp := &protocol.GS2C{
		Session:proto.Int32(0),
		Sequence:[]int32{def.RespMarquee},
		Resp_Marquee:&protocol.Resp_Marquee_{
		},
	}

	var gpDel []int32 //准备删除的id 列表
	curTime := time.Now().Unix()
	
	// 游戏内消息推送
	for _,gPush := range mgr.GamePush {
		if len(gpDel) >= 5 {	// 最多给5条
			break
		}
		if gPush.Times <= 0{//不能在播放了, 已经没有次数了
			gpDel = append(gpDel,gPush.ID)
			continue
		}
		if gPush.Trigger > curTime{//还没有到时间触发
			continue
		}
		infos := &protocol.GamePush_Info_{
			Type:proto.Int32(gPush.Type),
			Content:proto.String( gPush.Content ),
			Args:gPush.Args,
		}
		resp.Resp_Marquee.PushInfos = append( resp.Resp_Marquee.PushInfos, infos)

		gPush.Times -= 1
		gPush.Trigger = curTime + int64( gPush.Gap ) //重新设置下触发时间a
		if gPush.Times <= 0{//往预删除列表里面插东西
			gpDel = append(gpDel,gPush.ID)
		}
	}
	if len( gpDel ) > 0{
		for _, id := range gpDel{
			for i, marquee := range mgr.GamePush{
				if id != marquee.ID{//移除
					continue
				}
				mgr.GamePush = append( mgr.GamePush[:i], mgr.GamePush[i+1:]...)
			}
		}
	}
	if len( resp.Resp_Marquee.PushInfos ) > 0{
		return resp //需要广播这条消息包
	}
	return nil
}

// 添加游戏内消息推送
func (mgr *MarqueeManager) InsertGamePush(m *GamePush) {
	mgr.Lock()
	defer mgr.Unlock()
	
	if len(mgr.GamePush) >= Msg_Max {	// 超上限了，就不发了
		return
	}
	if m.Type == Mar_Type_None {
		return
	}
	mgr.AutoId += 1
	m.ID = mgr.AutoId
	if m.Times <= 0 {
		m.Times = 1
	}
	mgr.GamePush = append(mgr.GamePush,m)
}

// 移除所有跑马灯
func (mgr *MarqueeManager) RemoveAll() {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.GamePush = []*GamePush{}
}
