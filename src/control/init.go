package control

import (
	. "common"
)

func init() {
	CmdParser.RegisterMsg(&CmdNewGamerC2S{}, &CmdNewGamerS2C{})
	CmdHandler.RegisterMsg(&CmdNewGamerC2S{}, HandleNewGamer)

	CmdParser.RegisterMsg(&CmdGetStatusC2S{}, &CmdGetStatusS2C{})
	CmdHandler.RegisterMsg(&CmdGetStatusC2S{}, HandleGetStatus)

	CmdParser.RegisterMsg(&CmdSetPvpMinFrameTimeC2S{}, nil)
	CmdHandler.RegisterMsg(&CmdSetPvpMinFrameTimeC2S{}, HandleSetPvpMinFrameTime)
}
