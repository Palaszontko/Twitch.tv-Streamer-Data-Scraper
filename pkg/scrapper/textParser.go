package scrapper

import (
	"strconv"
	"strings"
)

func parseName(text string) string {
	textArray := strings.Split(text, " ")
	if len(textArray) > 1 {
		return textArray[0] + strings.Title(textArray[1])
	} else {
		return textArray[0]
	}
}

func parseValue(text string) int {
	textValue := strings.Replace(text, ",", "", -1)
	value, err := strconv.Atoi(textValue)
	if err != nil {
		return 0
	}
	return value
}
