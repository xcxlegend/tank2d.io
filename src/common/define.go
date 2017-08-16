package common

const (
	GAME_CMD_LOGIN      = 1   //登陆命令
	GAME_CMD_MAIN       = 3   //主命令
	GAME_CMD_PVE        = 4   //PVE战斗
	GAME_CMD_PVP        = 5   //PVP战斗
	GAME_CMD_CARD       = 6   //卡牌操作
	GAME_CMD_LEAGUE     = 10  //联盟命令
	GAME_CMD_MAIL       = 11  //邮件命令
	GAME_CMD_SOCIAL     = 12  //社交命令
	GAME_CMD_BUILD      = 13  //建筑命令
	GAME_CMD_SOLDIER    = 14  //士兵命令
	GAME_CMD_CARD_GROUP = 15  //卡组命令
	GAME_CMD_BAG        = 16  //背包命令
	GAME_CMD_TECH       = 17  //科技命令
	GAME_CMD_MAP        = 18  //地图命令
	GAME_CMD_COMBAT     = 19  //战斗命令
	GAME_CMD_EFFECT     = 20  //效果命令
	GAME_CMD_QUEST      = 21  //任务
	GAME_CMD_PLAYER     = 22  //球员
	GAME_CMD_ACTIVITY   = 23  //活动
	GAME_CMD_INNER      = 200 //
	GAME_CMD_NOTIFY     = 253 //通知消息
	GAME_CMD_TEST       = 254 //测试
)

const (
	GAME_CMD_LOGIN_BYSESSION = 1 //登陆
	GAME_CMD_LOGIN_BYRECONN  = 2 //重新登陆
)

const (
	GAME_CMD_INNER_LOGOUT           = 1
	GAME_CMD_INNER_PVP_MATCH_L2M    = 2
	GAME_CMD_INNER_PVP_MATCH_M2L    = 3
	GAME_CMD_INNER_PVP_MATCH_STOP   = 4
	GAME_CMD_INNER_PVP_SERVER_HELLO = 5
	GAME_CMD_INNER_NEW_PVP_RESULT   = 6
	GAME_CMD_INNER_NEW_FRIEND_REQ   = 7
	GAME_CMD_INNER_NEW_GAMER_MAIL   = 8
	GAME_CMD_INNER_NEW_FRIEND       = 9
)

const (
	GAME_CMD_MAIN_ECHO             = 1  //心跳
	GAME_CMD_MAIN_GET_INFO         = 2  //重新获取信息信息和登陆信息一致
	GAME_CMD_MAIN_GET_SERVER_TIME  = 3  //获取服务器时间
	GAME_CMD_MAIN_FIND_GAMER       = 4  //查找玩家
	GAME_CMD_MAIN_GET_FRIEND       = 5  //获取好友列表
	GAME_CMD_MAIN_ADD_FRIEND       = 6  //添加好友
	GAME_CMD_MAIN_FRIREQ_ECHO      = 7  //好友请求回复
	GAME_CMD_MAIN_GET_COMBAT_GIFT  = 9  //获取战斗奖励
	GAME_CMD_MAIN_GET_COMBAT       = 10 //获取战斗信息
	GAME_CMD_MAIN_FRIEND_TALK      = 11 //聊天
	GAME_CMD_MAIN_GET_LB_MAIN      = 12 //获取主排行榜
	GAME_CMD_MAIN_GET_OTHER        = 13 //获取其他玩家信息
	GAME_CMD_MAIN_CHANGE_NAME      = 14 //换名字
	GAME_CMD_MAIN_GET_ACHIEVE_GIFT = 15
	GAME_CMD_MAIN_GET_ACHIEVE      = 16 //获得成就
	GAME_CMD_MAIN_SET_SIGN         = 17 //设置签名
	GAME_CMD_MAIN_SET_ICON         = 18 //设置头像
	GAME_CMD_MAIN_SET_NAME         = 19 //设置名字
)

const (
	GAME_CMD_PVP_START_MATCH = 1
	GAME_CMD_PVP_SYNC_C2S    = 2
)

const (
	GAME_CMD_NOTIFY_OTHER_LOGIN        = 1 //其他地方登陆
	GAME_CMD_NOTIFY_NEW_FRIEND_REQUEST = 2 //新的好友请求
	GAME_CMD_NOTIFY_NEW_FRIEND         = 3 //新的好友
	GAME_CMD_NOTIFY_NEW_FRIEND_TALK    = 4 //新的聊天
	GAME_CMD_NOTIFY_NEW_ACHIEVE        = 5 //新的成就
	GAME_CMD_NOTIFY_PVP_INVITE         = 6 //pvp邀请
	GAME_CMD_NOTIFY_PVP_DENY_INVITE    = 7 //拒绝pvp邀请
	GAME_CMD_NOTIFY_PVP_MATCH_OK       = 8
	GAME_CMD_NOTIFY_PVP_SYNC           = 9  //pvp同步消息
	GAME_CMD_NOTIFY_NEW_GAMER_MAIL     = 10 // 新的邮件通知
)

const (
	GAME_CMD_PLAYER_OWNER             = 1 //获取已拥有球员
	GAME_CMD_PLAYER_EMPLOY            = 2 //球员招募
	GAME_CMD_PLAYER_UPDATE_LV         = 3 //球员升级
	GAME_CMD_PLAYER_BREAKTHOUGH       = 4 //球员突破
	GAME_CMD_PLAYER_SPECIALITY_SLOT   = 5 //球员特质槽解锁
	GAME_CMD_PLAYER_SPECIALITY_UPDATE = 6 //球员特质刷新
)

const (
	GAME_CMD_BAG_GET_PACK   = 1 //获取背包
	GAME_CMD_BAG_USE_GOODS  = 2 //使用物品
	GAME_CMD_BAG_SELL_GOODS = 3 //出售物品
)

const (
	GAME_CMD_SOCIAL_ADD_FRIEND           = 1 //发起好友请求
	GAME_CMD_SOCIAL_DEAL_FRIEND_REQUEST  = 2 //处理好友请求
	GAME_CMD_SOCIAL_CLEAR_FRIEND_REQUEST = 3 //清空好友请求
	GAME_CMD_SOCIAL_GET_FRIEND_REQUEST   = 4 //获取好友请求
	GAME_CMD_SOCIAL_GET_FRIENDS          = 5 //获取好友
	GAME_CMD_SOCIAL_DEL_FRIENDS          = 6 //删除好友
)

const (
	GAME_CMD_MAIL_GET_MAIL           = 1 //玩家获取邮件
	GAME_CMD_MAIL_DEL_MAIL           = 2 //玩家删除邮件
	GAME_CMD_MAIL_RECEIVE_ATTACHMENT = 3 //玩家领取附件
	GAME_CMD_MAIL_SET_MAIL_READ      = 4 //设置邮件已读
)

const (
	GAME_CMD_ACTIVITY_DAILY_SIGN_IN  = 1 //玩家日常签到
	GAME_CMD_ACTIVITY_TOTAL_SIGN_IN  = 2 //玩家累计签到
	GAME_CMD_ACTIVITY_REMEDY_SIGN_IN = 3 //玩家补签到
)
