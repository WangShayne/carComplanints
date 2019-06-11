package main

import (
	"carComplanints/common"
	"carComplanints/db"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func init() {
	db.NewDB()
}

func main() {
	url := "http://www.12365auto.com/zlts/0-0-0-0-0-0_0-0-1.shtml"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("error code is :", resp.StatusCode)
		return
	}

	regComplanints(resp.Body)
}

func regComplanints(contents io.ReadCloser) {
	doc, err := goquery.NewDocumentFromReader(contents)
	if err != nil {
		log.Fatal(err)
	}
	var tmpClt *common.Complanint
	doc.Find(".com_sec .tslb_b table tbody").Each(func(i int, s *goquery.Selection) {
		trs := s.Find("tr")
		trs.Each(func(j int, s1 *goquery.Selection) {
			if j > 0 {
				td := s1.Find("td")

				bid, _ := td.Eq(2).Attr("bid")
				sid, _ := td.Eq(2).Attr("sid")
				mid, _ := td.Eq(2).Attr("mid")

				tmpUrl, _ := td.Eq(4).Find("a").Attr("href")

				tmpClt = &common.Complanint{
					SN:        td.Eq(0).Text(),
					Brand:     td.Eq(1).Text(),
					BrandID:   bid,
					Series:    td.Eq(2).Text(),
					SeriesID:  sid,
					Model:     td.Eq(3).Text(),
					ModelID:   mid,
					Sketch:    td.Eq(4).Text(),
					SketchURL: tmpUrl,
					Date:      td.Eq(6).Text(),
				}
				fmt.Println(tmpClt)
				db.InsertComplanint(tmpClt)
			}

		})
	})
}
