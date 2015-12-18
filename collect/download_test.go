package collect

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestDownload(t *testing.T) {
	return
	ret, _ := GetData(time.Now())
	fmt.Println(ret)
	ret, _ = GetData(time.Now().AddDate(0, 0, -1))
	fmt.Println(ret)
}

func TestSaveToFiel(t *testing.T) {
	return
	f, err := os.OpenFile("/tmp/data.csv", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Error(err)
	}
	defer func() {
		f.Close()
	}()
	ret, _ := GetData(time.Now())
	err = ret.SaveFile(f, -1)
	if err != nil {
		t.Error(err)
	}

	ret, _ = GetData(time.Now().AddDate(0, 0, -1))
	err = ret.SaveFile(f, -1)
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateData(t *testing.T) {
	UpdateData()
}
