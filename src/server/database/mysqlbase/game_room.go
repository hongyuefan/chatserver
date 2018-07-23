package mysqlbase

import (
	"github.com/name5566/leaf/db/orm"
)

type GameRoom struct {
	Id             int64  `orm:"column(room_id);auto"`
	RoomName       string `orm:"column(room_name);size(64)"`
	Game_Class_Id  int64  `orm:"column(game_class_id)"`
	Game_Player_Id int64  `orm:"column(game_player_id)"`
	RoomPass       string `orm:"column(room_pass);size(64)"`
	Is_Need_Pass   bool   `orm:"column(is_need_pass)"`
}

func init() {
	orm.RegisterModel(new(GameRoom))
}

func (m *GameRoom) TableName() string {
	return "game_room"
}
