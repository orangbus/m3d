package movie_utils

import (
	"github.com/orangbus/m3d/app/models"
	"strings"
)

type MovieUrlItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// vod_play_from  vod_play_note vod_play_url
func ParseMovieUrl(movie models.Movies) []MovieUrlItem {
	// 是否有分隔符
	var cate []string       // 分类列表
	var cateList []string   // 分类视频列表
	var list []MovieUrlItem // 返回视频链接地址
	if movie.VodPlayNote != "" {
		cate = strings.Split(movie.VodPlayFrom, movie.VodPlayNote)
		cateList = strings.Split(movie.VodPlayURL, movie.VodPlayNote)

		for cateIndex, cateName := range cate {
			// 判断名称里面是否包含 M3u8
			if strings.Contains(cateName, "m3u8") {
				urlListmap := strings.Split(cateList[cateIndex], "#")
				for _, urlStr := range urlListmap {
					urlList := strings.Split(urlStr, "$")
					switch len(urlList) {
					case 2:
						list = append(list, MovieUrlItem{Name: urlList[0], Url: urlList[1]})
					case 1:
						list = append(list, MovieUrlItem{Name: movie.VodName, Url: urlList[0]})
					}
				}
			}
		}
		return list
	}

	// 是否存在 # 分割
	if strings.Contains(movie.VodPlayURL, "#") {
		cateList = strings.Split(movie.VodPlayURL, "#")
		for _, urlStr := range cateList {
			urlList := strings.Split(urlStr, "$")
			switch len(urlList) {
			case 2:
				list = append(list, MovieUrlItem{Name: urlList[0], Url: urlList[1]})
			case 1:
				list = append(list, MovieUrlItem{Name: movie.VodName, Url: urlList[0]})
			}
		}
		return list
	} else { // 只有单个地址的情况
		item := strings.Split(movie.VodPlayURL, "$")
		switch len(item) {
		case 2:
			list = append(list, MovieUrlItem{Name: item[0], Url: item[1]})
		case 1:
			list = append(list, MovieUrlItem{Name: movie.VodName, Url: item[0]})
		}
		return list
	}
}
