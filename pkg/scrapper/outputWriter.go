package scrapper

import (
	"encoding/json"
	"os"
	"time"
)

func outputWriter(streamersDataArray []StreamerData, sessionId string) {
	file, err := os.OpenFile("data/output/json/output_" + sessionId + ".json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		logger.Println("Error while opening the file: ", err)
	}

	defer file.Close()

	jsonData, err := json.MarshalIndent(streamersDataArray, "", "    ")
	
	if err != nil {
		logger.Println("Error while marshaling the data: ", err)
	}

	_, err = file.Write(jsonData)

	if err != nil {
		logger.Println("Error while writing the data: ", err)
	}

	logger.Println("Data written to the file.")
}

func sessionId() string {
	sessionId := time.Now().Format("2006-01-02 15_04_05")
	return sessionId
}
