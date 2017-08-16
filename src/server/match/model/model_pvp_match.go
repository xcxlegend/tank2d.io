package model

import (
	"antnet"
	. "common"
	. "model"
	"proto/pb"
	"time"
)

func SetPVPSession(session string, p *pb.NotifyGamerPVPMatch) error {
	_, err := GetGlobalRedis().Set(antnet.Sprintf("session.%v", session), antnet.DBData(p), time.Duration(Config.Global.PvpMatchChoseTime+30)*time.Second).Result()
	if err != nil {
		return err
	}
	return nil
}
