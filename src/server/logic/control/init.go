package control

import (
	. "common"
	"model"
	"proto/pb"
)

var MaxGamerId = 0

func init() {

	initServerFunc() //服务器相关处理函数

}

func initServerFunc() {
	//快速断线重连
	//SetControlFunc(GAME_CMD_LOGIN, GAME_CMD_LOGIN_BYRECONN, &pb.FastReconnC2S{}, &pb.GamerLoginS2C{}, HandleFastReConn)
	////发送服务器时间戳
	//SetControlFunc(GAME_CMD_MAIN, GAME_CMD_MAIN_ECHO, &pb.ServerTimeC2S{}, &pb.ServerTime{}, HandleServerTime)
	////匹配
	//SetControlFunc(GAME_CMD_PVP, GAME_CMD_PVP_START_MATCH, &pb.GamerStartPVPMatchC2S{}, nil, HandleStartPVPMatchC2S)
}
