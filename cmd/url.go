/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cast"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	number int
	outDir string
)

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "下载一个m3u8链接",
	Long:  `使用方式：m3d url http://xxx.meu8 -n xxx.mp4 -c 30 -o /download/path | out/path`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("请输入m3u8链接地址")
			return
		}
		url := args[0]
		if !strings.Contains(url, "m3u8") {
			log.Println("请输入正确的m3u8链接地址")
			return
		}
		var err error
		outDir, err = getOutDirPath(outDir)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("保存目录:", outDir)
		if err := startDownload(outDir, url, name, number); err != nil {
			log.Println(err)
		}
		log.Printf("下载成功:%s", outDir)
	},
}

func init() {
	rootCmd.AddCommand(urlCmd)
	urlCmd.Flags().StringVarP(&name, "name", "n", cast.ToString(time.Now().Unix()), "文件名，默认为当前时间戳")
	urlCmd.Flags().IntVarP(&number, "number", "c", 25, "线程数")
	urlCmd.Flags().StringVarP(&outDir, "outDir", "o", "out", "保存路径")
}
