package mysqlbase

import (
	"server/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/name5566/leaf/db/orm"
)

func init() {
	orm.RegisterDataBase("default", conf.Server.SqlType, conf.Server.Conn)
}
