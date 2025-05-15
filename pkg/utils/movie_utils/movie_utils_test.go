package movie_utils

import (
	"fmt"
	"github.com/orangbus/m3d/app/models"
	"testing"
)

var (
	vod_play_from = "hhyun$$$hhm3u8"
	vod_play_note = "$$$"
	vod_play_url  = "第01集$https://play.hhuus.com/play/PdygjWwa#第02集$https://play.hhuus.com/play/RdGq0NQa#第3集$https://play.hhuus.com/play/YaOBKpGe#第4集$https://play.hhuus.com/play/Ddw43Lme#第05集$https://play.hhuus.com/play/1aKx4Erb#第06集$https://play.hhuus.com/play/9aAk4KBd#第07集$https://play.hhuus.com/play/nel4qZje#第08集$https://play.hhuus.com/play/Yer4zP2b#第09集$https://play.hhuus.com/play/Rb40AW1b#第10集$https://play.hhuus.com/play/BeXOyYWd# 第11集$https://play.hhuus.com/play/1aKxPyzb#第12集$https://play.hhuus.com/play/qaQEYGqa#第13集$https://play.hhuus.com/play/QeZQ32Qd# 第14集$https://play.hhuus.com/play/9b6Jg3Ob#第15集$https://play.hhuus.com/play/yb8Lm5lb#第16集$https://play.hhuus.com/play/rb2B65Jb#第17集$https://play.hhuus.com/play/pen4nnlb#第18集$https://play.hhuus.com/play/peny5Bpd#第19集$https://play.hhuus.com/play/PdyNl8we#第20集$https://play.hhuus.com/play/PdygQWga#第21集完结\t$https://play.hhuus.com/play/pen4nBRb$$$第01集$https://play.hhuus.com/play/PdygjWwa/index.m3u8#第02集$https://play.hhuus.com/play/RdGq0NQa/index.m3u8#第3集$https://play.hhuus.com/play/YaOBKpGe/index.m3u8#第4集$https://play.hhuus.com/play/Ddw43Lme/index.m3u8#第05集$https://play.hhuus.com/play/1aKx4Erb/index.m3u8#第06集$https://play.hhuus.com/play/9aAk4KBd/index.m3u8#第07集$https://play.hhuus.com/play/nel4qZje/index.m3u8#第08集$https://play.hhuus.com/play/Yer4zP2b/index.m3u8#第09集$https://play.hhuus.com/play/Rb40AW1b/index.m3u8#第10集$https://play.hhuus.com/play/BeXOyYWd/index.m3u8# 第11集$https://play.hhuus.com/play/1aKxPyzb/index.m3u8#第12集$https://play.hhuus.com/play/qaQEYGqa/index.m3u8#第13集$https://play.hhuus.com/play/QeZQ32Qd/index.m3u8# 第14集$https://play.hhuus.com/play/9b6Jg3Ob/index.m3u8#第15集$https://play.hhuus.com/play/yb8Lm5lb/index.m3u8#第16集$https://play.hhuus.com/play/rb2B65Jb/index.m3u8#第17集$https://play.hhuus.com/play/pen4nnlb/index.m3u8#第18集$https://play.hhuus.com/play/peny5Bpd/index.m3u8#第19集$https://play.hhuus.com/play/PdyNl8we/index.m3u8#第20集$https://play.hhuus.com/play/PdygQWga/index.m3u8#第21集完结\t$https://play.hhuus.com/play/pen4nBRb/index.m3u8"
	vod_play_url2 = "第1集$https://99recdn1.jjrbf.top/20240430/Z7ZvAiOS/index.m3u8#第2集$https://99recdn1.jjrbf.top/20240430/Qr5m5fDf/index.m3u8"
)

func TestParse(t *testing.T) {
	var movie models.Movies
	movie.VodPlayFrom = vod_play_from
	movie.VodPlayNote = vod_play_note
	movie.VodPlayURL = vod_play_url
	list := ParseMovieUrl(movie)
	for _, item := range list {
		fmt.Println(item)
	}
}

func TestParse2(t *testing.T) {
	var movie models.Movies
	movie.VodPlayFrom = ""
	movie.VodPlayNote = ""
	movie.VodPlayURL = vod_play_url2
	list := ParseMovieUrl(movie)
	for _, item := range list {
		fmt.Println(item)
	}
}
