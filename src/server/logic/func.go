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
}
