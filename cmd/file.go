/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "批量下载 m3u8 文件",
	Long:  `m3d file filename.txt -c 30 -o download/path`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) == 0 {
			fmt.Println("请输入文件地址")
			return
		}
		fileName := args[0]
		if !strings.HasPrefix(fileName, "/") {
			basePath, _ := os.Getwd()
			fileName = filepath.Join(basePath, fileName)
		}

		info, err := os.Stat(fileName)
		if err != nil {
			log.Printf("文件不存在：%s", err.Error())
			return
		}
		if info.Size() == 0 {
			log.Println("文件为空")
			return
		}
		outDir, err = getOutDirPath(info.Name()[:len(info.Name())-len(filepath.Ext(info.Name()))] + "_out")
		if err != nil {
			log.Println(err.Error())
			return
		}
		if _, err := os.Stat(outDir); err != nil {
			if err2 := os.MkdirAll(outDir, os.ModePerm); err2 != nil {
				log.Printf("文件创建失败：%s", err2.Error())
				return
			}
		}

		file, err := os.Open(fileName)
		if err != nil {
			log.Printf("文件打开失败：%s", err.Error())
			return
		}
		defer file.Close()

		var start int64 // 开始
		tmpFile, err := os.OpenFile(fileName+".tmp", os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Printf("临时文件创建失败：%s", err.Error())
			return
		}
		defer tmpFile.Close()
		numStr, err := io.ReadAll(tmpFile)
		if err != nil {
			start = 0
		}
		if len(numStr) > 0 {
			num, err := strconv.Atoi(string(numStr))
			if err == nil {
				start = int64(num)
			}
		}
		if start > 0 {
			log.Printf("下载开始位置第 %d 行", start)
		}

		// 读取每一行
		var line string
		var total int64
		scanner := bufio.NewScanner(file)
		buf := make([]byte, 0, 64*1024) // 初始缓冲区大小为 64KB
		scanner.Buffer(buf, 1024*1024)  // 最大缓冲区设为 1MB
		for scanner.Scan() {
			total++
			if total <= start {
				continue
			}
			line = scanner.Text()
			line = strings.TrimSpace(line)
			items := strings.Split(line, ".m3u8")
			if len(items) == 2 {
				downloadName := strings.TrimSpace(items[1])
				downloadUrl := strings.TrimSpace(items[0] + ".m3u8")
				if err := startDownload(outDir, downloadUrl, downloadName, number); err != nil {
					log.Printf("下载任务创建失败:%s", err.Error())
				}
			}

			if _, err := tmpFile.Seek(0, 0); err != nil {
				log.Printf("临时文件写入失败:%s", err.Error())
			}
			if _, err := tmpFile.WriteString(fmt.Sprintf("%d", total)); err != nil {
				log.Printf("下载记录失败:%s", err.Error())
			}
		}
		if err := scanner.Err(); err != nil {
			log.Printf("下载中断:%s", err.Error())
			return
		}
		log.Printf("【%s】下载完成", fileName)
		if err := os.Remove(tmpFile.Name()); err != nil {
			log.Printf("%s 临时文件删除失败：%s", tmpFile.Name(), err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
	fileCmd.Flags().IntVarP(&number, "number", "c", 30, "线程数")
	fileCmd.Flags().StringVarP(&outDir, "outDir", "o", "out", "保存路径")
}
