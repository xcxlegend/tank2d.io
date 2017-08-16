package common

import (
	"antnet"
	"proto/pb"
	"sync"
	"time"
)

var rpcNetMap = &struct {
	sync.RWMutex
	M map[int32]antnet.IMsgQue
}{M: make(map[int32]antnet.IMsgQue)}

var matchServerId int32

func onNewRpcMsgQue(msgque antnet.IMsgQue) bool {
	msgque.SetTimeout(0)
	if msgque.GetConnType() == antnet.ConnTypeConn {
		sid := msgque.GetUser().(int32)
		antnet.LogInfo("connect to server sid:%d", sid)
		rpcNetMap.Lock()
		rpcNetMap.M[sid] = msgque
		rpcNetMap.Unlock()
	} else {

	}
	return true
}

func onDelRpcMsgQue(msgque antnet.IMsgQue) {
	if msgque.GetConnType() == antnet.ConnTypeConn {
		msgque.Reconnect(3)
	}
}

func onRpcConnectComplete(msgque antnet.IMsgQue, ok bool) bool {
	if Config.Server.Type != "auth" && ok {
		msgque.Send(antnet.NewMsg(GAME_CMD_INNER, GAME_CMD_INNER_PVP_SERVER_HELLO, 0, 0, antnet.PbData(ServerInfo)))
	}

	return true
}

func SendRpcMsg(server int32, msg *antnet.Message) bool {
	rpcNetMap.RLock()
	msgque, ok := rpcNetMap.M[server]
	rpcNetMap.RUnlock()
	if !ok || !msgque.Available() {
		return false
	}

	return msgque.Send(msg)
}

func SendRpcCallback(server int32, msg *antnet.Message, callback func(msg *antnet.Message), ws int) bool {
	rpcNetMap.RLock()
	msgque, ok := rpcNetMap.M[server]
	rpcNetMap.RUnlock()
	if !ok || !msgque.Available() {
		return false
	}
	c := make(chan *antnet.Message, 1)
	re := msgque.SendCallback(msg, c)
	if re == false {
		close(c)
	} else {
		select {
		case m := <-c:
			callback(m)
		case <-time.After(time.Duration(ws) * time.Second):
			callback(nil)
		}
	}
	return re
}

func IsServerOk(server int32) bool {
	rpcNetMap.RLock()
	msgque, ok := rpcNetMap.M[server]
	rpcNetMap.RUnlock()
	if !ok {
		return false
	}

	return msgque.Available()
}

func IsMatchServerOk() bool {
	return IsServerOk(matchServerId)
}

func SendToMatchServer(msg *antnet.Message) bool {
	return SendRpcMsg(matchServerId, msg)
}

type serverMap struct {
	sync.RWMutex
	M map[int32]*pb.ServerInfo
}

var PvpServerMap = &serverMap{M: make(map[int32]*pb.ServerInfo)}
var LogicServerMap = &serverMap{M: make(map[int32]*pb.ServerInfo)}

func HandleServerHello(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	c2s := msg.C2S().(*pb.ServerInfo)
	sid := c2s.GetId()
	antnet.LogInfo("new server id:%v", sid)
	if c2s.GetServerType() == "pvp" {
		PvpServerMap.Lock()
		PvpServerMap.M[sid] = c2s
		PvpServerMap.Unlock()
	} else if c2s.GetServerType() == "logic" {
		LogicServerMap.Lock()
		LogicServerMap.M[sid] = c2s
		LogicServerMap.Unlock()
	}

	rpcNetMap.Lock()
	rpcNetMap.M[sid] = msgque
	rpcNetMap.Unlock()

	return true
}
