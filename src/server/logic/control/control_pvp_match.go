package control

import (
	"antnet"
	. "common"
	"proto/pb"
)

var matchMap = map[int32]*pb.GamerStartPVPMatchC2S{}

func HandleStartPVPMatchC2S(gamer *ControlGamer, msgque antnet.IMsgQue, msg *antnet.Message) bool {
	c2s := msg.C2S().(*pb.GamerStartPVPMatchC2S)
	antnet.LogInfo("gamer start pvp match gid:%v type:%v session:%v", gamer.Id, c2s.GetPvpType(), c2s.GetSession())
	if c2s.Session == nil {
		c2s.Players = nil
		for k, v := range gamer.Model.Player().All() {
			c2s.Players = append(c2s.Players, &pb.PVPPlayer{Id: pb.Int32(k), Level: v.Level})

		}
	}
	l2m := &pb.GamerStartPVPMatchL2M{
		Id:       c2s.Id,
		PvpType:  c2s.PvpType,
		MapId:    c2s.MapId,
		Cancel:   c2s.Cancel,
		Players:  c2s.Players,
		Session:  c2s.Session,
		ServerId: pb.Int32(Config.Server.Id),
		Gamer: &pb.GamerMini{
			Id:    gamer.Model.Main().Id,
			Name:  gamer.Model.Main().Name,
			Level: gamer.Model.Main().Level,
		},
	}

	if SendToMatchServer(antnet.NewMsg(GAME_CMD_INNER, GAME_CMD_INNER_PVP_MATCH_L2M, 0, 0, antnet.PbData(l2m))) {
		msgque.Send(antnet.NewErrMsg(ErrOk).CopyTag(msg))
		if c2s.Cancel == nil {
			gamer.Proxy.MatchType = c2s.GetPvpType()
		} else {
			gamer.Proxy.MatchType = pb.PVPType_None
		}
	} else {
		msgque.Send(antnet.NewErrMsg(ErrServerNotFound).CopyTag(msg))
	}

	return true
}

func HandleStartPVPMatchM2L(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	c2s := msg.C2S().(*pb.NotifyGamerPVPMatch)
	for _, v := range c2s.Match {
		proxy := GetGamerProxy(v.GetGamer().GetId())
		if c2s.GetState() == pb.PVPMatchState_Done {
			proxy.MatchType = pb.PVPType_None
		}
		antnet.LogInfo("pvp match ok gid:%d msgque:%t", proxy.Id, proxy.MsgQue != nil)
		if proxy.MsgQue != nil {
			proxy.MsgQue.Send(antnet.NewMsg(GAME_CMD_NOTIFY, GAME_CMD_NOTIFY_PVP_MATCH_OK, 0, 0, msg.Data))
		}
	}

	return true
}
