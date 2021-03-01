package helper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
)

var (
	adUrl      string
	isLastPage bool
	listUrl    string
)

func startParser() {
	for _, cat := range urlCats {
		ind := 1
		for true {
			listUrl = fmt.Sprintf("%sc%d-p%d.html", baseUrl, cat, ind)
			geziyor.NewGeziyor(&geziyor.Options{
				StartURLs: []string{listUrl},
				ParseFunc: parseListPage,
			}).Start()
			if isLastPage {
				break
			} else {
				ind++
			}
		}
	}
}

func parseListPage(_ *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.topmess").Add("div.stradv").Each(func(i int, s *goquery.Selection) {
		adUrl = fmt.Sprintf("%s%s", baseUrl, s.Find("b").Find("a").AttrOr("href", ""))
		geziyor.NewGeziyor(&geziyor.Options{
			StartURLs: []string{adUrl},
			ParseFunc: parseAdPage,
		}).Start()
	})
	isLastPage = r.HTMLDoc.Find("i.fa-chevron-right").Length() == 0
}

func parseAdPage(_ *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.well").Each(func(i int, s *goquery.Selection) {
		protectedEmail, exists := s.Find("sup").Find("em").Find("a").Attr("data-cfemail")
		if exists {
			setCellValue(encode(protectedEmail))
		}
	})
}
