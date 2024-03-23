package scrapper

import (
	"reflect"
)

type StreamerData struct {
	Nickname        string       `json:"nickname"`
	Followers       int          `json:"followers"`
	Language        string       `json:"language"`
	AverageViewers  int          `json:"averageViewers"`
	HoursWatched    int          `json:"hoursWatched"`
	FollowersGained int          `json:"followersGained"`
	PeakViewers     int          `json:"peakViewers"`
	HoursStreamed   int          `json:"hoursStreamed"`
	Streams         int          `json:"streams"`
	GamesPlayed     []GamePlayed `json:"gamesPlayed"`
}

type GamePlayed struct {
	GameName       string `json:"gameName"`
	StreamTime     int    `json:"streamTime"`
	TotalWatchTime int    `json:"totalWatchTime"`
	AverageViewers int    `json:"averageViewers"`
	PeakViewers    int    `json:"peakViewers"`
}

/*
This function sets the value of the field in StreamerData struct.
*/
func setFiled(streamer *StreamerData, field string, value interface{}) {
	streamerReflect := reflect.ValueOf(streamer).Elem()
	fieldReflect := streamerReflect.FieldByName(field)

	if fieldReflect.IsValid() {
		switch v := value.(type) {
		case string:
			fieldReflect.SetString(v)
		case int:
			fieldReflect.SetInt(int64(v))
		default:
			logger.Println("Unknown type: ", v)
		}
	}

}
