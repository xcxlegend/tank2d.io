package auth

import (
	"antnet"
	. "common"
	"model"
	"net/http"
	"proto/pb"
	"server/auth/sdk"
)

var areas map[int32]*pb.LogicServerInfo

func SendError(w http.ResponseWriter, err error) {
	m := map[string]interface{}{
		"error": antnet.GetErrId(err),
	}
	SendHttpResponse(w, m)
}
func SendHttpResponse(w http.ResponseWriter, m map[string]interface{}) {
	data, _ := antnet.JsonPack(m)
	w.Write(data)
}

func Areas(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"error": 0,
	}
	if len(areas) == 0 {
		m["error"] = ErrAreaEmpty.Id
	}
	var s []map[string]interface{}
	for _, v := range areas {
		s = append(s, map[string]interface{}{
			"id":    v.GetId(),
			"name":  v.GetName(),
			"start": v.GetStart(),
			"end":   v.GetEnd(),
		})
	}
	m["areas"] = s
	SendHttpResponse(w, m)
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var m map[string]interface{}
	channel := r.Form.Get("channel")
	if channel == "mine" {
		m = sdk.Mine.Register(r.Form.Get("name"), r.Form.Get("passwd"))
	} else {
		SendError(w, ErrChannelNotFound)
		return
	}
	SendHttpResponse(w, m)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var m map[string]interface{}
	channel := r.Form.Get("channel")
	if channel == "mine" {
		m = sdk.Mine.Login(r.Form.Get("name"), r.Form.Get("passwd"))
	} else {
		SendError(w, ErrChannelNotFound)
		return
	}

	if m["error"] == 0 {
		roles, _ := model.GetUserRoles(m["uid"].(string))
		var s []map[string]interface{}
		for _, v := range roles {
			role := map[string]interface{}{
				"id":     v.GetId(),
				"name":   v.GetName(),
				"type":   v.GetTypeId(),
				"server": v.GetServer(),
			}
			s = append(s, role)
		}

		m["roles"] = s
	}

	SendHttpResponse(w, m)
}

func NewRole(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var typeId int32 = 0
	var max int32 = 1 //玩家最多角色数量
	session := r.Form.Get("session")
	name := r.Form.Get("name")
	serverStr := r.Form.Get("server")
	if serverStr == "" {
		serverStr = "1"
	}

	server := int32(antnet.Atoi(serverStr))
	gid, err := model.GetUIdBySession(session)
	if err != nil {
		SendError(w, err)
		return
	}

	count := model.GetUserRoleCount(gid)
	if count >= max {
		SendError(w, ErrRoleCount)
		return
	}

	gm, err := model.NewModelGamer(name, server)
	if err != nil {
		SendError(w, err)
		return
	}

	role, err := model.NewRole(gid, server, typeId, max, gm.Main())
	if err != nil {
		SendError(w, err)
		return
	}

	m := map[string]interface{}{
		"role": map[string]interface{}{
			"id":     role.GetId(),
			"name":   role.GetName(),
			"server": role.GetServer(),
			"type":   role.GetTypeId(),
		},
		"error": 0,
	}
	antnet.LogInfo("new role name:%v gid:%v server:%v type:%v", name, gid, server, typeId)
	SendHttpResponse(w, m)
}

func UseRole(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	session := r.Form.Get("session")
	roleid := r.Form.Get("id")

	uid, err := model.GetUIdBySession(session)
	if err != nil {
		SendError(w, err)
		return
	}

	roles, err := model.GetUserRoles(uid)
	if err != nil {
		SendError(w, err)
		return
	}

	for _, v := range roles {
		if v.GetId() == int32(antnet.Atoi(roleid)) {
			var m map[string]interface{}
			s := getServer()
			if s == nil {
				m = map[string]interface{}{
					"error": ErrServerNotFound.Id,
				}
			} else {
				addrs := antnet.SplitStr(s.GetExtAddr(), ":")
				m = map[string]interface{}{
					"server": map[string]interface{}{
						"id":   s.GetId(),
						"net":  addrs[0],
						"ip":   addrs[1][2:],
						"port": antnet.Atoi(addrs[2]),
					},
					"error": 0,
				}
			}

			model.SetLoginSession(v.GetId(), s.GetId())
			SendHttpResponse(w, m)
			return
		}
	}

	SendError(w, ErrRoleNotFound)
}

var MinServer *pb.ServerInfo

func updateServer() {
	var min int32
	infos, err := model.UpdateServerInfo()
	if err == nil {
		StartConnect(infos, func(info *pb.ServerInfo) bool {
			return info.GetServerType() == "logic"
		})
	}

	if MinServer != nil && !IsServerOk(MinServer.GetId()) {
		MinServer = nil
	}

	if err == nil {
		for k, v := range infos {
			if v.GetServerType() == "logic" {
				if MinServer == nil && IsServerOk(k) {
					MinServer = v
					min = v.GetLoad()
				}

				if min-v.GetLoad() > 1000 && IsServerOk(k) {
					MinServer = v
					min = v.GetLoad()
				}
			}
		}
	}
}

func getServer() *pb.ServerInfo {
	return MinServer
}

//register?channel=mine&name=test&passwd=test
//newrole?session=&name=&type=&server=
//userrole?session=&id=
//login?channel=mine&name=test&passwd=test
func Start() {
	areas = model.InitServerArea()
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/newrole", NewRole)
	http.HandleFunc("/userole", UseRole)
	http.HandleFunc("/areas", Areas)
	staticHandler := http.StripPrefix("/data/", http.FileServer(http.Dir("www")))
	http.HandleFunc("/data/", func(w http.ResponseWriter, req *http.Request) {
		staticHandler.ServeHTTP(w, req)
	})

	s := &http.Server{Addr: ":5000"}
	antnet.Go(func() {
		s.ListenAndServe()
	})
	antnet.Go2(func(cstop chan struct{}) {
		select {
		case <-cstop:
			s.Close()
		}
	})

	updateServer()
	antnet.SetTimeout(3000, func(args ...interface{}) int {
		updateServer()
		return 3000
	})
}
