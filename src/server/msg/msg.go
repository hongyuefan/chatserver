package msg

import (
	"sync"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&User{})
	Processor.Register(&ChatTo{})
	Processor.Register(&ChatLogin{})
	Processor.Register(&ReqFriend{})
	Processor.Register(&AcpFriend{})
	Processor.Register(&Response{})
}

type ChatTo struct {
	P
	OpType BussTypeId
	Token  string
	Msg    string
	To_Ids []int64
}

func (c *ChatTo) GetMsg() string {
	return c.Name + ": " + c.Msg
}

type ChatLogin struct {
	Token string
}

type P struct {
	Id   int64
	Name string
}

type ReqFriend struct {
	Token string
	P
	To_Id   int64
	To_Name string
}

type AcpFriend struct {
	Token string
	P
	To_Id   int64
	To_Name string
}

type User struct {
	P
	Friends    map[int64]string
	Blacks     map[int64]string
	Msg        string
	Count      uint64
	Token      string
	friendLock sync.RWMutex
	blackLock  sync.RWMutex
}

func NewUser(id int64, name, token string) *User {
	user := new(User)
	user.Id = id
	user.Name = name
	user.Token = token
	user.Friends = make(map[int64]string, 200)
	user.Blacks = make(map[int64]string, 200)
	return user
}

func (u *User) GetUserId() int64 {
	return u.Id
}
func (u *User) GetUserName() string {
	return u.Name
}
func (u *User) GetUserMsg() string {
	return u.Msg
}

func (u *User) GetFriends() map[int64]string {
	u.friendLock.RLock()
	defer u.friendLock.RUnlock()
	return u.Friends
}

func (u *User) GetBlacks() map[int64]string {
	u.blackLock.RLock()
	defer u.blackLock.RUnlock()
	return u.Blacks
}

func (u *User) AddFriend(p P) {
	u.friendLock.Lock()
	defer u.friendLock.Unlock()
	u.Friends[p.Id] = p.Name
}

func (u *User) AddBlack(p P) {
	u.blackLock.Lock()
	defer u.blackLock.Unlock()
	u.Blacks[p.Id] = p.Name
}

func (u *User) DelFriend(p P) {
	u.friendLock.Lock()
	defer u.friendLock.Unlock()
	delete(u.Friends, p.Id)
}

func (u *User) DelBlack(p P) {
	u.blackLock.Lock()
	defer u.blackLock.Unlock()
	delete(u.Blacks, p.Id)
}

type OnLine struct {
	Id   int64
	Name string
}

type OffLine struct {
	Id   int64
	Name string
}

type Response struct {
	Success bool
	BussId  BussTypeId
	Message string
	Data    interface{}
}

func SuccessHandler(agent gate.Agent, buss BussTypeId, data interface{}) {
	agent.WriteMsg(&Response{
		Success: true,
		BussId:  buss,
		Data:    data,
	})
}

func FailedHandler(agent gate.Agent, err error) {
	agent.WriteMsg(&Response{
		Success: false,
		Message: err.Error(),
	})
}
