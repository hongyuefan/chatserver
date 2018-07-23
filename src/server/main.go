package main

import (
	_ "server/agent_manager"
	"server/chat"
	"server/conf"
	_ "server/database/mysqlbase"
	"server/gate"
	"server/login"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		chat.Module,
		gate.Module,
		login.Module,
	)
}
