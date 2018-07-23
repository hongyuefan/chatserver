package internal

import (
	"reflect"
	"server/msg"
	//	"github.com/name5566/leaf/log"
	agent "server/agent_manager"
)

func init() {
	handler(&msg.ChatTo{}, handlerChat)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerChat(args []interface{}) {

	m := args[0].(*msg.ChatTo)

	MCHandler(msg.Buss_Chat_Code, m.GetMsg(), m.To_Ids)

}

func BCHandler(buss msg.BussTypeId, data interface{}, except []int64) {
	res := &msg.Response{
		Success: true,
		BussId:  buss,
		Data:    data,
	}
	agent.MAgent.AgentBC(res, except)
}

func MCHandler(buss msg.BussTypeId, data interface{}, ks []int64) {
	res := &msg.Response{
		Success: true,
		BussId:  buss,
		Data:    data,
	}
	agent.MAgent.AgentMC(res, ks)
}

func P2PHandler(buss msg.BussTypeId, data interface{}, k int64) {
	res := &msg.Response{
		Success: true,
		BussId:  buss,
		Data:    data,
	}
	agent.MAgent.AgentP2P(res, k)
}
