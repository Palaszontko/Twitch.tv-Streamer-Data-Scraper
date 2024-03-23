package scrapper

import (
	"encoding/json"
	"io"
	"os"
	"slices"
)

type Config struct {
	TimePeriod int `json:"timePeriod"`
}

func readConfig() Config{ 
	var config Config

	file := openFile()

	json.Unmarshal(file, &config)

	avilaibleTimePeriods := []int{3, 7, 14, 30, 90, 180, 365}
	
	if !slices.Contains(avilaibleTimePeriods, config.TimePeriod){
		config.TimePeriod = 7
	}

	return config
}

func openFile() []byte{ 
	jsonFile, err := os.Open("configs/config.json")

	if err != nil {
		logger.Println("Error while opening the file: ", err)
	}
	logger.Println("Config file opened successfully.")
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	return byteValue
} 