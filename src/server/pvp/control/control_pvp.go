package control

import (
	"antnet"
	. "common"
	"proto/pb"
	"server/pvp/model"
	"sync"
	"unsafe"
)

type PVPMsg struct {
	MsgQue antnet.IMsgQue
	Msg    *antnet.Message
	C2S    *pb.GamerPVPSyncC2S
}

type GamerInputChangePlayer struct {
	Bef uint8
	Aft uint8
}

type GamerInput struct {
	Flg int8
	Yaw int8
	Btn uint8
	Pos int8
}

type PVPGamer struct {
	Input    *pb.GamerInput
	Msg      *PVPMsg
	Match    *pb.PVPMatchInfo
	Index    uint16
	ReSend   []int
	LastRecv int64
}

type PVPFrame struct {
	Data []byte
	Len  int
}

func PVPing(cstop chan struct{}, c chan *PVPMsg, pvp *pb.NotifyGamerPVPMatch) {

}

var sessionMap = &struct {
	sync.RWMutex
	M map[string]chan *PVPMsg
}{M: make(map[string]chan *PVPMsg)}

func OnProcessPVPMsg(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	var c chan *PVPMsg
	c2s := msg.C2S().(*pb.GamerPVPSyncC2S)
	if c2s == nil {
		antnet.LogInfo("recv c2s is null")
		return true
	}
	if msgque.GetUser() == c2s.GetSession() {
		return true
	}
	sessionMap.Lock()
	c = sessionMap.M[c2s.GetSession()]
	if c == nil {
		pvp, err := model.GetPVPInfo(c2s.GetSession())
		if err != nil {
			msgque.SetUser(c2s.GetSession)
		} else {
			c = make(chan *PVPMsg, 64)
			sessionMap.M[c2s.GetSession()] = c
			antnet.Go2(func(cstop chan struct{}) {
				PVPing(cstop, c, pvp)
			})
		}
	}
	sessionMap.Unlock()
	if c != nil && len(c) < cap(c) {
		c <- &PVPMsg{msgque, msg, c2s}
	}
	return true
}
