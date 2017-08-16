package common

import (
	"antnet"
	"proto/pb"
)

//服务器网络配置
type configServerNet struct {
	Rpc      string
	Cmd      string
	Ext      string
	ExtIpUrl string
	Timeout  int
}

//服务器配置
type configServer struct {
	Id       int32
	Notify   int32
	Type     string
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
	PlayerSelectData           map[int32]*pb.PlayerSelect           //player_select.csv表,key为招募id
	PlayerSelectPackageData    map[int32]*pb.PlayerSelectPackage    //player_select_package.csv表,key为卡池id
	PlayerSelectPackWeightData map[int32]*pb.PlayerSelectPackWeight //player_select_pack_weight.csv表,key为Index
	PlayerSdMainData           map[int32]*pb.SdPlayerMain           //sd_player_main.csv表,key为球员id
	PlayerLvUpData             map[int32]*pb.SdPlayerLvUp           //sd_player_lv_up.csv表,key为Index
	GamerItemBaseData          map[int32]*pb.ItemBase               //item_base.csv物品表
	GamerFragmentData          map[int32]*pb.FragmentCompound       //fragment_compound.csv,key为碎片id
	GamerTeamInitData          map[int32]*pb.TeamInit               //team_init.csv,key为id
	GamerDropMatching          map[int32]*pb.DropMatching           //drop_matching.csv,key为匹配类型
	PlayerSpecialityConfigData map[int32]*pb.SpecialityConfig       //speciality_config.csv,key为index
	PlayerSpecialityPoolData   map[int32]*pb.SpecialityPoolWeight   //speciality_pool_weight.csv,key为特质池id
	PlayerSpecialityPackData   map[int32]*pb.SpecialityPackage      //speciality_package.csv,key为Index
	PlayerCostAllData          map[int32]*pb.CostAll                //cost_all.csv,key为次数
}

//配置
type config struct {
	Server    *configServer
	Global    *configGlobal
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
	err, conf := antnet.ReadConfigFromCSVLie("conf/doc/global.csv", 2, 4, 1, func() interface{} {
		return &configGlobal{}
	})
	if err != nil {
		antnet.LogError("read global conf error %v", err)
		return err
	}
	Config.Global = conf.(*configGlobal)
	antnet.LogInfo("read global conf %#v", Config.Global)
	ServerInfo.Conf = pb.String(path)
	return antnet.JsonUnPack(data, Config.Server)
}

