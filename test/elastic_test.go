package test

import (
	"encoding/json"
	"github.com/orangbus/m3d/app/models"
	"github.com/orangbus/m3d/app/service/elastic_service"
	"github.com/orangbus/m3d/bootstrap"
	"github.com/orangbus/m3d/pkg/assert"
	"github.com/orangbus/m3d/pkg/database"
	"log"
	"testing"
)

func init() {
	bootstrap.SetUp()
}

func TestCon(t *testing.T) {
	log.Println("666")
}

func TestMappingCreate(t *testing.T) {
	mapping := map[string]interface{}{
		"mappings": map[string]interface{}{
			"properties": map[string]interface{}{
				"id": map[string]interface{}{
					"type": "long",
				},
			},
		},
	}
	err := elastic_service.CreateMapping("demo", mapping)
	if assert.Error("index", err) {
		return
	}
}

func TestGetMapping(t *testing.T) {
	err := elastic_service.GetMapping("demo")
	if assert.Error("index", err) {
		return
	}
}

func TestFindOne(t *testing.T) {
	data, _ := elastic_service.FindOne("movies", "206410")
	log.Print(string(data))
}

func TestAnalyzer(t *testing.T) {
	list, _ := elastic_service.Analyzer("韩国极品推特网红美少女《Sulaa》男友肉棒抽插慢玩白虎粉嫩美穴极具带感")
	for _, item := range list {
		log.Println(item)
	}
	var video models.Video
	video.Tags, _ = json.Marshal(list)
	if err := database.DB.Create(&video).Error; err != nil {
		t.Log(err)
	}
}

func TestGetTags(t *testing.T) {
	var video models.Video
	database.DB.First(&video, 1)
	log.Println(string(video.Tags))
}

func TestLikeQuery(t *testing.T) {
	tags, _ := elastic_service.Analyzer("韩国极品推特网红美少女《Sulaa》男友肉棒抽插慢玩白虎粉嫩美穴极具带感")

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"more_like_this": map[string]interface{}{
				"fields": []string{
					"vod_name",
				},
				"like":            tags,
				"min_term_freq":   1,
				"max_query_terms": 12,
			},
		},
	}
	list, _, _ := elastic_service.Search("movies", query, 1, 20)
	for _, item := range list {
		log.Println(item)
	}
}
