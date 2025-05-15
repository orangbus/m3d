package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Github struct {
	api_url          string `json:"api_Url"`
	downloadDir      string `json:"downloadDir"`
	downloadFileName string `json:"downloadFileName"`
}

type GithubRelease struct {
	TagName         string `json:"tag_name"`         // 版本号 v0.0.1
	HtmlUrl         string `json:"html_url"`         // release 地址
	Body            string `json:"body"`             // 更新内容
	ZipballUrl      string `json:"zipball_url"`      // 下载地址
	TargetCommitish string `json:"target_commitish"` // 类似签名，解压的时候需要用到 156b117 6eca78e4e63cc96bb920e627ac72ababa
}

func NewGithub() *Github {
	return &Github{
		api_url:          "https://api.github.com/repos/orangbus/m3u8/releases",
		downloadDir:      "ui",
		downloadFileName: "ui.zip",
	}
}

func (c *Github) Get() (GithubRelease, error) {
	var list []GithubRelease
	resp, err := http.Get(c.api_url)
	if err != nil {
		return GithubRelease{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return GithubRelease{}, errors.New(fmt.Sprintf("网络请求错误，错误状态码:%d", resp.StatusCode))
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return GithubRelease{}, err
	}
	if err := json.Unmarshal(all, &list); err != nil {
		return GithubRelease{}, err
	}
	if len(list) == 0 {
		return GithubRelease{}, errors.New("未获取到最新版本")
	}
	return list[0], nil
}

// 检查更新
func (c *Github) Check() {

}

func (c *Github) Download() error {
	resp, err := c.Get()
	if err != nil {
		return err
	}
	response, err := http.Get(resp.ZipballUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return errors.New(fmt.Sprintf("网络请求错误，错误状态码:%d", response.StatusCode))
	}
	baseDir, err := os.Getwd()
	if err != nil {
		return err
	}
	downloadPath := filepath.Join(baseDir, c.downloadDir, c.downloadFileName)
	outFile, err := os.Create(downloadPath)
	if err != nil {
		return err
	}
	if _, err := io.Copy(outFile, response.Body); err != nil {
		return err
	}
	outFile.Close()
	log.Printf("下载成功：%s", outFile.Name())
	if err := exec.Command("unzip", "-o", outFile.Name(), "-d", filepath.Join(baseDir, c.downloadDir)).Run(); err != nil {
		return err
	}
	srcName := "orangbus-m3u8-" + resp.TargetCommitish[:7]
	srcDir := filepath.Join(baseDir, c.downloadDir, srcName, "/")
	log.Println(srcDir)
	log.Println(filepath.Join(baseDir, c.downloadDir))
	if err := exec.Command("mv", srcDir, filepath.Join(baseDir, c.downloadDir)).Run(); err != nil {
		return err
	}
	if err := os.RemoveAll(srcDir); err != nil {
		return err
	}
	if err := os.Remove(outFile.Name()); err != nil {
		return err
	}
	return nil
}
