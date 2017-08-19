package common

import (
	"antnet"
	//"proto/pb"
)

var OnlineCount int32 = 0

func RegisterRpc(cmd, act uint8, c2s, s2c interface{}, fun antnet.HandlerFunc) {
	PbParser.Register(cmd, act, c2s, s2c)
	RpcHandler.Register(cmd, act, fun)
}
func init() {

	//读取配置相关
	AddConfig()
	SetConfigData(3, 4, "conf/doc/")
	SetFunc()
}
