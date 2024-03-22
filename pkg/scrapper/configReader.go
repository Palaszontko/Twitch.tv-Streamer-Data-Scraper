package scrapper

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	TimePeriod int `json:"timePeriod"`
	GamesAmount int `json:"gamesAmount"`
}

func ReadConfig() Config{ 
	var config Config

	file := openFile()

	json.Unmarshal(file, &config)

	return config
}

func openFile() []byte{ 
	jsonFile, err := os.Open("config/config.json")

	if err != nil {
		logger.Println("Error while opening the file: ", err)
	}
	logger.Println("Config file opened successfully.")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
} 