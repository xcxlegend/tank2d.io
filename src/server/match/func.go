package match

import (
	"antnet"
	. "common"
	"model"
	"proto/pb"
	"server/match/control"
)

func Start() {
	StartListen()
	model.UploadServerInfo()
	antnet.SetTimeout(3000, func(args ...interface{}) int {
		model.UploadServerStatus()
		return 3000
	}, nil)

	antnet.Go2(func(cstop chan struct{}) {
		control.PVPMatchingM1(cstop)
	})

	antnet.Go2(func(cstop chan struct{}) {
		control.PVPMatching(cstop, pb.PVPType_M1V1)
	})

	antnet.Go2(func(cstop chan struct{}) {
		control.PVPMatching(cstop, pb.PVPType_M2V2)
	})

	antnet.Go2(func(cstop chan struct{}) {
		control.PVPMatching(cstop, pb.PVPType_M3V3)
	})

	antnet.Go2(func(cstop chan struct{}) {
		control.PVPMatching(cstop, pb.PVPType_M5V5)
	})
}
