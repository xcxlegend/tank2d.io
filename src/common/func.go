package common

import (
	"antnet"
	"sync/atomic"
	"unsafe"
)

type UUID struct {
	Timestamp int32
	ServerId  int16
	Index     int16
}

var uuidIndex int16 = 0

func GetUUID() int64 {
	uuidIndex++
	uuid := UUID{int32(antnet.Timestamp), int16(Config.Server.Id), uuidIndex}
	return *((*int64)(unsafe.Pointer(&uuid)))
}

func AddOnlineCount() int32 {
	return atomic.AddInt32(&OnlineCount, 1)
}
func SubOnlineCount() int32 {
	return atomic.AddInt32(&OnlineCount, -1)
}
func SetOnlineCount(c int32) int32 {
	OnlineCount = c
	return c
}
