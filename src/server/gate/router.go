package gate

import (
	"server/chat"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.ChatTo{}, chat.ChanRPC)
	msg.Processor.SetRouter(&msg.ChatLogin{}, login.ChanRPC)
}
