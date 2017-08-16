package model

import (
	"antnet"
	. "common"
	"proto/pb"
	"time"
)

var (
	RedisScriptNewRole = antnet.NewRedisScript("RedisScriptNewRole", antnet.Sprintf(`
local uid = ARGV[1]
local gid = ARGV[2]
local data = ARGV[3]
local max = tonumber(ARGV[4])
local count = redis.call('HLEN', 'role.' .. uid)
if count >= max then
	return %d
end	
redis.call('HSET', 'role.' .. uid, gid, data)
return 0
`, ErrRoleCount.Id,
	))
)

type modelUserRoleInfo struct {
	*pb.UserRoleInfo
	antnet.RedisModel
}

func GetUIdBySession(session string) (string, error) {
	uid, err := GetGlobalRedis().Get(antnet.Sprintf("session.%s", session)).Result()
	if err != nil {
		return uid, ErrSession
	}
	return uid, err
}

//查询玩家角色数量
func GetUserRoleCount(uid string) int32 {
	count, _ := GetGlobalRedis().HLen(antnet.Sprintf("role.%s", uid)).Result()
	return int32(count)
}

//获得玩家所有角色
func GetUserRoles(uid string) ([]*modelUserRoleInfo, error) {
	var roles []*modelUserRoleInfo
	data, err := GetGlobalRedis().HVals(antnet.Sprintf("role.%s", uid)).Result()
	for _, v := range data {
		m := &modelUserRoleInfo{UserRoleInfo: &pb.UserRoleInfo{}}
		if m.ParseDBStr(v, m.UserRoleInfo) {
			roles = append(roles, m)
		} else {

		}
	}

	return roles, err
}

func NewRole(uid string, server, typeId, max int32, main *modelGamerMain) (*modelUserRoleInfo, error) {
	m := &modelUserRoleInfo{
		UserRoleInfo: &pb.UserRoleInfo{
			Id:     main.Id,
			Name:   main.Name,
			Server: pb.Int32(server),
			Ctime:  main.TimeRecord.Create,
			Uid:    pb.String(uid),
			TypeId: pb.Int32(typeId),
		},
	}

	code, err := GetGlobalRedis().ScriptInt64(RedisScriptNewRole, nil, uid, main.GetId(), m.DBStr(m.UserRoleInfo), max)
	if err != nil {
		return nil, err
	}
	if code > 0 {
		return nil, antnet.GetError(uint16(code))
	}

	return m, nil
}

func SetLoginSession(gid int32, server int32) error {
	_, err := GetGamerRedis(gid).Set(antnet.Sprintf("session.%v", gid), server, time.Duration(10)*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}

func InitServerArea() map[int32]*pb.LogicServerInfo {
	m := map[int32]*pb.LogicServerInfo{}
	data, err := GetGlobalRedis().HVals("server.area").Result()
	if err != nil {
		return m
	}

	for _, v := range data {
		s := &pb.LogicServerInfo{}
		if antnet.ParseDBStr(v, s) {
			m[s.GetId()] = s
		}
	}

	return m
}
