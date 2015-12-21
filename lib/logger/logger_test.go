package logger

import (
	"testing"
)

func TestLog(t *testing.T) {
	err := SetLogger("name", "./logs")
	if err != nil {
		t.Error(err)
	}
	log := GetLogger()
	log.Debug("debug...")
	log.Warning("warning...")
	log.Error("error...")
}
