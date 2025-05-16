/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/orangbus/m3d/app/models"
	"github.com/orangbus/m3d/pkg/database"
	"github.com/orangbus/m3d/pkg/spider"
	"github.com/orangbus/m3d/pkg/utils/movie_utils"
	"log"

	"github.com/spf13/cobra"
)

// timerCmd represents the timer command
var timerCmd = &cobra.Command{
	Use:   "timer",
	Short: "定时下载 m3d timer",
	Long:  `定时下载收藏的频道`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := spider.NewSpider("https://cj.rycjapi.com/api.php/provide/vod").SetHour(24).Get()
		if err != nil {
			log.Print(err.Error())
			return
		}
		var dwonloadList []models.Download
		for _, v := range client.List {
			list := movie_utils.ParseMovieUrlItem(v.VodName, v.VodPlayFrom, v.VodPlayNote, v.VodPlayURL)
			for _, item := range list {
				fmt.Println(item.Name, item.Url)
				dwonloadList = append(dwonloadList, models.Download{
					ApiId:  1,
					Name:   item.Name,
					Url:    item.Url,
					Status: 0,
					Proxy:  0,
					Remark: "",
					OutDir: "",
				})
			}
		}

		// 保存数据库
		if len(dwonloadList) > 0 {
			if err := database.DB.Model(&models.Download{}).Create(&dwonloadList).Error; err != nil {
				log.Printf("写入数据失败：%s", err.Error())
			}
		}

		fmt.Println("更新成功")
	},
}

func init() {
	rootCmd.AddCommand(timerCmd)
	// timerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
