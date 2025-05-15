/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/orangbus/m3d/app/provider"
	"github.com/orangbus/m3d/console"
	"github.com/orangbus/m3d/pkg/utils/movie_utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

type Iptv struct {
	Name       string `json:"name"`
	Logo       string `json:"logo"`
	GroupTitle string `json:"group-title"`
	Url        string `json:"url"`
}

var (
	api_url = "" // https://cj.yayazy.net/api.php/provide/vod
	name    = "movie"
	extName = "m3u8"
)

// iptvCmd represents the iptv command
var iptvCmd = &cobra.Command{
	Use:   "iptv",
	Short: "生成一个直播列表",
	Run: func(cmd *cobra.Command, args []string) {
		spider := provider.NewMovie(api_url).SetAcVideoList()
		movieVideo, err := spider.GetList()
		if err != nil {
			console.Error(err.Error())
			return
		}

		filename := fmt.Sprintf("%s.%s", name, extName)
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer file.Close()
		file.WriteString("#EXTM3U\n")

		ch := make(chan []string, 20)
		for i := 0; i <= movieVideo.Pagecount; i++ {
			go spiderVideo(i+1, ch)
			continue
		}
		//movieVideo.Pagecount = 1

		for i := 1; i <= movieVideo.Pagecount; i++ {
			res := <-ch
			if len(res) > 0 {
				data := ""
				for _, item := range res {
					data += fmt.Sprintf("%s\n", item)
				}
				file.WriteString(data)
			}
			p := fmt.Sprintf("采集进度：%d/%d ", i, movieVideo.Pagecount)
			log.Println(p)
		}
		console.Success("采集完毕")
	},
}

func init() {
	rootCmd.AddCommand(iptvCmd)

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	iptvCmd.Flags().StringVarP(&api_url, "api_url", "u", api_url, "url地址")
	iptvCmd.Flags().StringVarP(&name, "name", "n", name, "生成文件名称")
}

func spiderVideo(pg int, ch chan []string) {
	var list []string
	spider := provider.NewMovie(api_url).SetAcVideoList().SetPage(pg)
	movieVideo, err := spider.GetList()
	if err != nil {
		console.Error(err.Error())
		ch <- list
		return
	}
	for _, movie := range movieVideo.List {
		urlList := movie_utils.ParseMovieUrl(movie)
		for _, url := range urlList {
			play_name := fmt.Sprintf("%s_%s", movie.VodName, url.Name)
			item := fmt.Sprintf("#EXTINF:-1 logo='%s' group-title=%s name=%s,%s\n%s", movie.VodPic, movie.TypeName, movie.VodName, play_name, url.Url)
			list = append(list, item)
		}
	}
	ch <- list
}
