package collect

import (
	"fmt"
	"testing"
)

func TestDownloadAll(t *testing.T) {
	return
	downloadAll()
}

func TestGetLatestLine(t *testing.T) {
	ball, err := GetLastLine()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ball)
	fmt.Println(ball.ID % 1000)
}
