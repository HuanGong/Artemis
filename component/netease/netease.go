package netease

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	//新闻客户端主页api
	NewsUrl string = "http://c.m.163.com"
)

func PaseNewsCategory() (map[string]string, error) {
	res, err := http.Get("http://c.m.163.com")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		logrus.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, errors.New("http status not 200")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		logrus.Errorln("parse response body error; ", err.Error())
		return nil, err
	}

	category := map[string]string{}
	doc.Find("body ul li a").Each(func(i int, s *goquery.Selection) {
		href, hit := s.Attr("href")
		if hit {
			category[s.Text()] = href
		}
	})
	return category, nil
}

func ParseCategoryContent(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		logrus.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, errors.New("http status not 200")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)

	contentUrls := []string{}
	doc.Find("urlset url").Each(func(i int, s *goquery.Selection) {
		url := s.Text()
		logrus.Debugln("url:", url)
		if len(url) > 0 {
			contentUrls = append(contentUrls, url)
		}
	})
	return contentUrls, nil
}
