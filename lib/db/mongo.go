package db

import (
	"fmt"
	"github.com/black-tech/realtime/conf"
	mgopkg "gopkg.in/mgo.v2"
)

var (
	database *mgopkg.Database
	session  *mgopkg.Session
)

const (
	MGODB_INIT = "initial"  // 原始数据
	MGODB_ANAL = "analysis" // 分析结果
	MGODB_TMP  = "tmp"      // 临时数据
)

func GetDB(name string) (db *mgopkg.Database, err error) {
	mconf, err := conf.GetMGOConfig(name)
	if err != nil {
		return
	}
	connstr := fmt.Sprintf("mongodb://%s:%d", mconf.Host, mconf.Port)

	s, err := mgopkg.Dial(connstr)
	if err != nil {
		return
	}

	return s.DB(mconf.Database), nil
}

type execHandler func(db *mgo.Database) error

func Exec(name string, callback execHandler) error {
	db, err := GetDB(name)
	if err != nil {
		return err
	}
	defer ss.Close()
	return callback(db)
}
