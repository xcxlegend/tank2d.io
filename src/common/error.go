package common

import (
	"antnet"
)

var (
	ErrUserNameRepeated             = antnet.NewError("用户名重复", 256)
	ErrSession                      = antnet.NewError("session错误", 257)
	ErrLoginRepeated                = antnet.NewError("重复登陆", 258)
	ErrNeedLogin                    = antnet.NewError("需要登陆", 259)
	ErrSendInnerMsg                 = antnet.NewError("发送内部消息错误", 260)
	ErrInnerMsgTimeout              = antnet.NewError("内部消息超时", 261)
	ErrUserNotFound                 = antnet.NewError("用户没有找到", 262)
	ErrPasswdErr                    = antnet.NewError("密码错误", 263)
	ErrChannelNotFound              = antnet.NewError("渠道没有找到", 264)
	ErrRoleCount                    = antnet.NewError("角色数量错误", 265)
	ErrRoleNotFound                 = antnet.NewError("角色没有找到", 266)
	ErrServerNotFound               = antnet.NewError("服务器没有找到", 267)
	ErrPVPSession                   = antnet.NewError("战斗session没有找到", 268)
	ErrRequestDataIsNull            = antnet.NewError("请求数据为空", 269)
	ErrWrongPlayerId                = antnet.NewError("错误的球员ID", 270)
	ErrRepeatedPlayerId             = antnet.NewError("重复的球员ID", 271)
	ErrParametersError              = antnet.NewError("玩家id参数错误", 272)
	ErrGetPlayer                    = antnet.NewError("获取球员列表失败", 273)
	ErrLessDiamond                  = antnet.NewError("钻石不足", 274)
	ErrLessMoney                    = antnet.NewError("金币不足", 275)
	ErrMoneyType                    = antnet.NewError("错误的货币类型", 276)
	ErrPackDBError                  = antnet.NewError("背包数据库错误", 277)
	ErrGoodsCountError              = antnet.NewError("物品数量错误", 278)
	ErrReConnSession                = antnet.NewError("重连session没有找到", 279)
	ErrGoodsCountMaxError           = antnet.NewError("物品数量超出上限", 280)
	ErrRecruitPlayerType            = antnet.NewError("错误的球员招募类型", 281)
	ErrRecruitPropertyLess          = antnet.NewError("招募道具不足", 282)
	ErrFreeRecruit                  = antnet.NewError("免费招募时间未到", 283)
	ErrBreakthroughMaterialLess     = antnet.NewError("球员突破材料不足", 284)
	ErrGoodsIdIsNull                = antnet.NewError("物品ID为0", 285)
	ErrGoodsCountLess               = antnet.NewError("物品数量不足", 286)
	ErrGoodsCannotSell              = antnet.NewError("物品不能出售", 287)
	ErrSellGoodsFail                = antnet.NewError("物品出售失败", 288)
	ErrUseGoodsFail                 = antnet.NewError("使用物品失败", 289)
	ErrPackReachMaxCount            = antnet.NewError("背包达到最大容量", 290)
	ErrNotExistInCfg                = antnet.NewError("物品配置表中不存在", 291)
	ErrGoodsCannotUse               = antnet.NewError("物品不能使用", 292)
	ErrSellGoodsCount               = antnet.NewError("出售物品数量错误", 293)
	ErrEverydayRecruitCnt           = antnet.NewError("每日招募次数不足", 294)
	ErrFriendCountReachedMax        = antnet.NewError("达到好友数量上限", 295)
	ErrAlreadyIsFriend              = antnet.NewError("已经是好友", 296)
	ErrFriendRequestNameIsNull      = antnet.NewError("好友请求名字为空", 297)
	ErrFriendRequestFailed          = antnet.NewError("好友请求失败", 298)
	ErrFriendRequestNotDeal         = antnet.NewError("好友请求通知未处理", 299)
	ErrNotExistInFriendRequests     = antnet.NewError("不存在此好友请求", 300)
	ErrAddNewFriendErr              = antnet.NewError("添加新好友失败", 301)
	ErrFriendRequestCountReachedMax = antnet.NewError("达到好友请求数量上限", 302)
	ErrFriendNotExist               = antnet.NewError("好友不存在", 303)
	ErrFriendRequestExist           = antnet.NewError("好友请求已存在", 304)
	ErrGamerNotExist                = antnet.NewError("玩家不存在", 305)
	ErrClearFriendRequestFail       = antnet.NewError("清空好友请求失败", 306)
	ErrDelFriendRequestFail         = antnet.NewError("删除好友失败", 307)
	ErrMailIdIsNull                 = antnet.NewError("邮件ID为0", 308)
	ErrMailIdNotExist               = antnet.NewError("邮件ID不存在", 309)
	ErrDelMailFail                  = antnet.NewError("删除邮件失败", 310)
	ErrMailHasNotReceivedAttachment = antnet.NewError("邮件有未接收的附件", 311)
	ErrMailHasNoAttachmentReceive   = antnet.NewError("没有可接收的附件", 312)
	ErrReceiveAttachmentsFail       = antnet.NewError("附件接收失败", 313)
	ErrGamerSetMailReadFail         = antnet.NewError("设置邮件已读失败", 314)
	ErrPlayerLvReq                  = antnet.NewError("球员升级请求错误", 315)
	ErrPlayerBreakthroughReq        = antnet.NewError("球员突破请求错误", 316)
	ErrMaxPlayerLv                  = antnet.NewError("球员升级已达到最大等级", 317)
	ErrPlayerSpecialitySlot         = antnet.NewError("球员解锁特质槽位失败", 318)
	ErrPlayerUpdateSpeciality       = antnet.NewError("球员更新特质失败", 319)
	ErrGamerId                      = antnet.NewError("玩家Id错误", 320)
	ErrAreaNotFound                 = antnet.NewError("区服没有找到", 321)
	ErrAreaEmpty                    = antnet.NewError("区服为空", 322)
	ErrPlayerSpecialityGoodsLess    = antnet.NewError("球员更新特质材料不足", 323)
	ErrPlayerSpecialityIsExist      = antnet.NewError("球员初始化槽位已经存在", 324)
	ErrCommonRecruitType            = antnet.NewError("错误的普通招募类型", 325)
	ErrSeniorRecruitType            = antnet.NewError("错误的高级招募类型", 326)
	ErrPlayerCompoundId             = antnet.NewError("错误的球员碎片ID配置", 327)
	ErrGamerAlreadySign             = antnet.NewError("今日已经签到", 328)
)

