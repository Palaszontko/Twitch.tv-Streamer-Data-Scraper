package scrapper

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func csvWriter(streamersDataArray []StreamerData, sessionId string) {
	file, err := os.OpenFile("data/output/csv/output_" + sessionId + ".csv", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger.Println("CSV file created")

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Nickname", "Followers", "Language", "Average Viewers", "Hours Watched", "Followers Gained", "Peak Viewers", "Hours Streamed", "Streams", "Game 1", "Game 2", "Game 3", "Game 4", "Game 5"})

	for _, streamerData := range streamersDataArray {
		dataToWrite := []string{
			streamerData.Nickname,
			strconv.Itoa(streamerData.Followers),
			streamerData.Language,
			strconv.Itoa(streamerData.AverageViewers),
			strconv.Itoa(streamerData.HoursWatched),
			strconv.Itoa(streamerData.FollowersGained),
			strconv.Itoa(streamerData.PeakViewers),
			strconv.Itoa(streamerData.HoursStreamed),
			strconv.Itoa(streamerData.Streams),
		}

		//write as many games as possible if more than 5 break if less than 5 is in array of games write 0
		for i := 0; i < 5; i++ {
			if i < len(streamerData.GamesPlayed) {
				dataToWrite = append(dataToWrite, streamerData.GamesPlayed[i].GameName)
			} else {
				dataToWrite = append(dataToWrite, "")
			}
		}

		writer.Write(dataToWrite)
	}
	logger.Println("Streamers data written to CSV file")
}