func SetConfigData(index, data int, directory string) {
	configData = make(map[string][]interface{})
	for _, v := range configInfo {
		_, s := readConfigCsv(directory+v.path, index, data, v.f)
		configData[v.path] = s
	}

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
	setConfigInfo(configInfoStr{
		path: "player_select_package.csv",
		f: func() interface{} {
			return &pb.PlayerSelectPackage{
				Index:     pb.Int32(0),
				PackageId: pb.Int32(0),
				ItemConfig: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)}}
		}})

	setConfigInfo(configInfoStr{
		path: "player_select_pack_weight.csv",
		f: func() interface{} {
			return &pb.PlayerSelectPackWeight{
				PackageId: pb.Int32(0),
				Pack_1:    pb.Int32(0),
				Weight_1:  pb.Int32(0),
				Pack_2:    pb.Int32(0),
				Weight_2:  pb.Int32(0),
				Pack_3:    pb.Int32(0),
				Weight_3:  pb.Int32(0),
				Pack_4:    pb.Int32(0),
				Weight_4:  pb.Int32(0),
				Pack_5:    pb.Int32(0),
				Weight_5:  pb.Int32(0),
				Pack_6:    pb.Int32(0),
				Weight_6:  pb.Int32(0)}
		}})

	setConfigInfo(configInfoStr{
		path: "player_select.csv",
		f: func() interface{} {
			return &pb.PlayerSelect{
				SelectId:          pb.Int32(0),
				SelectType:        pb.Int32(0),
				SelectName:        pb.String(""),
				SelectDescribe:    pb.String(""),
				OpenRequirement:   pb.Int32(0),
				RequirementValue:  pb.Int32(0),
				FreeTimesEveryday: pb.Int32(0),
				FreeCd:            pb.Int32(0),
				SelectNum:         pb.Int32(0),
				SelectBunus: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				FreeSelectPackage:    pb.Int32(0),
				FirstSelectPackId:    pb.Int32(0),
				MormalPackId:         pb.Int32(0),
				LuckyNum:             pb.Int32(0),
				LuckySelectPackId:    pb.Int32(0),
				SelectPack:           pb.Int32(0),
				HotspotPack:          pb.Int32(0),
				HotspotOdds:          pb.Int32(0),
				HotspotIncreasedOdds: pb.Int32(0),
				HotspotLuckyNum:      pb.Int32(0),
				SelectImgId:          pb.Int32(0),
				SelectPlayerImgId:    pb.Int32(0),
				OpenTime:             pb.Int32(0),
				DurationTime:         pb.Int32(0),
				ServerOnOff:          pb.Int32(0),
				NeedSelectItem: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				IsShow:       pb.Int32(0),
				LimitedTimes: pb.Int32(0)}
		}})

	setConfigInfo(configInfoStr{
		path: "item_base.csv",
		f: func() interface{} {
			return &pb.ItemBase{
				ItemId:      pb.Int32(0),
				ItemName:    pb.String(""),
				ItemType:    pb.Int32(0),
				ItemQuality: pb.Int32(0),
				IsSell:      pb.Int32(0),
				IsUse:       pb.Int32(0),
				ItemSell: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0),
				},
				UseResultType:  pb.Int32(0),
				UseResultValue: pb.Int32(0),
			}
		},
	})

	setConfigInfo(configInfoStr{
		path: "sd_player_main.csv",
		f: func() interface{} {
			return &pb.SdPlayerMain{
				Id:                       pb.Int32(0),
				PositionId:               pb.Int32(0),
				Name:                     pb.String(""),
				SchoolId:                 pb.Int32(0),
				Difficulty:               pb.Float32(0.0),
				PlayerLvUpId:             pb.Int32(0),
				BaseLvPlayer:             pb.Int32(0),
				BaseExpPlayer:            pb.Int32(0),
				PlayerIcon:               pb.String(""),
				PlayerIconSmall:          pb.String(""),
				PlayerOriginalPainting:   pb.String(""),
				ModeId:                   pb.Int32(0),
				Skill:                    []int32{0},
				EmojiId:                  pb.Int32(0),
				PlayerPropertyId:         pb.Int32(0),
				PlayerSpecialityConfigId: pb.Int32(0),
				PlayerDescribe:           pb.String("")}
		}})

	setConfigInfo(configInfoStr{
		path: "sd_player_lv_up.csv",
		f: func() interface{} {
			return &pb.SdPlayerLvUp{
				Id:                   pb.Int32(0),
				PlayerLvUpId:         pb.Int32(0),
				PlayerLv:             pb.Int32(0),
				PlayerLvUpPropertyId: pb.Int32(0),
				CritParameter: &pb.CritPara{
					Min:      pb.Int32(0),
					Max:      pb.Int32(0),
					Expected: pb.Int32(0),
					Variance: pb.Int32(0)},
				LvCrit: pb.Int32(0),
				LvCritExp: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				RequreGold: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				LvUpRequreExp: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				IsNeedGold:     pb.Int32(0),
				IsNeedMaterial: pb.Int32(0),
				LvMaterial: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				TypeLvUp: pb.Int32(0)}
		}})

	setConfigInfo(configInfoStr{
		path: "fragment_compound.csv",
		f: func() interface{} {
			return &pb.FragmentCompound{
				FragmentId:     pb.Int32(0),
				CompoundType:   pb.Int32(0),
				CompoundTarget: pb.Int32(0),
				RequireNum:     pb.Int32(0),
				RequireOther: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)}}
		}})

	setConfigInfo(configInfoStr{
		path: "team_init.csv",
		f: func() interface{} {
			return &pb.TeamInit{
				Id:         pb.Int32(0),
				BaseLvTeam: pb.Int32(1),
				BaseExpTeam: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				BaseGold: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				BaseDiamond: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				BasePlayerOne: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				BasePlayerTwo: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				BasePlayerThree: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)}}
		}})

	setConfigInfo(configInfoStr{
		path: "drop_matching.csv",
		f: func() interface{} {
			return &pb.DropMatching{
				MatchingType: pb.Int32(0),
				WinExp: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				WinGold: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				FailExp: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				FailGold: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)},
				MvpExtraGold: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)}}
		}})

	setConfigInfo(configInfoStr{
		path: "speciality_config.csv",
		f: func() interface{} {
			return &pb.SpecialityConfig{
				Index: pb.Int32(0),
				PlayerSpecialityConfigId: pb.Int32(1),
				SpecialityNum:            pb.Int32(0),
				SpecialityUnclockLv:      pb.Int32(0),
				SpecialityInitPool: &pb.SpecialityInitPool{
					IsUnlockFree: pb.Int32(0),
					InitPoolId:   pb.Int32(0)},
				SpecialityResetPool: pb.Int32(0),
				SpecialityResetItem: &pb.ItemConfig{
					Id:     pb.Int32(0),
					Number: pb.Int32(0)}}
		}})

	setConfigInfo(configInfoStr{
		path: "speciality_pool_weight.csv",
		f: func() interface{} {
			return &pb.SpecialityPoolWeight{
				SpecialityPoolId: pb.Int32(0),
				SpecialityPack_1: pb.Int32(0),
				Weight_1:         pb.Int32(0),
				SpecialityPack_2: pb.Int32(0),
				Weight_2:         pb.Int32(0),
				SpecialityPack_3: pb.Int32(0),
				Weight_3:         pb.Int32(0),
				SpecialityPack_4: pb.Int32(0),
				Weight_4:         pb.Int32(0),
				SpecialityPack_5: pb.Int32(0),
				Weight_5:         pb.Int32(0),
				SpecialityPack_6: pb.Int32(0),
				Weight_6:         pb.Int32(0)}
		}})

	setConfigInfo(configInfoStr{
		path: "speciality_package.csv",
		f: func() interface{} {
			return &pb.SpecialityPackage{
				SpecialityId:       pb.Int32(0),
				SpecialityPackId:   pb.Int32(0),
				SpecialityQuality:  pb.Int32(0),
				SpecialityIcon:     pb.String(""),
				PropertyId:         pb.Int32(0),
				PropertyValue:      pb.Int32(0),
				SpecialityDescribe: pb.String("")}
		}})

	setConfigInfo(configInfoStr{
		path: "cost_all.csv",
		f: func() interface{} {
			return &pb.CostAll{
				Times:                pb.Int32(0),
				SpecialityResetRatio: pb.Int32(0)}
		}})

}

