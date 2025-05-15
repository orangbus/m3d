package request

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

func Request(api_url string) (*http.Response, error) {
	// 设置代理地址
	//proxyURL, err := url.Parse("http://127.0.0.1:7890")
	//if err != nil {
	//	fmt.Println("解析代理地址失败:", err)
	//	return nil, err
	//}
	// 创建一个自定义的http.Transport
	transport := &http.Transport{
		//Proxy: http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	// 使用自定义的http.Transport创建http.Client
	client := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 60,
	}

	// 发送请求
	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.1")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Connection", "keep-alive")

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("请求错误")
	}
	return response, nil
}
