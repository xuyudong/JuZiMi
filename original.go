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
func runJuziMi() {

	m := map[string]string{"本周热门原创": "week", "最新原创句子": "ju", "推荐原创句子": "recommend"}
	//fileName := "D:/phpStudy/PHPTutorial/WWW/test.md"
	fileName := "C:/Users/Administrator/Desktop/test.md"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	st := ""
	for k, v := range m {
		//默认取10页
		for i := 0; i < 10; i++ {
			page := "?page=" + strconv.Itoa(i)
			url := "https://www.juzimi.com/original/" + v + page
			st = getJuzi(url)
			dstFile.WriteString(k + "\n")
			dstFile.WriteString(st + "\n")
		}
	}
	defer dstFile.Close()
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
