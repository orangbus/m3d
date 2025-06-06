package download_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/orangbus/m3d/app/models"
	"github.com/orangbus/m3d/pkg/database"
	"github.com/orangbus/m3d/pkg/m3u8/dl"
	"github.com/orangbus/m3d/pkg/spider"
	"github.com/orangbus/m3d/pkg/utils/movie_utils"
	"github.com/spf13/cast"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	mu         sync.Mutex
	status     bool // 是否在下载
	ctx        context.Context
	cancelFunc context.CancelFunc
)

func Start() {
	mu.Lock()
	if status {
		log.Println("已经在下载了")
		mu.Unlock()
		return
	}
	ctx, cancelFunc = context.WithCancel(context.Background())
	status = true
	mu.Unlock()
	log.Println("开始下载")
	if err := download(ctx); err != nil {
		log.Println(err.Error())
	}
	log.Println("下载结束")
}

func Stop() {
	mu.Lock()
	defer mu.Unlock()
	if cancelFunc != nil {
		cancelFunc()
	}
	status = false
}

func download(ctx context.Context) error {
	d, err := getDownloadUrl()
	if err != nil {
		return err
	}
	if d.ID == 0 {
		return nil
	}
	baseDir, err := os.Getwd()
	if err != nil {
		return err
	}
	name := fmt.Sprintf("%s.mp4", d.Name)
	outDir := filepath.Join(baseDir, "download/out")
	log.Println("下载位置：", outDir)
	// 判断文件是否存在
	filePath := filepath.Join(outDir, name)
	log.Printf("文件位置：%s", filePath)
	if _, err := os.Stat(filePath); err == nil {
		log.Printf("文件已存在，%s", filePath)
		if err2 := database.DB.Model(&models.Download{}).Where("id =?", d.ID).Updates(map[string]any{
			"status": 1,
			"remark": "文件已经存在",
		}).Error; err2 != nil {
			return err2
		}
		return download(ctx)
	}

	select {
	case <-ctx.Done():
		mu.Lock()
		status = false
		mu.Unlock()
		log.Println("手动下载停止", status)
		time.Sleep(time.Second)
	default:
		downloadStatus, err := downloadTask(outDir, d, 30)
		if err != nil {
			log.Printf(err.Error())
		}
		if err2 := database.DB.Model(&models.Download{}).Where("id =?", d.ID).Update("status", downloadStatus).Error; err2 != nil {
			return err2
		}
	}
	return download(ctx)
}

func getDownloadUrl() (models.Download, error) {
	var d models.Download
	if err := database.DB.Model(&models.Download{}).Where("status = ?", 0).First(&d).Error; err != nil {
		return d, err
	}
	if d.ID == 0 {
		return d, errors.New("暂无下载任务")
	}
	return d, nil
}

func downloadTask(outDir string, d models.Download, number int) (int, error) {
	name := fmt.Sprintf("%s.mp4", d.Name)
	downloader, err := dl.NewTask(outDir, d.Url)
	if err != nil {
		log.Printf("下载任务创建失败:%s", err.Error())
		if err2 := database.DB.Model(&models.Download{}).Where("id =?", d.ID).Update("status", 2).Error; err2 != nil {
			return 2, err
		}
		return 2, err
	}
	if err := downloader.Start(number, name, ctx); err != nil {
		log.Printf("下载失败:%s", err.Error())
		if err2 := database.DB.Model(&models.Download{}).Where("id =?", d.ID).Update("status", 2).Error; err2 != nil {
			return 2, err
		}
		return 2, err
	}
	if err := database.DB.Model(&models.Download{}).Where("id =?", d.ID).Update("status", 1).Error; err != nil {
		return 2, err
	}
	return 1, nil
}

func DownloadFavorite(favorite models.Favorite, day int) {
	res, err := spider.NewSpider(favorite.ApiUrl).SetHour(day * 24).Get()
	if err != nil {
		log.Printf("采集错误:%s", err.Error())
		return
	}
	if res.Total <= 1 {
		// 解析下载地址：
		saveDownload(res.List)
		return
	}

	// 采集更多
	for i := 2; i < cast.ToInt(res.Total); i++ {
		res, err := spider.NewSpider(favorite.ApiUrl).SetPg(i).SetHour(day * 24).Get()
		if err != nil {
			log.Printf("采集错误:%s", err.Error())
			continue
		}
		if res.Total <= 1 {
			// 解析下载地址：
			saveDownload(res.List)
		}
	}
}

func saveDownload(movieList []models.Movies) {
	var dwonloadList []models.Download
	for _, v := range movieList {
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
}
