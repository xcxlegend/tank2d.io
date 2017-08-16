package control

import (
	"antnet"
	. "common"
	"model"
)

type CmdNewGamerC2S struct {
	New   string `match:"k"`
	Gamer string `match:"k"`
	Name  string
}

type CmdNewGamerS2C struct {
	Name  string
	Id    int32
	Error error
}

type CmdGetStatusC2S struct {
	Status string `match:"k"`
}

type CmdGetStatusS2C struct {
	AntNet *antnet.Statis
}

type CmdSetPvpMinFrameTimeC2S struct {
	Set          string `match:"k"`
	PVP          string `match:"k"`
	MinFrameTime int
}

func HandleNewGamer(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	c2s := msg.C2S().(*CmdNewGamerC2S)
	s2c := msg.S2C().(*CmdNewGamerS2C)
	m, err := model.NewModelGamer(c2s.Name, 1)
	if m == nil {
		s2c.Error = err
	} else {
		s2c.Name = m.Main().GetName()
		s2c.Id = m.GetId()
	}
	msgque.SendStringLn(msg.S2CString())
	return true
}

func HandleGetStatus(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	s2c := msg.S2C().(*CmdGetStatusS2C)
	s2c.AntNet = antnet.GetStatis()
	msgque.SendStringLn(msg.S2CString())
	return true
}

func HandleSetPvpMinFrameTime(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	c2s := msg.C2S().(*CmdSetPvpMinFrameTimeC2S)

	Config.Global.PvpMinFrameTime = int32(c2s.MinFrameTime)
	antnet.LogInfo("set PvpMinFrameTime to %v", Config.Global.PvpMinFrameTime)
	msgque.SendStringLn(antnet.Sprintf("set PvpMinFrameTime to %v", Config.Global.PvpMinFrameTime))
	return true
}