//新添的map结构赋值在此添加,每张表的map维护一个赋值函数
func SetFunc() {
	setPlayerSelectData()           //player_select.csv表
	setPlayerSelectPackageData()    //player_select_package.csv表
	setPlayerSelectPackWeightData() //player_select_pack_weight.csv表
	setItemBaseData()               //item_base.csv物品表
	setPlayerSdMainData()           //sd_player_main.csv表
	setPlayerLvUpData()             //sd_player_lv_up.csv表
	setFragmentData()               //fragment_compound.csv表
	setTeamInitData()               //team_init.csv表
	setDropMatching()               //drop_matching.csv表
	setSpecialityConfigData()       //speciality_config.csv表
	setSpecialityPoolData()         //speciality_pool_weight.csv表
	setSpecialityPackData()         //speciality_package.csv表
	setCostAllData()                //cost_all.csv,key为次数
}

//player_select.csv表
func setPlayerSelectData() {
	Config.Doc.PlayerSelectData = make(map[int32]*pb.PlayerSelect)
	for _, v := range getConfigData("player_select.csv") {
		p, _ := v.(*pb.PlayerSelect)
		Config.Doc.PlayerSelectData[p.GetSelectId()] = p
	}
}

//player_select_package.csv表
func setPlayerSelectPackageData() {
	Config.Doc.PlayerSelectPackageData = make(map[int32]*pb.PlayerSelectPackage)
	for k, v := range getConfigData("player_select_package.csv") {
		p, _ := v.(*pb.PlayerSelectPackage)
		Config.Doc.PlayerSelectPackageData[int32(k)] = p
	}
}

