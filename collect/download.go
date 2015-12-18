package collect

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"net/http"
)

type SpiderData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Rows int `json:"rows"`
		Data []struct {
			Phasetype string `json:"phasetype"`
			Phase     string `json:"phase"`
			TimeDraw  string `json:"time_draw"`
			Result    struct {
				Result []struct {
					Key  string   `json:"key"`
					Data []string `json:"data"`
				} `json:"result"`
			} `json:"result"`
			Ext struct {
				Ten  string `json:"ten"`
				Unit string `json:"unit"`
				Last string `json:"last"`
			} `json:"ext"`
		} `json:"data"`
	} `json:"data"`
	Redirect  string `json:"rediret"`
	Timestamp int64  `json:"timestamp"`
}

func (s *SpiderData) SaveFile(f *os.File) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("err in save, ", r)
			err = r.(error)
		}
	}()
	if s == nil {

	}
	if s.Code == 0 {
		if set := s.Data.Data; len(set) > 0 {
			for i := len(set) - 1; i >= 0; i-- {
				cell := set[i]
				if balls := cell.Result.Result[0].Data; len(balls) == 5 {
					sData := []string{cell.Phase, cell.TimeDraw, strings.Join(balls, ",")}
					str := strings.Join(sData, ",")
					f.WriteString(str + "\n")
				}
			}
		}
	}
	return nil
}

func GetData(date time.Time) (*SpiderData, error) {
	req := NewRequest(date)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer func() {
		resp.Body.Close()
	}()
	ret := new(SpiderData)
	err = json.NewDecoder(resp.Body).Decode(ret)
	return ret, err
}

func NewRequest(t time.Time) *http.Request {
	URL := "http://baidu.lecai.com/lottery/draw/sorts/ajax_get_draw_data.php"
	date := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
	req, err := http.NewRequest("GET", URL+"?lottery_type=200&date="+date, nil)
	if err != nil {
		fmt.Println("new req, ", err)
	}

	req.Header.Add("Accept-Encoding", "deflate")
	req.Header.Add("Accept-Encoding", "sdch")
	req.Header.Add("Accept-Language", "zh-CN,zh")
	req.Header.Add("Accept-Language", "q=0.8,en")
	req.Header.Add("Accept-Language", "q=0.6,de")
	req.Header.Add("Accept-Language", "q=0.4,zh-TW")
	req.Header.Add("Accept-Language", "q=0.2")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept", "text/javascript")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept", "q=0.01")
	req.Header.Add("Referer", "http://baidu.lecai.com/lottery/draw/view/200?agentId=5555")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	strings.Count("sd", "sep")
	time.Now()
	headCookie := []string{
		"Hm_lpvt_6c5523f20c6865769d31a32a219a6766=1449293219",
		"Hm_lpvt_9b75c2b57524b5988823a3dd66ccc8ca=1449293219",
		"Hm_lvt_6c5523f20c6865769d31a32a219a6766=1448809736,1448809796,1449126226,1449293219",
		"Hm_lvt_9b75c2b57524b5988823a3dd66ccc8ca=1448809736,1448809796,1449126226,1449293219",
		"Hm_lvt_ddaa40fe0ef9967e65e6956736d327af=1448809651",
		"_lhc_uuid=sp_565b0507160987.09890935",
		"_source=5555",
		"_source_pid=0",
		"_srcsig=b109a5da",
		"lehecai_request_control_stats=2",
	}

	for _, v := range headCookie {
		kies := strings.Split(v, "=")
		c := &http.Cookie{
			Name:    kies[0],
			Value:   kies[1],
			Expires: time.Now(),
			Path:    "/",
			Domain:  "baidu.lecai.com",
		}
		req.AddCookie(c)
	}
	return req
}
