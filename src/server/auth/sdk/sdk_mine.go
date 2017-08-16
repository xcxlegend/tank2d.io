package sdk

import (
	"antnet"
	"server/auth/model"
)

type mine struct{}

var Mine = &mine{}

func (r *mine) Register(name, passwd string) map[string]interface{} {
	session := antnet.MD5Str(antnet.Sprintf("%s%v", name, antnet.Timestamp))
	_, err := model.Register(name, passwd, session)
	if err != nil {
		return map[string]interface{}{
			"error":   antnet.GetErrId(err),
			"session": nil,
		}

	}

	return map[string]interface{}{
		"error":   0,
		"session": session,
	}
}

func (r *mine) Login(name, passwd string) map[string]interface{} {
	session := antnet.MD5Str(antnet.Sprintf("%s%v", name, antnet.Timestamp))
	user, err := model.SDKMineLogin(name, passwd, session)
	if err != nil {
		return map[string]interface{}{
			"error":   antnet.GetErrId(err),
			"session": nil,
		}
	}
	return map[string]interface{}{
		"error":   0,
		"uid":     user.GetId(),
		"session": session,
	}
}
