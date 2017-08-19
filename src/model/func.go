package model

import (
	"antnet"
	. "common"
)

const (
	RedisNotifyNone = iota
	RedisNotifyNewGamer
	RedisNotifyGamerOnline
	RedisNotifyGamerUpdate
	RedisNotifyGamerOffline
	RedisNotifyFriendRequest
	RedisNotifyFriendRequestAccept
	RedisNotifyNewPVPResult
	RedisNotifyGamerPlayerLvUpdate         //恭喜xx玩家达到xx等级
	RedisNotifyGamerPlayerInitSpeciality   //恭喜xx玩家解锁槽位获得xx特质
	RedisNotifyGamerPlayerUpdateSpeciality //恭喜xx玩家更新特质获得xx特质
	RedisNotifyGamerNewMail
)

var RedisManager *antnet.RedisManager
var GamersPerReis int = 300000

func GetGamerRedis(gid int32) *antnet.Redis {
	return RedisManager.GetByRid(int(gid) / GamersPerReis)
}

func GetGlobalRedis() *antnet.Redis {
	return RedisManager.GetGlobal()
}

func Start() {
	config := []*antnet.RedisConfig{}
	for _, v := range Config.Server.Redis {
		ss := antnet.SplitStr(v, "@")
		c := &antnet.RedisConfig{Passwd: antnet.TrimStr(ss[0]), Addr: antnet.TrimStr(ss[1]), PoolSize: 10}
		config = append(config, c)
	}

	RedisManager = antnet.NewRedisManager(config)

}
