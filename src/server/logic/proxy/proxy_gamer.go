package proxy

import (
	"antnet"
	"common"
	"model"
	"proto/pb"
	"sync"
	"time"
)

type IActor interface {
	GetMutex() *sync.Mutex
}

type IControlGamer interface {
	GetId() int32
	SetProxy(proxy *GamerProxy)
	GetProxy() *GamerProxy

	ProcessClientMsg(msgque antnet.IMsgQue, msg *antnet.Message) bool
	ProcessInnerMsg(msgque antnet.IMsgQue, msg *antnet.Message) interface{}
}

//玩家消息
type GamerMsgType int

const (
	GamerMsgTypeClient = iota
	GamerMsgTypeInner
)

type GamerMsg struct {
	Type   GamerMsgType
	MsgQue antnet.IMsgQue
	Msg    *antnet.Message
	cout   chan interface{}
}

type extModel struct {
	sync.RWMutex
	data []byte
}

type GamerProxy struct {
	Id int32

	Online    bool //是否在线
	MatchType pb.PVPType
	MsgQue    antnet.IMsgQue //玩家的消息队列

	control IControlGamer

	mutex       sync.Mutex
	modelMain   extModel
	modelPlayer extModel
}

func GenGamerProxy(control IControlGamer) *GamerProxy {
	proxy := &GamerProxy{
		Id:      control.GetId(),
		control: control,
	}

	control.SetProxy(proxy)
	return proxy
}

func (r *GamerProxy) GetMutex() *sync.Mutex {
	return &r.mutex
}

func (r *GamerProxy) SendMsg(msg *GamerMsg) (re bool) {
	r.handler(msg)
	return true
}

func (r *GamerProxy) SendClientMsg(msgque antnet.IMsgQue, msg *antnet.Message) {
	r.SendMsg(&GamerMsg{
		MsgQue: msgque,
		Msg:    msg,
	})
}

func (r *GamerProxy) SendInnerMsg(msgque antnet.IMsgQue, msg *antnet.Message) {
	r.SendMsg(&GamerMsg{
		MsgQue: msgque,
		Msg:    msg,
		Type:   GamerMsgTypeInner,
	})
}

func (r *GamerProxy) SendInnerMsgWait(act uint8, actor IActor, msg *antnet.Message) (interface{}, error) {
	cout := make(chan interface{}, 1)
	m := &GamerMsg{cout: cout, Msg: msg, Type: GamerMsgTypeInner}

	var data interface{} = nil
	err := common.ErrSendInnerMsg
	if !r.SendMsg(m) {
		close(cout)
		return nil, err
	}

	actor.GetMutex().Unlock()
	select {
	case data = <-cout:
	case <-time.After(time.Millisecond * time.Duration(10)):
		return nil, common.ErrInnerMsgTimeout
	}

	if data == nil {
		err = nil
	}
	close(cout)
	actor.GetMutex().Lock()
	return data, err
}

func (r *GamerProxy) Logout() {
	r.SendMsg(&GamerMsg{
		Msg:  antnet.NewTagMsg(common.GAME_CMD_INNER, common.GAME_CMD_INNER_LOGOUT, 0),
		Type: GamerMsgTypeInner,
	})
}

func (r *GamerProxy) SetMain(data *pb.Gamer) {
	r.modelMain.Lock()
	r.modelMain.data = antnet.PbData(data)
	r.modelMain.Unlock()
}

func (r *GamerProxy) SetPlayer(data []*pb.Player) {
	r.modelPlayer.Lock()
	r.modelPlayer.data = antnet.PbData(&pb.GamerProxy{Players: data})
	r.modelPlayer.Unlock()
}

func (r *GamerProxy) Player() []*pb.Player {
	playerSlice := []*pb.Player{}
	if r.modelPlayer.data == nil {
		m := model.ModelGamerPlayer(r.Id)
		if m != nil {
			for _, v := range m.All() {
				playerSlice = append(playerSlice, v.Player)
			}
			r.modelPlayer.Lock()
			r.modelPlayer.data = antnet.PbData(&pb.GamerProxy{Players: playerSlice})
			r.modelPlayer.Unlock()
		}
	}

	if r.modelPlayer.data == nil {
		return nil
	}

	if len(playerSlice) == 0 {
		p := &pb.GamerProxy{}
		r.modelPlayer.Lock()
		antnet.ParsePbData(r.modelPlayer.data, p)
		r.modelPlayer.Unlock()
		playerSlice = p.Players
	}

	return playerSlice
}

func (r *GamerProxy) Main() *pb.Gamer {
	if r.modelMain.data == nil {
		m := model.ModelGamerMain(r.Id)
		if m != nil {
			r.modelMain.Lock()
			r.modelMain.data = m.PbData(m.Gamer)
			r.modelMain.Unlock()
		}
	}

	if r.modelMain.data == nil {
		return nil
	}

	g := &pb.Gamer{}
	r.modelMain.Lock()
	antnet.ParsePbData(r.modelMain.data, g)
	r.modelMain.Unlock()
	return g
}

func (r *GamerProxy) handler(gmsg *GamerMsg) {
	antnet.Go(func() {
		r.mutex.Lock()
		if gmsg.Type == GamerMsgTypeClient {
			if !r.control.ProcessClientMsg(gmsg.MsgQue, gmsg.Msg) {
				gmsg.MsgQue.Stop()
			}
		} else {
			re := r.control.ProcessInnerMsg(gmsg.MsgQue, gmsg.Msg)
			if gmsg.cout != nil {
				gmsg.cout <- re
			}
		}
		r.mutex.Unlock()
	})
}
