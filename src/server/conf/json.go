package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/name5566/leaf/log"
)

var Server struct {
	LogLevel        string
	LogPath         string
	WSAddr          string
	CertFile        string
	KeyFile         string
	TCPAddr         string
	SqlType         string
	Conn            string
	Email_Stmp      string
	Email_From_Name string
	Email_Port      int
	Email_Sender    string
	Email_Pass      string
	AppId           string
	AppKey          string
	TplId           int
	Token_Salt      string
	Token_Exp       int64
	SessionExpire   int64
	SessionCheck    int64
	MaxConnNum      int
	ConsolePort     int
	ProfilePath     string
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
