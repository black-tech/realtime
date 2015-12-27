package collect

import (
	"fmt"
	"testing"
)

func TestDownloadAll(t *testing.T) {
	return
	pullAll()
}

func TestGetLatestLine(t *testing.T) {
	ball, err := GetLastLine()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ball)
	fmt.Println(ball.ID % 1000)
}

func TestConnMysql(t *testing.T) {
	qr := DB().Query("select 1+1")
	if qr.Error != nil {
		t.Error(qr.Error)
	}
	for _, row := range qr.Rows {
		for k, v := range qr.Cols {
			fmt.Println(v, ": ", row.Int(k))
			if row.Int(0) != 2 {
				t.Error("none")
			}
		}
	}
}
