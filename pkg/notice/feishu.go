package notice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	MsgType string                 `json:"msg_type"`
	Content map[string]interface{} `json:"content"`
	Card    map[string]interface{} `json:"card"` // 卡片消息
}

/*
*
https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot#f62e72d5
*/
type FeishuNotice struct {
	BaseUrl string  `json:"base_url"`
	Body    Message `json:"body"`
}

func NewFeiShuNotice(api_url string) *FeishuNotice {
	notice := FeishuNotice{BaseUrl: api_url}
	return &notice
}

func (m *FeishuNotice) SeedMsg(msg string) error {
	body := Message{
		MsgType: "text",
		Content: map[string]interface{}{
			"text": fmt.Sprintf("[新消息]%s", msg),
		},
	}
	return seed(m.BaseUrl, body)
}

func (m *FeishuNotice) SeedText(title string, content string) error {
	body := Message{
		MsgType: "post",
		Content: map[string]interface{}{
			"post": map[string]interface{}{
				"zh_cn": map[string]interface{}{
					"title": title,
					"content": []interface{}{
						map[string]interface{}{
							"tag":  "text",
							"text": content,
						},
					},
				},
			},
		},
	}
	return seed(m.BaseUrl, body)
}

func (m *FeishuNotice) SeedCard(title string, msg string) error {
	body := Message{
		MsgType: "interactive",
		Card: map[string]interface{}{
			"header": map[string]interface{}{
				"title": map[string]interface{}{
					"tag":     "plain_text",
					"content": title,
				},
			},
			"elements": []interface{}{
				map[string]interface{}{
					"tag": "div",
					"text": map[string]interface{}{
						"content": msg,
						"tag":     "lark_md",
					},
				},
			},
		},
	}
	return seed(m.BaseUrl, body)
}

func seed(api_url string, body Message) error {
	marshal, err := json.Marshal(body)
	if err != nil {
		return err
	}

	response, err := http.Post(api_url, "application/json", bytes.NewReader(marshal))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	_, err = io.ReadAll(response.Body) // {"StatusCode":0,"StatusMessage":"success","code":0,"data":{},"msg":"success"}
	if err != nil {
		return err
	}
	fmt.Println("消息发送成功！")
	return nil
}
