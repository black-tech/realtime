package collect

import (
	"github.com/black-tech/realtime/conf"
	"github.com/ckeyer/commons/db/mysql"
	"log"
)

const (
	MS_DB_INSTANCE = "db_realtime"
)

var (
	db mysql.DBWrapper
)

func DB() mysql.DBWrapper {
	if db.DB != nil {
		return db
	}
	mc := conf.GetMysqlConfig(MS_DB_INSTANCE)
	db = mysql.ConnectMysqlDB(mc.Host,
		mc.Port,
		mc.Database,
		mc.Username,
		mc.Password)
	if db.DB == nil {
		log.Fatal("mysql connect failed")
	}
	return db
}
