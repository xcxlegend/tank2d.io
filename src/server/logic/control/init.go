package control

import (
	. "common"
	"model"
	"proto/pb"
)

var MaxGamerId = 0

func init() {
	initReadWeightConfig()  //权重配置
	initNotifyHandler()     //NotifyHandler初始化
	initInnerFunc()         //InnerFunc初始化
	initBagFunc()           //背包处理函数
	initMailFunc()          //邮件处理函数
	initServerFunc()        //服务器相关处理函数
	initPlayerFunc()        //球员相关处理函数
	initFriendFunc()        //好友相关处理函数
	initGamerInfoFunc()     //玩家信息处理函数
	initDailyActivityFunc() //日常活动处理函数
}

func initNotifyHandler() {
	model.NotifyHandlerMap[model.RedisNotifyNewGamer] = OnNotifyGamerNew
	model.NotifyHandlerMap[model.RedisNotifyFriendRequest] = OnNotifyFriendRequest             //好友请求通知
	model.NotifyHandlerMap[model.RedisNotifyFriendRequestAccept] = OnNotifyFriendRequestAccept //好友请求结果通知
	model.NotifyHandlerMap[model.RedisNotifyNewPVPResult] = OnNotifyGamerNewPVPResult
	model.NotifyHandlerMap[model.RedisNotifyGamerNewMail] = OnNotifyGamerNewMail
}

func initInnerFunc() {
	RegisterRpc(GAME_CMD_INNER, GAME_CMD_INNER_PVP_MATCH_M2L, &pb.NotifyGamerPVPMatch{}, nil, HandleStartPVPMatchM2L)
	SetControlInnerFunc(GAME_CMD_INNER, GAME_CMD_INNER_NEW_FRIEND_REQ, nil, nil, HandleNewFriendReq)
	SetControlInnerFunc(GAME_CMD_INNER, GAME_CMD_INNER_LOGOUT, nil, nil, HandleLogout)
	SetControlInnerFunc(GAME_CMD_INNER, GAME_CMD_INNER_NEW_PVP_RESULT, nil, nil, HandleNewPVPResult)
	SetControlInnerFunc(GAME_CMD_INNER, GAME_CMD_INNER_NEW_GAMER_MAIL, nil, nil, HandleGamerNewMail)
}

func initServerFunc() {
	//快速断线重连
	SetControlFunc(GAME_CMD_LOGIN, GAME_CMD_LOGIN_BYRECONN, &pb.FastReconnC2S{}, &pb.GamerLoginS2C{}, HandleFastReConn)
	//发送服务器时间戳
	SetControlFunc(GAME_CMD_MAIN, GAME_CMD_MAIN_ECHO, &pb.ServerTimeC2S{}, &pb.ServerTime{}, HandleServerTime)
	//匹配
	SetControlFunc(GAME_CMD_PVP, GAME_CMD_PVP_START_MATCH, &pb.GamerStartPVPMatchC2S{}, nil, HandleStartPVPMatchC2S)
}

func initPlayerFunc() {
	//球员招募
	SetControlFunc(GAME_CMD_PLAYER, GAME_CMD_PLAYER_EMPLOY, &pb.GamerEmployPlayerC2S{}, &pb.GamerEmployPlayerS2C{}, HandlRecruitPlayer)
	//球员升级
	SetControlFunc(GAME_CMD_PLAYER, GAME_CMD_PLAYER_UPDATE_LV, &pb.GamerPlayerLvC2S{}, &pb.GamerPlayerLvS2C{}, HandlePlayerLv)
	//球员突破
	SetControlFunc(GAME_CMD_PLAYER, GAME_CMD_PLAYER_BREAKTHOUGH, &pb.GamerPlayerBreadthroughC2S{}, &pb.GamerPlayerBreadthroughS2C{}, HandlePlayerBreakthrough)
	//球员特质槽位解锁
	SetControlFunc(GAME_CMD_PLAYER, GAME_CMD_PLAYER_SPECIALITY_SLOT, &pb.GamerGetPlayerSpecialityC2S{}, &pb.GamerGetPlayerSpecialityS2C{}, HandlePlayerSpecialitySlot)
	//球员特质刷新
	SetControlFunc(GAME_CMD_PLAYER, GAME_CMD_PLAYER_SPECIALITY_UPDATE, &pb.GamerUpdatePlayerSpecialityC2S{}, &pb.GamerUpdatePlayerSpecialityS2C{}, HandlePlayerSpecialityUpdate)
}

func initBagFunc() {
	//更新背包
	SetControlFunc(GAME_CMD_BAG, GAME_CMD_BAG_GET_PACK, &pb.GamerGetPackC2S{}, &pb.GamerGetPackS2C{}, HandleGetPack)
	//使用物品
	SetControlFunc(GAME_CMD_BAG, GAME_CMD_BAG_USE_GOODS, &pb.GamerUseGoodsC2S{}, &pb.GamerUseGoodsS2C{}, HandleUseGoods)
	//出售物品
	SetControlFunc(GAME_CMD_BAG, GAME_CMD_BAG_SELL_GOODS, &pb.GamerSellGoodsC2S{}, &pb.GamerSellGoodsS2C{}, HandleSellGoods)
}

