package scrapper

import (
	"io"
	"os"
	"strings"
)

func inputReader() []string{
	file, err := os.Open("data/input/input.txt")
	if err != nil {
		logger.Println("Error while opening the file: ", err)
	}
	defer file.Close()
	logger.Println("Nicknames file opened successfully.")

	nicknamesByteArray, err := io.ReadAll(file)

	if err != nil {
		logger.Println("Error while reading the file: ", err)
	}

	nicknamesString := string(nicknamesByteArray[:])
	nicknamesStringArray := strings.Split(nicknamesString, "\n") 

	return nicknamesStringArray
}
