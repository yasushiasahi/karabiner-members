package data

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// Member is Scraped from Web site
type Member struct {
	Img   string
	Title string
	Name  string
}

var url = "https://www.karabiner.tech"

// Scrape get memberInfo from Web site
func Scrape() (ms []Member, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		m := "status code error: " + strconv.FormatInt(int64(res.StatusCode), 10) + res.Status
		err = dataError{m}
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}

	doc.Find(".memberListItem").Each(func(i int, s *goquery.Selection) {
		m := Member{Img: url}
		src, _ := s.Find(".memberImg").Find("img").Attr("src")
		m.Img += src
		m.Title = s.Find(".nameTitle").Text()
		m.Name = s.Find(".memberName").Text()
		ms = append(ms, m)
	})
	return
}
