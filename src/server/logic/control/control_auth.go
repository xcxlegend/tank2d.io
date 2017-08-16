package control

import (
	"antnet"
	. "common"
	"model"
	"proto/pb"
	"time"
)

func HandleLogin(gamer *ControlGamer, msgque antnet.IMsgQue, msg *antnet.Message) bool {
	gamer.Proxy.Online = true
	gamer.Model = model.GetModelGamer(gamer.Id)

	return true
}

func HandleFastReConn(gamer *ControlGamer, msgque antnet.IMsgQue, msg *antnet.Message) bool {
	return true
}

func HandleLogout(gamer *ControlGamer, msgque antnet.IMsgQue, msg *antnet.Message) interface{} {
	gamer.Proxy.Online = false
	if gamer.Proxy.MsgQue != nil {
		gamer.Proxy.MsgQue.SetUser(nil)
		gamer.Proxy.MsgQue.Stop()
		gamer.Proxy.MsgQue = nil
	}
	if gamer.Proxy.MatchType != pb.PVPType_None {
		l2m := &pb.GamerStartPVPMatchL2M{
			Id:      gamer.Model.Main().Id,
			PvpType: gamer.Proxy.MatchType.Enum(),
		}
		gamer.Proxy.MatchType = pb.PVPType_None
		SendToMatchServer(antnet.NewMsg(GAME_CMD_INNER, GAME_CMD_INNER_PVP_MATCH_STOP, 0, 0, antnet.PbData(l2m)))
	}
	gamer.Model.Main().Save(model.RedisNotifyGamerOffline)
	antnet.LogInfo("gamer logout %v", gamer.Id)
	SubOnlineCount()
	return nil
}
