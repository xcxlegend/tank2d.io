package common

type Reason struct {
	Id  int
	Str string
}

func (r *Reason) String() string {
	return r.Str
}

var idReasonMap = map[int]*Reason{}
var reasonIdMap = map[*Reason]int{}

func NewReason(str string, id int) *Reason {
	reason := &Reason{id, str}
	idReasonMap[id] = reason
	reasonIdMap[reason] = id
	return reason
}

var (
	ReasonUnknown                  = NewReason("未知原因", 0)
	ReasonNewGamer                 = NewReason("新建角色", 1)
	ReasonGamerSign                = NewReason("玩家签到", 2)
	ReasonGamerSell                = NewReason("玩家出售", 3)
	ReasonGamerBuyPlayerUseRmb     = NewReason("使用钻石购买球员", 4)
	ReasonGamerBuyPlayerUseMoney   = NewReason("使用金币购买球员", 5)
	ReasonGamerCommonRecruit       = NewReason("普通招募", 6)
	ReasonGamerSeniorRecruit       = NewReason("高级招募", 7)
	ReasonGamerObtainBySell        = NewReason("玩家出售获得", 8)
	ReasonGamerPlayerLv            = NewReason("球员升级", 9)
	ReasonGamerPlayerBreakthrough  = NewReason("球员突破", 10)
	ReasonGamerPlayerSynthetic     = NewReason("球员合成", 11)
	ReasonGamerReceiveAttachment   = NewReason("玩家收取附件获得", 12)
	ReasonGamerPVPResult           = NewReason("战斗获得", 13)
	ReasonGamerPlayerSpecialSlot   = NewReason("球员特质槽位解锁", 14)
	ReasonGamerPlayerSpecialUpdate = NewReason("球员特质刷新", 15)
	ReasonGamerDailySignIn         = NewReason("日常签到", 16)
	ReasonGamerTotalSignIn         = NewReason("累计签到", 17)
	ReasonGamerRemedySignIn        = NewReason("补签到", 18)
	ReasonGamerUse                 = NewReason("玩家使用", 19)
)
