package spider_serve

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/orangbus/m3d/app/models"
	"github.com/orangbus/m3d/pkg/debug"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func Start() {
	siteList := map[string]string{
		"小故事": "https://www.xigushi.com",
	}

	for key, site := range siteList {
		getNav(key, site)
	}
}

func getNav(siteName, site_url string) {
	response, _ := httpGet(site_url)
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		debug.Error("goquery解析错误", err)
		return
	}
	selection := doc.Selection.Find(".conter ul li")
	selection.Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		href, ok := a.Attr("href")
		if ok && href != "/" {
			getArticleList(s.Text(), fmt.Sprintf("%s%s", site_url, href), siteName, site_url)
		}
	})
}

func getArticleList(cateName, cateUrl, siteName, siteUrl string) {
	response, _ := httpGet(cateUrl)
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		debug.Error("goquery解析错误", err)
		return
	}

	selection := doc.Selection.Find(".list dl dd ul li a")
	selection.Each(func(i int, s *goquery.Selection) {
		article := models.Articles{}
		article.SiteName = siteName
		article.CateName = cateName
		article.Title = s.Text()
		article.Url = fmt.Sprintf("%s%s", siteUrl, s.AttrOr("href", ""))
		getArticleDetail(article)
	})
	nextUrl := getNextPage(doc)
	if nextUrl != "" {
		getArticleList(cateName, fmt.Sprintf("%s%s", siteUrl, nextUrl), siteName, siteUrl)
	}
}

func getArticleDetail(article models.Articles) {
	if article.Url == "" {
		return
	}
	response, _ := httpGet(article.Url)
	defer response.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(response.Body)
	content := doc.Find(".by dl dd p")
	log.Println(content.Text())
}

func httpGet(url string) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr, Timeout: 30 * time.Second}
	response, err := client.Get(url)
	if err != nil {
		debug.Log("请求错误", err)
		return nil, err
	}
	return response, nil
}

func getNextPage(document *goquery.Document) string {
	se := document.Selection.Find(".pages ul .sy2")
	if se == nil {
		return ""
	}
	return se.AttrOr("href", "")
}

func gerFullUrl(baseUrl, path string, dot ...string) (string, error) {
	sep := "/"
	if len(dot) > 0 {
		sep = dot[0]
	}
	parse, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	// 获取后缀的
	index := strings.LastIndex(parse.Path, sep)
	prefurl := baseUrl[:index]
	return prefurl, nil
}
