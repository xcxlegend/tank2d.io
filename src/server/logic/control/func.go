package control

import (
	"antnet"
	. "common"
	"model"
	"proto/pb"
	. "server/logic/proxy"
	"sync"
)

var gamerProxyMap = &struct {
	sync.Mutex
	M map[int32]*GamerProxy
}{M: make(map[int32]*GamerProxy)}

//获取玩家代码
func GetGamerProxy(gid int32) *GamerProxy {
	gamerProxyMap.Lock()
	g, ok := gamerProxyMap.M[gid]
	if ok {
		gamerProxyMap.Unlock()
		return g
	}

	g = GenGamerProxy(&ControlGamer{Id: gid, Model: model.GetModelGamer(gid)})
	gamerProxyMap.M[gid] = g
	gamerProxyMap.Unlock()

	ServerInfo.Load = pb.Int32(int32(len(gamerProxyMap.M)))
	return g
}

//发送数据到在线玩家，msg可以是任意值
func SendInnerMsg2OnlineGamer(gid int32, act uint8, msg interface{}) bool {
	gamerProxyMap.Lock()
	g, ok := gamerProxyMap.M[gid]
	if ok && g.Online {
		m := antnet.NewTagMsg(GAME_CMD_INNER, act, 0)
		m.User = msg
		g.SendInnerMsg(nil, m)
	}
	gamerProxyMap.Unlock()
	return ok
}

//发送数据到在线玩家，msg可以是任意值
func SendNotifyMsg2OnlineGamer(gid int32, act uint8, data []byte) bool {
	gamerProxyMap.Lock()
	g, ok := gamerProxyMap.M[gid]
	if ok && g.Online {
		g.MsgQue.Send(antnet.NewMsg(GAME_CMD_NOTIFY, act, 0, 0, data))
	}
	gamerProxyMap.Unlock()
	return ok
}

//玩家是否在线
func IsGamerOnline(gid int32) bool {
	gamerProxyMap.Lock()
	g, ok := gamerProxyMap.M[gid]
	gamerProxyMap.Unlock()
	return ok && g.Online
}

func GetGamerId(data []byte) int32 {
	p := &pb.GamerId{}
	err := antnet.PBUnPack(data, p)
	if err != nil {
		return 0
	}
	return p.GetId()
}

//客服端连接到来
func OnNewMsgQue(msgque antnet.IMsgQue) bool {
	return true
}

//客服端连接断开
func OnGamerDisConnect(msgque antnet.IMsgQue) {
	gamer, ok := msgque.GetUser().(IControlGamer)
	if ok {
		gamer.GetProxy().Logout()
	}
}

//处理客服端消息
func OnProcessGamerMsg(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	var gid int32 = 0
	var gamer IControlGamer
	if msgque.GetUser() != nil {
		gamer = msgque.GetUser().(IControlGamer)
		gid = gamer.GetId()
	} else {
		gid = GetGamerId(msg.Data)
	}

	if gid == 0 {
		return false
	}
	if int(gid) > MaxGamerId+1000 {
		msgque.Send(antnet.NewErrMsg(ErrGamerId).CopyTag(msg))
		return false
	}

	if msg.Head.Cmd == GAME_CMD_LOGIN {
		if msg.Head.Act == GAME_CMD_LOGIN_BYSESSION {
			if gamer != nil {
				msgque.Send(antnet.NewErrMsg(ErrLoginRepeated).CopyTag(msg))
				return true
			}

			c2s := msg.C2S().(*pb.GamerLoginC2S)
			antnet.LogInfo("gamer login gid:%v channel:%v", gid, c2s.GetChannel())
			modelGamerMain, err := model.Login(gid)
			if err != nil {
				antnet.LogError("gamer login failed gid:%v err:%v", gid, err)
				msgque.Send(antnet.NewErrMsg(ErrSession).CopyTag(msg))
				return false
			}
			msg.User = modelGamerMain
		}
	}

	proxy := GetGamerProxy(gid)
	proxy.SendClientMsg(msgque, msg)
	return true
}

type Func func(gamer *ControlGamer, msgque antnet.IMsgQue, msg *antnet.Message) bool
type InnerFunc func(gamer *ControlGamer, msgque antnet.IMsgQue, msg *antnet.Message) interface{}

var FuncMap = map[int]Func{}
var InnerFuncMap = map[int]InnerFunc{}

func SetControlFunc(cmd, act uint8, c2s, s2c interface{}, fun Func) {
	PbParser.Register(cmd, act, c2s, s2c)
	FuncMap[antnet.CmdAct(cmd, act)] = fun
}

func SetControlInnerFunc(cmd, act uint8, c2s, s2c interface{}, fun InnerFunc) {
	PbParser.Register(cmd, act, c2s, s2c)
	InnerFuncMap[antnet.CmdAct(cmd, act)] = fun
}
