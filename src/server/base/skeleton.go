package base

import (
	"server/conf"

	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/module"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		AsynCallLen:        conf.AsynCallLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
		Email_Stmp:         conf.Server.Email_Stmp,
		Email_From_Name:    conf.Server.Email_From_Name,
		Email_Port:         conf.Server.Email_Port,
		Email_Sender:       conf.Server.Email_Sender,
		Email_Pass:         conf.Server.Email_Pass,
		AppId:              conf.Server.AppId,
		AppKey:             conf.Server.AppKey,
		TplId:              conf.Server.TplId,
		Token_Salt:         conf.Server.Token_Salt,
		Token_Exp:          conf.Server.Token_Exp,
		SessionExpire:      conf.Server.SessionExpire,
		SessionCheck:       conf.Server.SessionCheck,
	}
	skeleton.Init()
	return skeleton
}