func initFriendFunc() {
	//玩家获取好友请求
	SetControlFunc(GAME_CMD_SOCIAL, GAME_CMD_SOCIAL_GET_FRIEND_REQUEST,
		&pb.GamerGetFriendRequestC2S{}, &pb.GamerGetFriendRequestS2C{}, HandleGetFriendRequest)
	//发起好友请求
	SetControlFunc(GAME_CMD_SOCIAL, GAME_CMD_SOCIAL_ADD_FRIEND, &pb.GamerAddFriendC2S{}, nil, HandleSendFriendRequest)
	//处理好友请求
	SetControlFunc(GAME_CMD_SOCIAL, GAME_CMD_SOCIAL_DEAL_FRIEND_REQUEST,
		&pb.GamerDealFriendRequestC2S{}, &pb.GamerDealFriendRequestS2C{}, HandleDealFriendRequest)
	//清空好友请求
	SetControlFunc(GAME_CMD_SOCIAL, GAME_CMD_SOCIAL_CLEAR_FRIEND_REQUEST, &pb.GamerClearFriendRequestC2S{}, nil, HandleClearFriendRequest)
	//玩家获取好友
	SetControlFunc(GAME_CMD_SOCIAL, GAME_CMD_SOCIAL_GET_FRIENDS, &pb.GamerGetAllFriendC2S{}, &pb.GamerGetAllFriendS2C{}, HandleGamerGetFriends)
	//删除好友
	SetControlFunc(GAME_CMD_SOCIAL, GAME_CMD_SOCIAL_DEL_FRIENDS, &pb.GamerDelFriendC2S{}, &pb.GamerDelFriendS2C{}, HandleDelFriend)
}

func initMailFunc() {
	//玩家获取邮件
	SetControlFunc(GAME_CMD_MAIL, GAME_CMD_MAIL_GET_MAIL, &pb.GamerGetMailC2S{}, &pb.GamerGetMailS2C{}, HandleGamerGetMail)
	//玩家删除邮件
	SetControlFunc(GAME_CMD_MAIL, GAME_CMD_MAIL_DEL_MAIL, &pb.GamerDelMailC2S{}, &pb.GamerDelMailS2C{}, HandleGamerDelMail)
	//玩家领取附件
	SetControlFunc(GAME_CMD_MAIL, GAME_CMD_MAIL_RECEIVE_ATTACHMENT, &pb.GamerReceiveAttachmentC2S{}, &pb.GamerReceiveAttachmentS2C{}, HandleGamerReceiveAttachment)
	//设置邮件已读
	SetControlFunc(GAME_CMD_MAIL, GAME_CMD_MAIL_SET_MAIL_READ, &pb.GamerReadMailC2S{}, nil, HandleReadMail)
}

func initGamerInfoFunc() {
	//玩家登录
	SetControlFunc(GAME_CMD_LOGIN, GAME_CMD_LOGIN_BYSESSION, &pb.GamerLoginC2S{}, &pb.GamerLoginS2C{}, HandleLogin)
	//设置签名
	SetControlFunc(GAME_CMD_MAIN, GAME_CMD_MAIN_SET_SIGN, &pb.GamerSetSignC2S{}, nil, HandleGamerSetSign)
	//设置头像
	SetControlFunc(GAME_CMD_MAIN, GAME_CMD_MAIN_SET_ICON, &pb.GamerSetIconC2S{}, nil, HandleGamerSetIcon)
	//设置名字
	SetControlFunc(GAME_CMD_MAIN, GAME_CMD_MAIN_SET_NAME, &pb.GamerSetNameC2S{}, nil, HandleGamerSetName)
}

func initDailyActivityFunc() {
	//日常签到
	SetControlFunc(GAME_CMD_ACTIVITY, GAME_CMD_ACTIVITY_DAILY_SIGN_IN, &pb.GamerDailySignInC2C{}, &pb.GamerDailySignInS2C{}, HandleDailySignIn)
	//累计签到
	SetControlFunc(GAME_CMD_ACTIVITY, GAME_CMD_ACTIVITY_TOTAL_SIGN_IN, &pb.GamerTotalSignInC2C{}, &pb.GamerTotalSignInS2C{}, HandleTotalSignIn)
	//补签到
	SetControlFunc(GAME_CMD_ACTIVITY, GAME_CMD_ACTIVITY_REMEDY_SIGN_IN, &pb.GamerRemedySignInC2C{}, &pb.GamerRemedySignInS2C{}, HandleRemedySignIn)
}
