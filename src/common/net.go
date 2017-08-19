package common

import (
	"antnet"
	//"io/ioutil"
	//"net/http"
	//"proto/pb"
	//"sync"
	//"github.com/bitly/go-simplejson"
	"proto/pb"
)

var PbParser = &antnet.Parser{Type: antnet.ParserTypePB}
var CmdParser = &antnet.Parser{Type: antnet.ParserTypeCmd}
var JsonParser = &antnet.Parser{Type: antnet.ParserTypeJson}

type msgHandler struct {
	antnet.DefMsgHandler
	NewMsgQue       func(msgque antnet.IMsgQue) bool                      //新的消息队列
	DelMsgQue       func(msgque antnet.IMsgQue)                           //消息队列关闭
	ProcessMsg      func(msgque antnet.IMsgQue, msg *antnet.Message) bool //默认的消息处理函数
	ConnectComplete func(msgque antnet.IMsgQue, ok bool) bool             //连接成功
}

func (r *msgHandler) OnNewMsgQue(msgque antnet.IMsgQue) bool {
	if r.NewMsgQue != nil {
		return r.NewMsgQue(msgque)
	}
	return true
}
func (r *msgHandler) OnDelMsgQue(msgque antnet.IMsgQue) {
	if r.DelMsgQue != nil {
		r.DelMsgQue(msgque)
	}
}
func (r *msgHandler) OnProcessMsg(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	if r.ProcessMsg != nil {
		return r.ProcessMsg(msgque, msg)
	}

	return true
}
func (r *msgHandler) OnConnectComplete(msgque antnet.IMsgQue, ok bool) bool {
	if r.ConnectComplete != nil {
		return r.ConnectComplete(msgque, ok)
	}
	return true
}

var ExtHandler = &msgHandler{}

//
func StartExtListen() error {
	var err error
	err = antnet.StartServer(Config.Server.Net.Tcp, antnet.MsgTypeMsg, ExtHandler, PbParser)
	return err
}

//
func StartListen() error {
	StartExtListen()
	return nil
}

//
//var rpcAddrMap = &struct {
//	sync.RWMutex
//	M map[string]int
//}{M: make(map[string]int)}
//
//func StartConnect(m map[int32]*pb.ServerInfo, filter func(info *pb.ServerInfo) bool) error {
//	for k, v := range m {
//		if k == Config.Server.Id {
//			continue
//		} else if filter(v) {
//			rpcAddrMap.Lock()
//			_, ok := rpcAddrMap.M[v.GetRpcAddr()]
//			if ok {
//				rpcAddrMap.Unlock()
//				continue
//			}
//
//			rpcAddrMap.M[v.GetRpcAddr()] = 1
//			rpcAddrMap.Unlock()
//			addrs := antnet.SplitStr(v.GetRpcAddr(), "://")
//			antnet.StartConnect(addrs[0], addrs[1], antnet.MsgTypeMsg, RpcHandler, PbParser, k)
//			if v.GetServerType() == "match" {
//				matchServerId = k
//			}
//		}
//	}
//
//	return nil
//}