var ( //自动生成 复制于antnet错误码
	ErrOk            = antnet.ErrOk            //正确
	ErrDBErr         = antnet.ErrDBErr         //数据库错误
	ErrProtoPack     = antnet.ErrProtoPack     //协议解析错误
	ErrProtoUnPack   = antnet.ErrProtoUnPack   //协议打包错误
	ErrMsgPackPack   = antnet.ErrMsgPackPack   //msgpack打包错误
	ErrMsgPackUnPack = antnet.ErrMsgPackUnPack //msgpack解析错误
	ErrPBPack        = antnet.ErrPBPack        //pb打包错误
	ErrPBUnPack      = antnet.ErrPBUnPack      //pb解析错误
	ErrJsonPack      = antnet.ErrJsonPack      //json打包错误
	ErrJsonUnPack    = antnet.ErrJsonUnPack    //json解析错误
	ErrCmdUnPack     = antnet.ErrCmdUnPack     //cmd解析错误
	ErrFileRead      = antnet.ErrFileRead      //文件读取错误
	ErrDBDataType    = antnet.ErrDBDataType    //数据库数据类型错误
	ErrNetTimeout    = antnet.ErrNetTimeout    //网络超时
	ErrErrIdNotFound = antnet.ErrErrIdNotFound //错误没有对应的错误码
)
