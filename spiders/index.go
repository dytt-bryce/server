package spiders

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type DyttMovie struct {
	Name       string
	URL        string
	Cover      string
	Score      string
	Resolution string
}

type MovieData struct {
	Code    int
	Message string
	Data    interface{}
}

type Spider struct {
	Name string
}

func (spider Spider) GetPageData(url string) []DyttMovie {
	BaseURL := "http://m.dytt.com"
	timeout := time.Duration(time.Second * 10)
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println(spider.Name + "get page faild")
		return []DyttMovie{}
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(spider.Name + "get page faild")
		return []DyttMovie{}
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(spider.Name + "get page faild")
		return []DyttMovie{}
	}
	var allData []DyttMovie
	document.Find("div.list_mov").Each(func(i int, selection *goquery.Selection) {
		movieName := selection.Find("h4 a").Text()
		url, _ := selection.Find("h4 a").Attr("href")
		url = BaseURL + url
		cover, _ := selection.Find("img.img-responsive").Attr("src")
		cover = BaseURL + cover
		score := selection.Find("span.poster.score").Text()
		resolution := selection.Find("em").Text()
		dyttMovie := DyttMovie{
			Name:       movieName,
			URL:        url,
			Cover:      cover,
			Score:      score,
			Resolution: resolution,
		}
		allData = append(allData, dyttMovie)
	})
	return allData
}

type MovieDetail struct {
	Name string
	// Cover       string
	// Resolution  string
	// Score       string
	// Language    string
	// Type        string
	// Year        string
	// Location    string
	// Duration    string
	// Time        string
	// DoubanScore string
	// Actor       string
	// Director    string
	// Description string
	// PlayURL     string
}

func (spider Spider) GetMovieDetail(url string) MovieDetail {
	fmt.Println("get " + url)
	var detail MovieDetail
	timeout := time.Duration(time.Second * 10)
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println(spider.Name + "get page faild")
		return detail
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(spider.Name + "get page faild")
		return detail
	}
	defer res.Body.Close()
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(spider.Name + "get page faild")
		return detail
	}
	mainbody := document.Find("div#mainbody")
	desc := mainbody.Find("p.col-sm-12").Text()
	detail.Name = strings.TrimSpace(desc)
	return detail
}
