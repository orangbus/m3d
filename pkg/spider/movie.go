package spider

import (
	"encoding/json"
	"fmt"
	"github.com/orangbus/m3d/app/response/resp"
	"github.com/orangbus/m3d/pkg/config"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Spider struct {
	Url          string `json:"url"`          // 接口地址
	total        int64  `json:"total"`        // 总数
	hour         int    `json:"hour"`         // 小时
	typeId       int    `json:"type_id"`      // 分类
	ac           string `json:"ac"`           // 数据类型
	pg           int    `json:"pg"`           // 分页
	size         int    `json:"size"`         // 大小
	wd           string `json:"wd"`           // 搜索关键词
	proxy        string `json:"proxy"`        // 代理
	proxy_status int    `json:"proxy_status"` // 代理
	ids          string `json:"ids"`
}

func NewSpider(url string, proxy ...string) *Spider {
	proxy_url := config.GetString("proxy.url")
	if len(proxy) > 0 {
		proxy_url = proxy[0]
	}
	return &Spider{
		Url:   url,
		pg:    1,
		size:  20,
		ac:    "videolist",
		proxy: proxy_url,
	}
}

func (s *Spider) GetMovieCate() (resp.RespMovieVideoList, error) {
	s.ac = "list"
	return s.Get()
}

func (s *Spider) SetHour(hour int) *Spider {
	s.hour = hour
	return s
}

func (s *Spider) SetTypeId(type_id int) *Spider {
	s.typeId = type_id
	return s
}

func (s *Spider) SetKeyword(keyword string) *Spider {
	s.wd = keyword
	return s
}
func (s *Spider) SetProxy(proxy string) *Spider {
	s.proxy = proxy
	return s
}
func (s *Spider) SetProxyStatus(status int) *Spider {
	s.proxy_status = status
	return s
}

func (s *Spider) SetPg(pg int) *Spider {
	s.pg = pg
	return s
}

func (s *Spider) SetAc(ac string) *Spider {
	s.ac = ac
	return s
}

// 获取请求地址
func (s *Spider) GetRequestUrl() string {
	param := url.Values{}
	if s.hour > 0 {
		param.Set("h", strconv.Itoa(s.hour))
	}
	if s.typeId > 0 {
		param.Set("t", strconv.Itoa(s.typeId))
	}
	if s.wd != "" {
		param.Set("wd", s.wd)
	}

	if s.ac != "" {
		param.Set("ac", s.ac)
	} else {
		param.Set("ac", "videolist")
	}

	if s.pg > 0 {
		param.Set("pg", strconv.Itoa(s.pg))
		if s.size > 0 {
			if s.size > 20 {
				s.size = 20
			}
			param.Set("size", strconv.Itoa(s.size))
		}
	}
	req_url := fmt.Sprintf("%s?%s", s.Url, param.Encode())
	if s.proxy_status == 1 {
		req_url = fmt.Sprintf("%s%s", s.proxy, req_url)
	}
	if config.GetBool("app.debug") {
		log.Printf("请求地址：%s", req_url)
	}
	return req_url
}

func (s *Spider) GetDetail(ids string) (resp.RespMovieVideoList, error) {
	s.ids = ids
	param := url.Values{}
	param.Set("ac", "videolist")
	param.Set("ids", ids)
	req_url := fmt.Sprintf("%s?%s", s.Url, param.Encode())
	if s.proxy_status == 1 {
		req_url = fmt.Sprintf("%s%s", s.proxy, req_url)
	}

	var result resp.RespMovieVideoList
	response, err := http.Get(req_url)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return result, err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytes, &result)
	return result, nil
}

func (s *Spider) Get() (resp.RespMovieVideoList, error) {
	var result resp.RespMovieVideoList
	response, err := http.Get(s.GetRequestUrl())
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return result, err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytes, &result)
	return result, nil
}
