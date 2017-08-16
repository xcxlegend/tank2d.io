package pvp

import (
	"antnet"
	. "common"
	"model"
	"proto/pb"
	"server/pvp/control"
)

func OnNewMsgQue(msgque antnet.IMsgQue) bool {
	msgque.SetTimeout(Config.Server.Net.Timeout)
	return true
}

func Start() {
	ExtHandler.ProcessMsg = control.OnProcessPVPMsg
	ExtHandler.NewMsgQue = OnNewMsgQue
	StartListen()
	//model.UploadServerInfo()
	//m, err := model.UpdateServerInfo()
	//if err == nil {
	//	StartConnect(m, func(info *pb.ServerInfo) bool {
	//		return info.GetServerType() == "match"
	//	})
	//}
	//antnet.SetTimeout(3000, func(args ...interface{}) int {
	//	model.UploadServerStatus()
	//	return 3000
	//})
}
