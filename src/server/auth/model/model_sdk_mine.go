package model

import (
	"antnet"
	. "common"
	. "model"
	"proto/pb"
)

var (
	RedisScriptSdkMineRegister = antnet.NewRedisScript("RedisScriptSdkMineRegister", antnet.Sprintf(`
local user = cmsgpack.unpack(ARGV[1])
local ok = redis.call('HSETNX', 'user', user.Name, '')
if ok == 0 then 
	return %d
end	
user.Id = 'mine_' .. redis.call('HINCRBY', 'idgen', 'user', 1)
local data = cmsgpack.pack(user)
redis.call('HSET', 'user', user.Name, data)
redis.call('SETEX', 'session.' .. ARGV[2], 300, user.Id)
return data
`, ErrUserNameRepeated.Id,
	))

	RedisScriptSdkMineLogin = antnet.NewRedisScript("RedisScriptSdkMineLogin", antnet.Sprintf(`
local name = ARGV[1]
local passwd = ARGV[2]
local session = ARGV[3] 
local data = redis.call('HGET', 'user', name)
if not data then
	return %d
end	
local user = cmsgpack.unpack(data)
if user.Passwd ~= passwd then
	return %d
end
redis.call('SETEX', 'session.' .. session, 300, user.Id)
return data
`, ErrUserNotFound.Id, ErrPasswdErr.Id,
	))
)

type modelUserInfo struct {
	*pb.UserInfo
	antnet.RedisModel
}

func Register(name, passwd, session string) (*modelUserInfo, error) {
	user := &modelUserInfo{
		UserInfo: &pb.UserInfo{
			Name:   pb.String(name),
			Passwd: pb.String(passwd),
		},
	}

	_, err := GetGlobalRedis().ScriptStr(RedisScriptSdkMineRegister, nil, user.DBStr(user.UserInfo), session)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func SDKMineLogin(name, passwd, session string) (*modelUserInfo, error) {
	user := &modelUserInfo{UserInfo: &pb.UserInfo{}}
	data, err := GetGlobalRedis().ScriptStr(RedisScriptSdkMineLogin, nil, name, passwd, session)
	if err != nil {
		return nil, ErrUserNotFound
	}

	user.ParseDBStr(data, user.UserInfo)
	if user.GetPasswd() != passwd {
		return nil, ErrPasswdErr
	}
	return user, nil
}
