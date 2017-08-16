package logic

import (
	"antnet"
	. "common"
	"model"
	"proto/pb"
	"server/logic/control"
)

func Start() {
	ExtHandler.DelMsgQue = control.OnGamerDisConnect
	ExtHandler.ProcessMsg = control.OnProcessGamerMsg
	ExtHandler.NewMsgQue = control.OnNewMsgQue
	StartListen()
	model.UploadServerInfo()
	m, err := model.UpdateServerInfo()
	if err == nil {
		StartConnect(m, func(info *pb.ServerInfo) bool {
			return info.GetServerType() == "match"
		})
	}
	antnet.SetTimeout(3000, func(args ...interface{}) int {
		model.UploadServerStatus()
		return 3000
	})
}
