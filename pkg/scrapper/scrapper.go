package scrapper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func scrapper(nickname string, timePeriod int) StreamerData {
	streamer := StreamerData{
		Nickname: nickname,
	}
	scrapperStreamerData(streamer.Nickname, timePeriod, &streamer)
	logger.Println("Streamer data scrapped.")
	scrapperGamesData(streamer.Nickname, timePeriod, &streamer)
	logger.Println("Games data scrapped.")
	
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
	csvWriter(streamersDataArray, sessionId)
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

/* 
This function scraps the games data from the website.
*/
func scrapperGamesData(nickname string, timePeriod int, streamer *StreamerData) {
	type Response struct {
		Data []struct {
			GameName       string `json:"gamesplayed"`
			StreamTime     int    `json:"streamtime"`
			TotalWatchTime int    `json:"viewtime"`
			AverageViewers int    `json:"avgviewers"`
			PeakViewers    int    `json:"maxviewers"`
		} `json:"data"`
	}

	id := scrapeAndParseStreamerId(nickname)
	
	if id == -1 {
		fmt.Println("ID not found for the streamer: ", nickname)
		fmt.Println("Games data not provided.")
		return
	}
	
	url := fmt.Sprintf("https://sullygnome.com/api/tables/channeltables/games/%d/%d/%s/1/2/desc/0/25", timePeriod, id, " ")
	
	res, err := http.Get(url)
	
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	
	defer res.Body.Close()
	
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	
	jsonBody, err := ioutil.ReadAll(res.Body)
	
	if err != nil {
		log.Fatal("Error while reading JSON:", err)
	}
	
	var response Response
	
	err = json.Unmarshal(jsonBody, &response)
	
	if err != nil {
		log.Fatal("Error while unmarshalling JSON:", err)
	}
	
	for _, data := range response.Data {
	
		index := strings.Index(data.GameName, "|") 
		
		if index != -1 {
			data.GameName = data.GameName[:index]	
		}
		data.StreamTime = data.StreamTime / 60
		data.TotalWatchTime = data.TotalWatchTime / 60

		gamePlayed := GamePlayed{
			GameName:       data.GameName,
			StreamTime:     data.StreamTime,
			TotalWatchTime: data.TotalWatchTime,
			AverageViewers: data.AverageViewers,
			PeakViewers:    data.PeakViewers,
		}

		streamer.GamesPlayed = append(streamer.GamesPlayed, gamePlayed)
	}

	
}

/*
This function finds the "special" ID on the website. This ID is used to get the data about the games played by the streamer
*/
func scrapeAndParseStreamerId(nickname string) int {
	url := fmt.Sprintf("https://sullygnome.com/channel/%s/games", nickname)
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	_ = doc

	if err != nil {
		log.Fatal("Error while reading HTML:", err)
	}

	text := doc.Find("body script").Text()
	lines := strings.Split(text, "\n")

	type PageInfo struct {
		ID int `json:"id"`
	}

	for _, line := range lines {
		if strings.Contains(line, "var PageInfo = ") {
			variable := strings.Replace(line, " ", "", -1)
			variable = strings.Replace(variable, "varPageInfo=", "", -1)
			variable = strings.Replace(variable, ";", "", -1)

			var pageInfo PageInfo

			err := json.Unmarshal([]byte(variable), &pageInfo)

			if err != nil {
				log.Fatal("Error while unmarshalling JSON:", err)
			}

			return pageInfo.ID
		}
	}

	return -1
}
