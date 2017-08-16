package main

import (
	"antnet"
	. "common"
	_ "control"
	"flag"
	"model"
	"server/auth"
	"server/logic"
	"server/match"
	"server/pvp"
)

func main() {
	Config.BuildDate = "2017.8.16"
	antnet.Printf("build date:%v\n", Config.BuildDate)
	path := flag.String("f", "conf/pvp.json", "run config")
	daemon := flag.Int("d", 0, "1 for daemon")
	flag.Parse()
	if *daemon == 1 {
		antnet.Daemon([]string{"-d"})
		antnet.Println("go daemon")
		return
	}

	logPath := *path
	logPath = antnet.ReplaceStr(logPath, "conf", "log")
	logPath = antnet.ReplaceStr(logPath, ".json", ".log")
	antnet.DefLog.SetLogger(&antnet.FileLogger{Path: logPath, Ln: true, Timeout: 0, MaxSize: 1024 * 1024 * 10, OnFull: nil}, true)
	antnet.DefLog.SetLevel(antnet.LogLevelInfo)

	antnet.LogInfo("build date:%v", Config.BuildDate)
	ReadConfig(*path)
	model.Start()
	//if Config.Server.Type == "auth" {
	auth.Start()
	//} else if Config.Server.Type == "pvp" {
	pvp.Start()
	//} else if Config.Server.Type == "logic" {
	logic.Start()
	//} else if Config.Server.Type == "match" {
	//match.Start()
	//}

	antnet.WaitForSystemExit()
}
