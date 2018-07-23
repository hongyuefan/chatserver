package mysqlbase

import (
	"fmt"

	"github.com/name5566/leaf/db/orm"
)

type GamePlayer struct {
	Id            int64  `orm:"column(player_id);auto"`
	MobileOrEmail string `orm:"column(player_mobileoremail);size(32);null"`
	Password      string `orm:"column(player_password);size(64)"`
	Nickname      string `orm:"column(player_nickname);size(32);null"`
	Avatar        string `orm:"column(player_avatar);size(256);null"`
	Paypass       string `orm:"column(player_paypass);size(128);null"`
	Friends       string `orm:"column(player_friends);size(1024);null"`
	Blacks        string `orm:"column(player_blacks);size(1024);null"`
	UserType      int8   `orm:"column(player_usertype);size(32);null"`
	Createtime    int64  `orm:"column(player_createtime)"`
	Lastime       int64  `orm:"column(player_lastime)"`
	Isdel         int8   `orm:"column(player_isdel);null"`
}

func init() {
	orm.RegisterModel(new(GamePlayer))
}

func (m *GamePlayer) TableName() string {
	return "game_player"
}

func AddPlayer(m *GamePlayer) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPlayerById retrieves Player by Id. Returns error if
// Id doesn't exist
func GetPlayerById(id int64) (v *GamePlayer, err error) {
	o := orm.NewOrm()
	v = &GamePlayer{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func UpdatePlayerById(m *GamePlayer, cols ...string) (err error) {
	o := orm.NewOrm()
	v := GamePlayer{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m, cols...); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePlayer deletes Player by Id and returns error if
// the record to be deleted doesn't exist
func DeletePlayer(id int64) (err error) {
	o := orm.NewOrm()
	v := GamePlayer{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GamePlayer{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetPlayerBy(m *GamePlayer, cols ...string) (err error) {
	o := orm.NewOrm()
	return o.Read(m, cols...)
}
