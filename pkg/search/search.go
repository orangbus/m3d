package search

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/orangbus/m3d/pkg/config"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	Es *elasticsearch.TypedClient
)

/*
*
初始化es客户端
*/
func NewSearch() {
	addrs := []string{}
	es_hosts := config.GetString("es.host")
	name := config.GetString("es.username")
	password := config.GetString("es.password")
	if strings.Contains(es_hosts, ",") {
		split := strings.Split(es_hosts, ",")
		for _, v := range split {
			addrs = append(addrs, v)
		}
	} else {
		addrs = append(addrs, es_hosts)
	}
	c := elasticsearch.Config{
		Addresses: addrs,
		Username:  name,
		Password:  password,
		Transport: &http.Transport{
			MaxIdleConns:        100,             // 最大空闲连接数
			MaxIdleConnsPerHost: 10,              // 每个主机的最大空闲连接数
			IdleConnTimeout:     time.Second * 3, // 空闲连接超时时间
		},
	}
	client, err := elasticsearch.NewTypedClient(c)
	if err != nil {
		panic("elasticsearch连接失败")
	}
	Es = client
	response, err := client.Info().Do(context.Background())
	if err != nil {
		log.Println(err.Error())
		panic(err.Error())
	}
	log.Printf("elasticsearch version:%s\n", response.Version.Int)
}
