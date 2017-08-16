package common

import (
	"antnet"
	"proto/pb"
)

var ServerInfo = &pb.ServerInfo{}
var OnlineCount int32 = 0

func RegisterRpc(cmd, act uint8, c2s, s2c interface{}, fun antnet.HandlerFunc) {
	PbParser.Register(cmd, act, c2s, s2c)
	RpcHandler.Register(cmd, act, fun)
}
func init() {
	RpcHandler.NewMsgQue = onNewRpcMsgQue
	RpcHandler.DelMsgQue = onDelRpcMsgQue
	RpcHandler.ConnectComplete = onRpcConnectComplete

	RegisterRpc(GAME_CMD_INNER, GAME_CMD_INNER_PVP_SERVER_HELLO, &pb.ServerInfo{}, nil, HandleServerHello)

	//读取配置相关
	AddConfig()
	SetConfigData(3, 4, "conf/doc/")
	SetFunc()
}
