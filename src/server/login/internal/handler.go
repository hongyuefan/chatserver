package internal

import (
	"encoding/json"

	"reflect"

	db "server/database/mysqlbase"
	"server/msg"

	agent "server/agent_manager"
	"strconv"

	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.ChatLogin{}, handleLogin)
}

func handleLogin(args []interface{}) {

	var (
		err     error
		pId     int64
		play    *db.GamePlayer
		friends map[int64]string
		blacks  map[int64]string
	)

	m := args[0].(*msg.ChatLogin)

	if pId, err = getIdFromToken(m.Token); err != nil {
		goto errDeal
	}
	if play, err = db.GetPlayerById(pId); err != nil {
		err = msg.Err_Login_NotExist
		goto errDeal
	}
	agentAdd(play.Id, args[1].(gate.Agent))

	if len(play.Friends) > 0 {
		if err = json.Unmarshal([]byte(play.Friends), friends); err != nil {
			err = msg.Err_Get_Friend
			goto errDeal
		}
		msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_Chat_GetFriend_Code, friends)
	}

	if len(play.Blacks) > 0 {
		if err = json.Unmarshal([]byte(play.Blacks), blacks); err != nil {
			err = msg.Err_Get_Black
			goto errDeal
		}
		msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_Chat_GetBlack_Code, blacks)

	}
	return
errDeal:
	msg.FailedHandler(args[1].(gate.Agent), err)
}

func agentAdd(k int64, g gate.Agent) {
	agent.MAgent.InsertAgent(k, g)
	g.SetUserData(k)
}

func getIdFromToken(token string) (id int64, err error) {
	sId, err := skeleton.TokenValid(token)
	if err != nil {
		return
	}
	if id, err = strconv.ParseInt(sId, 10, 64); err != nil {
		return
	}
	return
}
