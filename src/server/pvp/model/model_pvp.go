package model

import (
	"antnet"
	. "common"
	. "model"
	"proto/pb"
)

var (
	RedisScriptGetPVPSession = antnet.NewRedisScript("RedisScriptGetPVPSession", antnet.Sprintf(`
local session = 'session.' .. ARGV[1]
local data = redis.call('GET', session)
if not data then return %d end
redis.call('DEL', session)
return data
`, ErrPVPSession.Id,
	))

	RedisScriptSetPVPResult = antnet.NewRedisScript("RedisScriptSetPVPResult", antnet.Sprintf(`
local id = ARGV[1]
local time = tonumber(ARGV[2])
local data = ARGV[3]
local len = redis.call('ZCARD', 'gamer.' .. id .. '.pvp')
if len > 50 then 
	redis.call('ZREMRANGEBYRANK', 'gamer.' .. id .. '.pvp', 0, len - 50)
end
redis.call('ZADD', 'gamer.' .. id .. '.pvp', time, data)
redis.call('PUBLISH', '%d', '%d.%d.' .. data)
return %v`, Config.Server.Notify, Config.Server.Id, RedisNotifyNewPVPResult, ErrOk.Id,
	))
)

func GetPVPInfo(session string) (*pb.NotifyGamerPVPMatch, error) {
	data, err := GetGlobalRedis().ScriptStr(RedisScriptGetPVPSession, nil, session)
	if err != nil {
		return nil, err
	}

	p := &pb.NotifyGamerPVPMatch{}
	antnet.ParseDBStr(data, p)
	return p, nil
}

func SetPVPResult(gid int32, data []byte) error {
	GetGamerRedis(gid).ScriptInt64(RedisScriptSetPVPResult, nil, gid, antnet.Timestamp, data)
	return nil
}
