package collect

import (
	"fmt"
	"testing"
	"time"
)

func TestDownload(t *testing.T) {
	ret := GetData(time.Now())
	fmt.Println(ret)
	ret = GetData(time.Now().AddDate(0, 0, -1))
	fmt.Println(ret)
}