//player_select_pack_weight.csv表
func setPlayerSelectPackWeightData() {
	Config.Doc.PlayerSelectPackWeightData = make(map[int32]*pb.PlayerSelectPackWeight)
	for _, v := range getConfigData("player_select_pack_weight.csv") {
		p, _ := v.(*pb.PlayerSelectPackWeight)
		Config.Doc.PlayerSelectPackWeightData[p.GetPackageId()] = p
	}
}

//item_base.csv物品表
func setItemBaseData() {
	Config.Doc.GamerItemBaseData = make(map[int32]*pb.ItemBase)
	for _, v := range getConfigData("item_base.csv") {
		item := v.(*pb.ItemBase)
		Config.Doc.GamerItemBaseData[item.GetItemId()] = item
	}
}

//sd_player_main.csv表
func setPlayerSdMainData() {
	Config.Doc.PlayerSdMainData = make(map[int32]*pb.SdPlayerMain)
	for _, v := range getConfigData("sd_player_main.csv") {
		p, _ := v.(*pb.SdPlayerMain)
		Config.Doc.PlayerSdMainData[p.GetId()] = p
	}
}

//sd_player_lv_up.csv表
func setPlayerLvUpData() {
	Config.Doc.PlayerLvUpData = make(map[int32]*pb.SdPlayerLvUp)
	for _, v := range getConfigData("sd_player_lv_up.csv") {
		p, _ := v.(*pb.SdPlayerLvUp)
		Config.Doc.PlayerLvUpData[p.GetId()] = p
	}
}

//fragment_compound.csv表
func setFragmentData() {
	Config.Doc.GamerFragmentData = make(map[int32]*pb.FragmentCompound)
	for _, v := range getConfigData("fragment_compound.csv") {
		p, _ := v.(*pb.FragmentCompound)
		Config.Doc.GamerFragmentData[p.GetFragmentId()] = p
	}
}

//team_init.csv表
func setTeamInitData() {
	Config.Doc.GamerTeamInitData = make(map[int32]*pb.TeamInit)
	for _, v := range getConfigData("team_init.csv") {
		p, _ := v.(*pb.TeamInit)
		Config.Doc.GamerTeamInitData[p.GetId()] = p
	}
}

//drop_matching.csv表
func setDropMatching() {
	Config.Doc.GamerDropMatching = make(map[int32]*pb.DropMatching)
	for _, v := range getConfigData("drop_matching.csv") {
		p, _ := v.(*pb.DropMatching)
		Config.Doc.GamerDropMatching[p.GetMatchingType()] = p
	}
}

//speciality_config.csv表
func setSpecialityConfigData() {
	Config.Doc.PlayerSpecialityConfigData = make(map[int32]*pb.SpecialityConfig)
	for _, v := range getConfigData("speciality_config.csv") {
		p, _ := v.(*pb.SpecialityConfig)
		Config.Doc.PlayerSpecialityConfigData[p.GetIndex()] = p
	}
}

//speciality_pool_weight.csv表
func setSpecialityPoolData() {
	Config.Doc.PlayerSpecialityPoolData = make(map[int32]*pb.SpecialityPoolWeight)
	for _, v := range getConfigData("speciality_pool_weight.csv") {
		p, _ := v.(*pb.SpecialityPoolWeight)
		Config.Doc.PlayerSpecialityPoolData[p.GetSpecialityPoolId()] = p
	}
}

//speciality_package.csv表
func setSpecialityPackData() {
	Config.Doc.PlayerSpecialityPackData = make(map[int32]*pb.SpecialityPackage)
	for _, v := range getConfigData("speciality_package.csv") {
		p, _ := v.(*pb.SpecialityPackage)
		Config.Doc.PlayerSpecialityPackData[p.GetSpecialityId()] = p
	}
}

//cost_all.csv表
func setCostAllData() {
	Config.Doc.PlayerCostAllData = make(map[int32]*pb.CostAll)
	for _, v := range getConfigData("cost_all.csv") {
		p, _ := v.(*pb.CostAll)
		Config.Doc.PlayerCostAllData[p.GetTimes()] = p
	}
}
