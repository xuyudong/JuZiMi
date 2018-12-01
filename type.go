package main

import (
	"crypto/tls"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func main() {
	runJuziMi()
}

const (
	//TYPE = "new"//最新发布
	//TYPE = "recommend"//推荐
	TYPE = "todayhot"//今日热门
	//TYPE = "totallike"//最受欢迎
	
	DIR = "C:/Users/Administrator/Desktop/juzimi"
)

func runJuziMi() {
	createDir(DIR)
	fileName := DIR+"/"+TYPE+".md"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	st := ""
		//默认取10页
	dstFile.WriteString(TYPE + "\n")
	for i := 0; i < 10; i++ {
		page := "?page=" + strconv.Itoa(i)
		url := "https://www.juzimi.com/"+TYPE+"/"  + page
		st = getJuzi(url)
		dstFile.WriteString(st + "\n")
	}
	defer dstFile.Close()
}

func createDir(dir string)  {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Print("Create Directory OK!")
	}
}

func getJuzi(urls string) string {
	//http://www.89ip.cn/index.html
	//ip代理库
	proxy, _ := url.Parse("103.14.198.144:53281")

	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 5,
	}

	resp, err := client.Get(urls)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("err1")
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic("err2")
	}
	str := "\n"
	doc.Find(".views-field-phpcode").Find(".views-field-phpcode-1").Each(func(i int, s *goquery.Selection) {
		str += s.Text() + "\n\n"
	})
	return str
}
