package provider

import (
	"encoding/json"
	"fmt"
	"github.com/orangbus/m3d/app/response/resp"
	"github.com/orangbus/m3d/pkg/request"
	"github.com/spf13/cast"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	AcList      = "list"
	AcVideoList = "videolist"
)

/*
*
spid
*/
type MovieProvider struct {
	apiUrl string `json:"api_url"`

	keywords string `json:"keywords"`
	pg       int    `json:"pg"`
	size     int    `json:"size"`
	h        int    `json:"h"`
	t        int    `json:"t"`
	ac       string `json:"ac"` // list | videolist
}

func NewMovie(baseUrl string) *MovieProvider {
	return &MovieProvider{
		apiUrl: baseUrl,
		pg:     1,
		size:   20,
		ac:     AcList,
	}
}

func (this MovieProvider) SetPage(page int) MovieProvider {
	if page > 0 {
		this.pg = page
	}
	return this
}
func (this MovieProvider) SetLimit(limit int) MovieProvider {
	if limit > 0 && limit < 20 {
		this.size = limit
	}
	return this
}
func (this MovieProvider) SetType(t int) MovieProvider {
	if t > 0 {
		this.t = t
	}
	return this
}
func (this MovieProvider) SetHour(h int) MovieProvider {
	if h > 0 {
		this.h = h
	}
	return this
}
func (this MovieProvider) SetAcVideoList() MovieProvider {
	this.ac = AcVideoList
	return this
}
func (this MovieProvider) SetKeyWords(keywords string) MovieProvider {
	if keywords != "" {
		this.keywords = keywords
	}
	return this
}

func (this MovieProvider) GetList() (resp.RespMovieVideoList, error) {
	result := resp.RespMovieVideoList{}
	bytes, err := get(this)
	log.Println(string(bytes))
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (this MovieProvider) Ping() bool {
	_, err := http.Get(this.apiUrl)
	if err != nil {
		return false
	}
	return true
}
func (this MovieProvider) Decode() string {
	return transformParam(this)
}

// 转化请求参数
func transformParam(param MovieProvider) string {
	data := url.Values{}
	if param.ac == "" {
		param.ac = "list"
	}
	data.Set("ac", param.ac)

	if param.pg > 0 {
		data.Set("pg", cast.ToString(param.pg))
	}
	if param.size > 0 {
		data.Set("size", strconv.Itoa(param.size))
	}
	if param.h > 0 {
		data.Set("h", strconv.Itoa(param.h))
	}
	if param.t > 0 {
		data.Set("t", strconv.Itoa(param.t))
	}
	if param.keywords != "" {
		data.Set("keywords", param.keywords)
	}
	return fmt.Sprintf("%s?%s", param.apiUrl, data.Encode())
}

// 发送请求
func get(movie MovieProvider) ([]byte, error) {
	api_url := transformParam(movie)
	response, err := request.Request(api_url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}
