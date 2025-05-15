package request

import (
	"io"
	"testing"
)

func TestReq(t *testing.T) {
	//url := "https://google.com"
	url := "https://www.2222.la/api.php/provide/vod"
	response, err := Request(url)
	if err != nil {
		return
	}
	defer response.Body.Close()
	d, err := io.ReadAll(response.Body)
	if err != nil {
		t.Log(err.Error())
	}
	t.Logf(string(d))
}
