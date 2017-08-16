package common

import (
	"antnet"
	"io/ioutil"
	"net/http"
	"proto/pb"
	"sync"

	"github.com/bitly/go-simplejson"
)

var PbParser = &antnet.Parser{Type: antnet.ParserTypePB}
var CmdParser = &antnet.Parser{Type: antnet.ParserTypeCmd}

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
var CmdHandler = &msgHandler{}
var RpcHandler = &msgHandler{}

func GetExtIpByUrl() string {
	resp, err := http.Get(Config.Server.Net.ExtIpUrl)
	if err != nil {

	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	json, err := simplejson.NewJson(body)
	if err != nil {

	}
	ip, _ := json.Get("eip").String()
	antnet.LogInfo("get %v from %v", ip, Config.Server.Net.ExtIpUrl)
	return ip
}

func StartExtListen() error {
	addrs := antnet.SplitStr(Config.Server.Net.Ext, "://")
	if addrs[1][0] == ':' {
		if Config.Server.Net.ExtIpUrl != "" {
			ServerInfo.ExtAddr = pb.String(addrs[0] + "://" + GetExtIpByUrl() + addrs[1])
		} else {
			for _, ip := range antnet.GetSelfExtraIp() {
				ServerInfo.ExtAddr = pb.String(addrs[0] + "://" + ip + addrs[1])
			}
			if ServerInfo.ExtAddr == nil {
				for _, ip := range antnet.GetSelfIntraIp() {
					if ip != "127.0.0.1" {
						ServerInfo.ExtAddr = pb.String(addrs[0] + "://" + ip + addrs[1])
						break
					}
				}
			}
		}
	} else {
		ServerInfo.ExtAddr = pb.String(Config.Server.Net.Ext)
	}
	return antnet.StartServer(Config.Server.Net.Ext, antnet.MsgTypeMsg, ExtHandler, PbParser)
}

//
//func StartCmdListen() error {
//	addrs := antnet.SplitStr(Config.Server.Net.Cmd, "://")
//	if addrs[1][0] == ':' {
//		for _, ip := range antnet.GetSelfIntraIp() {
//			if ip == "127.0.0.1" {
//				continue
//			}
//			ServerInfo.CmdAddr = pb.String(addrs[0] + "://" + ip + addrs[1])
//			err := antnet.StartServer(addrs[0]+"://"+ip+addrs[1], antnet.MsgTypeCmd, CmdHandler, CmdParser)
//			if err == nil {
//				err = antnet.StartServer(addrs[0]+"://127.0.0.1"+addrs[1], antnet.MsgTypeCmd, CmdHandler, CmdParser)
//			}
//			return err
//		}
//	}
//	ServerInfo.CmdAddr = pb.String(Config.Server.Net.Cmd)
//	return antnet.StartServer(Config.Server.Net.Cmd, antnet.MsgTypeCmd, CmdHandler, CmdParser)
//}
//
//func StartRpcListen() error {
//	addrs := antnet.SplitStr(Config.Server.Net.Rpc, "://")
//	if addrs[1][0] == ':' {
//		for _, ip := range antnet.GetSelfIntraIp() {
//			if ip == "127.0.0.1" {
//				continue
//			}
//			ServerInfo.RpcAddr = pb.String(addrs[0] + "://" + ip + addrs[1])
//			return antnet.StartServer(addrs[0]+"://"+ip+addrs[1], antnet.MsgTypeMsg, RpcHandler, PbParser)
//		}
//	}
//	ServerInfo.RpcAddr = pb.String(Config.Server.Net.Rpc)
//	return antnet.StartServer(Config.Server.Net.Rpc, antnet.MsgTypeMsg, RpcHandler, PbParser)
//}

func StartListen() error {
	//StartCmdListen()
	//StartRpcListen()
	StartExtListen()
	ServerInfo.Id = pb.Int32(int32(Config.Server.Id))
	ServerInfo.ServerType = pb.String(Config.Server.Type)
	ServerInfo.LastTick = pb.Int64(antnet.Timestamp)
	return nil
}

var rpcAddrMap = &struct {
	sync.RWMutex
	M map[string]int
}{M: make(map[string]int)}

func StartConnect(m map[int32]*pb.ServerInfo, filter func(info *pb.ServerInfo) bool) error {
	for k, v := range m {
		if k == Config.Server.Id {
			continue
		} else if filter(v) {
			rpcAddrMap.Lock()
			_, ok := rpcAddrMap.M[v.GetRpcAddr()]
			if ok {
				rpcAddrMap.Unlock()
				continue
			}

			rpcAddrMap.M[v.GetRpcAddr()] = 1
			rpcAddrMap.Unlock()
			addrs := antnet.SplitStr(v.GetRpcAddr(), "://")
			antnet.StartConnect(addrs[0], addrs[1], antnet.MsgTypeMsg, RpcHandler, PbParser, k)
			if v.GetServerType() == "match" {
				matchServerId = k
			}
		}
	}

	return nil
}
