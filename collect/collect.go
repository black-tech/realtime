package collect

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	// "github.com/black-tech/realtime/lib/db"
	// "gopkg.in/mgo.v2"
	// "time"
	"os"
)

var (
	START_TIME  = time.Date(2015, 3, 9, 0, 0, 0, 0, time.Local)
	SOURCE_FILE = "/data/realtime.csv"
	LOG_FILE    = "/data/realtime.log"
)

type Ball struct {
	ID   int
	Time time.Time
	Cell []int
}

func NewBallByBytes(bs []byte) (*Ball, error) {
	ss := strings.Split(string(bs), ",")
	if len(ss) == 7 {
		id, err := strconv.Atoi(ss[0])
		if err != nil {
			return nil, err
		}

		t, err := time.Parse("2006-01-02 15:04:05", ss[1])
		if err != nil {
			return nil, err
		}
		cell := make([]int, 5)
		for i, v := range ss[2:] {
			cell[i], err = strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
		}
		return &Ball{
			ID:   id,
			Time: t,
			Cell: cell,
		}, nil
	}
	return nil, fmt.Errorf("Error Ball Bytes ")
}

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
	f, err := os.OpenFile(SOURCE_FILE, os.O_APPEND|os.O_RDWR, 0644)
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

func GetLastLine() (*Ball, error) {
	f, err := os.Open(SOURCE_FILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 41)
	stat, err := os.Stat(SOURCE_FILE)
	start := stat.Size() - 42
	_, err = f.ReadAt(buf, start)
	if err != nil {
		return nil, err
	}
	return NewBallByBytes(buf)
}
