package internal

import (
	"reflect"
	agent "server/agent_manager"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	if reflect.ValueOf(a.UserData()).Kind() == reflect.Int64 {
		agent.MAgent.DelAgent(a.UserData().(int64))
	}
}
