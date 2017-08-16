package control

import (
	. "common"
	"proto/pb"
)

func init() {
	RegisterRpc(GAME_CMD_INNER, GAME_CMD_INNER_PVP_MATCH_STOP, &pb.GamerStartPVPMatchL2M{}, nil, HandleStartPVPMatchStop)
	RegisterRpc(GAME_CMD_INNER, GAME_CMD_INNER_PVP_MATCH_L2M, &pb.GamerStartPVPMatchL2M{}, nil, HandleStartPVPMatchL2M)
}
