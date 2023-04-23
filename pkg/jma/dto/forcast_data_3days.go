package dto

import (
	"fmt"
	"time"
)

// 直近の天気予報
type ForcastData3Days struct {
	PublishingOffice string    `json:"publishingOffice"`
	ReportDateTime   time.Time `json:"reportDateTime"`
	// 0要素目：TimeSeries3DaysWeathers
	// 1要素目：TimeSeries3DaysPops
	// 2要素目：TimeSeries3DaysTemps
	TimeSeries [3]map[string]any `json:"timeSeries"`
}

// 地域ごとの向こう3日間の天気
type TimeSeries3DaysWeathers struct {
	TimeDefines []time.Time `json:"timeDefines"` // 向こう3日間の日時
	Areas       []struct {
		Area         Area     `json:"area"`         // 地域
		WeatherCodes []string `json:"weatherCodes"` // 3日間の天気コード
		Weathers     []string `json:"weathers"`     // 3日間の天気
		Winds        []string `json:"winds"`        // 3日間の風
		Waves        []string `json:"waves"`        // 3日間の波
	} `json:"areas"`
}

// 地域ごとの向こう3日間の6時間毎の降水確率
type TimeSeries3DaysPops struct {
	TimeDefines []time.Time `json:"timeDefines"` // 向こう3日間の6時間毎の日時
	Areas       []struct {
		Area Area     `json:"area"` // 地域
		Pops []string `json:"pops"` // 向こう3日間の6時間毎の降水確率
	} `json:"areas"`
}

// 地域ごとの今日、明日の朝の最低気温と日中の最高気温
type TimeSeries3DaysTemps struct {
	TimeDefines []time.Time `json:"timeDefines"` // 今日、明日の「朝の最低気温」と「日中の最高気温」の日時
	Areas       []struct {
		Area  Area     `json:"area"`  // 地域
		Temps []string `json:"temps"` // 今日、明日の「朝の最低気温」と「日中の最高気温」
	} `json:"areas"`
}

func (d ForcastData3Days) TimeSeries3DaysWeathers() (TimeSeries3DaysWeathers, error) {
	var timeSeries3DaysWeathers TimeSeries3DaysWeathers
	if err := mapToStruct(d.TimeSeries[0], &timeSeries3DaysWeathers); err != nil {
		return timeSeries3DaysWeathers, fmt.Errorf("failed to convert TimeSeries3DaysWeathers from map to struct: %w", err)
	}
	return timeSeries3DaysWeathers, nil
}

func (d ForcastData3Days) TimeSeries3DaysPops() (TimeSeries3DaysPops, error) {
	var timeSeries3DaysPops TimeSeries3DaysPops
	if err := mapToStruct(d.TimeSeries[1], &timeSeries3DaysPops); err != nil {
		return timeSeries3DaysPops, fmt.Errorf("failed to convert TimeSeries3DaysPops from map to struct: %w", err)
	}
	return timeSeries3DaysPops, nil
}

func (d ForcastData3Days) TimeSeries3DaysTemps() (TimeSeries3DaysTemps, error) {
	var timeSeries3DaysTemps TimeSeries3DaysTemps
	if err := mapToStruct(d.TimeSeries[2], &timeSeries3DaysTemps); err != nil {
		return timeSeries3DaysTemps, fmt.Errorf("failed to convert TimeSeries3DaysPops from map to struct: %w", err)
	}
	return timeSeries3DaysTemps, nil
}
