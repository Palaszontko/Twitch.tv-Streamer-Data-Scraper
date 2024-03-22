package scrapper

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func scrapper(nickname string, timePeriod int) StreamerData {
	streamer := StreamerData{
		Nickname: nickname,
	}
	scrapperStreamerData(streamer.Nickname, timePeriod, &streamer)

	logger.Println("Streamer data scrapped.")
	return streamer
}

func Start() {
	var streamersDataArray []StreamerData

	sessionId := sessionId()
	loggerInit(sessionId)
	timePeriod := readConfig().TimePeriod

	logger.Println("Scrapper started.")

	nicknames := inputReader()

	for i, nickname := range nicknames {
		logger.Println("Scrapper started for: ", nickname, "(", i+1, "/", len(nicknames), ")")
		scrapedStreamer := scrapper(nickname, timePeriod)
		streamersDataArray = append(streamersDataArray, scrapedStreamer)
	}

	logger.Println("Scrapper finished.")
	outputWriter(streamersDataArray, sessionId)

}

/*
This function scraps the streamer data from the website.
With we can fulfill all the fields of the StreamerData struct except for the GamesPlayed field.
*/
func scrapperStreamerData(nickname string, timePeriod int, streamer *StreamerData) {
	url := fmt.Sprintf("https://sullygnome.com/channel/%s/%d", nickname, timePeriod)
	res, err := http.Get(url)

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

	/*
		Followers and Language
	*/
	doc.Find(".MiddleSubHeaderRow").Each(func(i int, s *goquery.Selection) {
		followersDiv := s.Find("[title='Total number of followers']")
		followersValue := parseValue(followersDiv.Next().Text())

		languageDiv := s.Find("[title='Broadcast language']")
		languageValue := languageDiv.Next().Text()

		streamer.Language = languageValue
		streamer.Followers = followersValue
	})
	/*
		Streams data
	*/
	doc.Find("[class*='InfoStatPanelWrapper']").Each(func(i int, s *goquery.Selection) {
		fieldName := s.Find(".InfoStatPanelBRCell").Text()
		fieldValue := s.Find(".InfoStatPanelTLCell").Text()
		setFiled(streamer, parseName(fieldName), parseValue(fieldValue))
	})
}