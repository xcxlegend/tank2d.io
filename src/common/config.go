package common

import (
	"antnet"
	//"proto/pb"
)

//服务器网络配置
type configServerNet struct {
	Tcp     string
	Udp     string
	Ws      string
	Timeout int
}

//服务器配置
type configServer struct {
	Id       int32
	Name     string
	Timezone int32
	Net      *configServerNet
	Redis    []string
}

//全局配置
type configGlobal struct {
	PvpDeltaTime       float32
	PvpInputDelta      float32
	PvpThreadRecv      bool
	PvpMaxReSend       int32
	PvpMinFrameTime    int32
	PvpMatchChoseTime  int32
	PvpMatchTime       int32
	EverydayUpdateTime int32 //每日刷新时间
	RoleNameRepated    int32
	NotifyMailHost     string
	NotifyMailUser     string
	NotifyMailPasswd   string
	RecvMailList       []string
}

type docConfig struct {
	//PlayerSelectData           map[int32]*pb.PlayerSelect           //player_select.csv表,key为招募id
}

//配置
type config struct {
	Server    *configServer
	Doc       docConfig
	BuildDate string
}

var Config = &config{
	Server: &configServer{},
}

var configInfo []configInfoStr
var configData map[string][]interface{}

type configInfoStr struct {
	path string
	f    antnet.GenConfigObject
}

//招募id,配置表字段select_id
const (
	COMMON_RECRUIT     = 101 //普通招募id
	SENIOR_RECRUIT     = 201 //高级招募id
	SENIOR_RECRUIT_TEN = 202 //高级招募*10id
)

const (
	PLAYER_TYPE          = 99 //球员类型
	PLAYER_FRAGMENT_TYPE = 90 //球员碎片类型
)

const (
	ID_GAMER_EXPERIENCE  = 10000001 //球队经验
	ID_MONEY             = 10000002 //金币
	ID_DIAMOND           = 10000003 //钻石
	ID_PLAYER_EXPERIENCT = 10000005 //球员经验
)

const PROBABILITY_UNIT = 10000 //概率单位:万分率

//对外可见提供需要的配置内容,访问方法Config.GetConfigData("需要读取配置的文件名")
func getConfigData(path string) []interface{} {
	if v, ok := configData[path]; ok {
		return v
	} else {
		antnet.LogInfo("err get config path:%v", path)
		return nil
	}
}

func ReadConfig(path string) error {
	data, err := antnet.ReadFile(path)

	if err != nil {
		return err
	}

	return antnet.JsonUnPack(data, Config.Server)
}

func SetConfigData(index, data int, directory string) {
	//configData = make(map[string][]interface{})
	//for _, v := range configInfo {
	//	_, s := readConfigCsv(directory+v.path, index, data, v.f)
	//	configData[v.path] = s
	//}

}

func setConfigInfo(str configInfoStr) {
	configInfo = append(configInfo, str)
}

func readConfigCsv(path string, index, data int, f antnet.GenConfigObject) (error, []interface{}) {
	err, configData := antnet.ReadConfigFromCSV(path, index, data, f)
	if err != nil {
		antnet.LogError("read config %v error %v", path, err)
		return err, nil
	}

	return err, configData
}

//需要读取的配置文件,依次在此添加,csv读取格式要求为横行读取
func AddConfig() {
	//setConfigInfo(configInfoStr{
	//	path: "player_select_package.csv",
	//	f: func() interface{} {
	//		//return &pb.PlayerSelectPackage{
	//		//	Index:     pb.Int32(0),
	//		//	PackageId: pb.Int32(0),
	//		//	ItemConfig: &pb.ItemConfig{
	//		//		Id:     pb.Int32(0),
	//		//		Number: pb.Int32(0)}}
	//	}})

}

//新添的map结构赋值在此添加,每张表的map维护一个赋值函数
func SetFunc() {
	setPlayerSelectData() //player_select.csv表

}

//player_select.csv表
func setPlayerSelectData() {
	//Config.Doc.PlayerSelectData = make(map[int32]*pb.PlayerSelect)
	//for _, v := range getConfigData("player_select.csv") {
	//	p, _ := v.(*pb.PlayerSelect)
	//	Config.Doc.PlayerSelectData[p.GetSelectId()] = p
	//}
}
