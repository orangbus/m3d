package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/orangbus/m3d/pkg/m3u8/dl"
	"os"
	"path/filepath"
	"strings"
)

// 设置公共下载目录
func getOutDirPath(outDir string) (string, error) {
	if !strings.HasPrefix(outDir, "/") {
		d, err := os.Getwd()
		if err != nil {
			return outDir, errors.New(fmt.Sprintf("获取当前目录失败:%s", err.Error()))
		}
		outDir = filepath.Join(d, "download", outDir)
	} else {
		outDir = filepath.Join("download", outDir)
	}
	return outDir, nil
}

func startDownload(outDir, url, name string, number int) error {
	ctx := context.Background()
	downloader, err := dl.NewTask(outDir, url)
	if err != nil {
		return errors.New(fmt.Sprintf("下载任务创建失败:%s", err.Error()))
	}
	name = fmt.Sprintf("%s.mp4", name)
	if err := downloader.Start(number, name, ctx); err != nil {
		return errors.New(fmt.Sprintf("下载失败:%s", err.Error()))
	}
	return nil
}
