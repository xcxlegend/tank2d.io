package control

import (
	. "common"
	"proto/pb"
)

func init() {
	//
	PbParser.Register(GAME_CMD_PVP, GAME_CMD_PVP_SYNC_C2S, &pb.GamerPVPSyncC2S{}, nil)
	//PbParser.ErrType = antnet.ParseErrTypeAlways
}
