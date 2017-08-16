package control

import (
	"antnet"
	. "common"
	"container/list"
	"proto/pb"
	"server/match/model"
	"sync"
	"time"
)

type t_match_chan chan *pb.GamerStartPVPMatchL2M

var matchChanMap = map[pb.PVPType]t_match_chan{
	pb.PVPType_M2V2: make(t_match_chan, 64),
	pb.PVPType_M1V1: make(t_match_chan, 64),
	pb.PVPType_M1:   make(t_match_chan, 64),
	pb.PVPType_M3V3: make(t_match_chan, 64),
	pb.PVPType_M5V5: make(t_match_chan, 64),
}

type roomMap struct {
	sync.RWMutex
	M map[string]t_match_chan
}

var RoomMap = &roomMap{M: make(map[string]t_match_chan)}

func HandleStartPVPMatchL2M(msgque antnet.IMsgQue, msg *antnet.Message) bool {

	return true
}

func HandleStartPVPMatchStop(msgque antnet.IMsgQue, msg *antnet.Message) bool {

	return true
}

func PVPMatchingM1(cstop chan struct{}) {

}

func PVPMatchReady(cstop chan struct{}, c t_match_chan, servers map[int32]int, p *pb.NotifyGamerPVPMatch, matchGamer map[int32]*pb.GamerStartPVPMatchL2M) {

}

func PVPMatchingOk(matchGamer map[int32]*pb.GamerStartPVPMatchL2M) {

}

func PVPMatching(cstop chan struct{}, typ pb.PVPType) {

}
