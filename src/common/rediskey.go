package common

/**
@title redis key collections
@author Legend.Xie <xcx_legender@qq.com>
@create 2017-08-03
@description key集合 便于 统一管理查看和使用
*/
const (
	REDIS_KEY_SERVER_INFO = "server.info" // HASH  服务器状态

	REDIS_KEY_IDGENS = "idgen" // HASH key=type value=id ID自增集

	REDIS_KEY_USERS = "user" // HASH key=username value=<data>user 用户登录信息

	REDIS_KEY_USER_ROLES = "role.%v" // role.<uidstr> HASH key=gid value=<data>role 用户角色信息

	REDIS_KEY_GAMERS = "gamer.%v" // gamer.<id> HASH . key=main value=<data>main  gamer主信息

	REDIS_KEY_MAP_GAMER_NAME2ID = "names.gamer" // HASH . key=name value=id 名字映射关系

	REDIS_KEY_GAMER_NAMECOUNT = "names.gamer.count'" // HASH . key=name value=count 重复名称带#count

	REDIS_KEY_SUB_SERVER_SESSION = "session.%v" // session.<id> STRING value=server_id 作为登录战斗服验证

	REDIS_KEY_GAMER_PACK = "gamer.%v.pack" // gamer.<gid>.pack HASH key=item_id value=item gamer背包

	REDIS_KEY_GAMER_PLAYERS = "gamer.%v.player" // gamer.<gid>.player HASH key=player_id value=player gamer的球员

	REDIS_KEY_GAMER_MAILS = "gamer.%v.mail" // gamer.<gid>.mail HASH key=mail_id value=<data>mail 个人邮件信息

	REDIS_KEY_SYS_MAILS = "server.%v.sysmail" // server.<lid>.sysmail HASH key=? value=<data>mail 系统邮件信息

	REDIS_KEY_GAMER_FRIENDS = "gamer.%v.friend" // gamer.<gid>.friend HASH KEY=gid value=<data>friend 个人好友信息

	REDIS_KEY_GAMER_FRIEND_REQUESTS = "gamer.%v.friend.request" // gamer.<gid>.friend.request HASH key=gid value=<data>friend_request 好友请求信息

	REDIS_KEY_GAMER_PVP_RESULT = "gamer.%v.pvp" // gamer.<gid>.pvp SORT-SETS value=<data>result score=timestamp 个人结果集合

)
