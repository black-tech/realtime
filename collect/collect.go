package collect

import (
	"fmt"
	"time"
	// "github.com/black-tech/realtime/lib/db"
	// "gopkg.in/mgo.v2"
	// "time"
	"os"
)

var (
	START_TIME  = time.Date(2012, 1, 1, 0, 0, 0, 0, time.Local)
	SOURCE_FILE = "/data/realtime.csv"
	LOG_FILE    = "/data/realtime.log"
)

// checkExistingData  检查已有数据完整性
func checkExistingData(path string, force bool) error {
	var f *os.File
	if force {
		f, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
	} else {

	}
	f.Close()
	return nil
}

func downloadAll() {
	f, err := os.OpenFile(SOURCE_FILE, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	logf, err := os.OpenFile(LOG_FILE, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		f.Close()
		logf.Close()
	}()

	for t := START_TIME; t.Before(time.Now()); t = t.Add(time.Hour * 24) {
		times := 0
	loop:
		ret, err := GetData(t)
		if err != nil {
			if times >= 3 {
				fmt.Println("err: ", err)
				fmt.Printf("%4d-%02d-%02d failed...\n", t.Year(), t.Month(), t.Day())
				logf.WriteString(fmt.Sprintf("%4d-%02d-%02d\n", t.Year(), t.Month(), t.Day()))
			}
			time.Sleep(time.Second)
			times++
			goto loop
		} else {
			err := ret.SaveFile(f)
			if err != nil {
				fmt.Println("err: ", err)
			} else {
				fmt.Printf("%4d-%02d-%02d over...\n", t.Year(), t.Month(), t.Day())
			}
		}
	}
}
