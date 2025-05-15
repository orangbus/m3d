package elastic_service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/orangbus/m3d/pkg/search"
)

var ctx = context.Background()

func GetMapping(indexName string) error {
	_, err := search.Es.Indices.GetMapping().Index(indexName).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func CreateMapping(indexName string, mapping map[string]interface{}) error {
	marshal, err := json.Marshal(mapping)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(marshal)

	response, err := search.Es.Indices.Create(indexName).Raw(reader).Do(ctx)
	if err != nil {
		return err
	}
	if !response.Acknowledged {
		return errors.New(fmt.Sprintf("%s 索引创建失败", indexName))
	}
	return nil
}

func CreateIndex(indexName string) error {
	response, err := search.Es.Indices.Create(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if !response.Acknowledged {
		return errors.New(fmt.Sprintf("[%s]索引创建失败", indexName))
	}
	return nil
}

func DeleteIndex(indexName string) error {
	response, err := search.Es.Indices.Delete(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if !response.Acknowledged {
		return errors.New(fmt.Sprintf("[%s]索引删除失败", indexName))
	}
	return nil
}

func Create(indexName string, id string, doc interface{}) error {
	_, err := search.Es.Index(indexName).Id(id).Document(doc).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
func FindOne(indexName string, id string) ([]byte, error) {
	response, err := search.Es.Get(indexName, id).Do(ctx)
	if err != nil {
		return nil, err
	}
	if !response.Found {
		return nil, errors.New("记录不存在")
	}
	marshalJSON, err := response.Source_.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return marshalJSON, nil
}
func Update(indexName string, id string, doc interface{}) error {
	_, err := search.Es.Index(indexName).Id(id).Document(doc).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}
func Delete(indexName string, id string) error {
	_, err := search.Es.Delete(indexName, id).Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func Search(indexName string, query map[string]interface{}, page, limit int) ([]interface{}, int64, error) {
	var list []interface{}
	var total int64

	if page <= 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 20
	}
	query["from"] = (page - 1) * limit
	query["size"] = limit

	marshal, err := json.Marshal(query)
	if err != nil {
		return list, total, err
	}
	reader := bytes.NewReader(marshal)
	res, err2 := search.Es.Search().Index(indexName).Raw(reader).Do(ctx)
	if err2 != nil {
		return list, total, err
	}
	for _, item := range res.Hits.Hits {
		var data interface{}
		err := json.Unmarshal(item.Source_, &data)
		if err != nil {
			continue
		}
		list = append(list, data)
	}
	total = res.Hits.Total.Value
	return list, total, nil
}

// 分词
func Analyzer(text string, analyzer ...string) ([]string, error) {
	analyzerType := "ik_max_word" // ik_smart
	if len(analyzer) > 0 {
		analyzerType = analyzer[0]
	}
	var tokens []string
	response, err := search.Es.Indices.Analyze().Index("movies").Analyzer(analyzerType).Text(text).Do(ctx)
	if err != nil {
		return tokens, err
	}
	for _, item := range response.Tokens {
		if len(item.Token) > 3 { // 3:表示一个汉字
			tokens = append(tokens, item.Token)
		}
	}
	return tokens, nil
}

func DeleteByApiId(apiId int64) error {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"api_id": apiId,
			},
		},
	}
	marshal, err := json.Marshal(query)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(marshal)
	_, err2 := search.Es.DeleteByQuery("movies").Raw(reader).Do(ctx)
	if err2 != nil {
		return err2
	}
	return nil
}
