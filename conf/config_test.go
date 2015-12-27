package conf

import (
	"testing"
)

func TestGetMysql(t *testing.T) {
	mc := GetMysqlConfig("db_realtime")
	if mc != nil {
		t.Log(mc)
	}
	if mc.Port != 3308 {
		t.Error("not equle ", 3308)
	}
}
