package scrapper

import (
	"log"
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func Scrapper() {
	log := log.New(log.Writer(), "Log:", log.Ltime)
	log.Printf("Starting...")

	res, err := http.Get("https://sullygnome.com/channel/xqc/30")

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		log.Fatal("Error while reading HTML:", err)
	}

	doc.Find(".InfoStatPanelInner").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".InfoStatPanelTL").AttrOr("title", "No title found")

		fmt.Println("Title:", title)
	})

}
