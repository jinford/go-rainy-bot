package dto

import (
	"fmt"
	"time"
)

// 週間予報
type ForcastData7Days struct {
	PublishingOffice string    `json:"publishingOffice"`
	ReportDateTime   time.Time `json:"reportDateTime"`
	// 0要素目：TimeSeries7DaysWeathers
	// 1要素目：TimeSeries7DaysTemps
	TimeSeries    [2]map[string]any `json:"timeSeries"`
	TempAverage   TempAverage       `json:"tempAverage"`
	PrecipAverage PrecipAverage     `json:"precipAverage"`
}

// 地域ごとの向こう7日間の天気
type TimeSeries7DaysWeathers struct {
	TimeDefines []string `json:"timeDefines"` // 向こう7日間の日時
	Areas       []struct {
		Area          Area     `json:"area"`          // 地域
		WeatherCodes  []string `json:"weatherCodes"`  // 天気コード
		Pops          []string `json:"pops"`          // 降水確率(%)
		Reliabilities []string `json:"reliabilities"` // 信頼度
	} `json:"areas"`
}

// 地域ごとの向こう7日間の気温
type TimeSeries7DaysTemps struct {
	TimeDefines []string `json:"timeDefines"` // 向こう7日間の日時
	Areas       []struct {
		Area          Area     `json:"area"`          // 地域
		TempsMin      []string `json:"tempsMin"`      // 最低気温
		TempsMinUpper []string `json:"tempsMinUpper"` // 最低気温の予測下限
		TempsMinLower []string `json:"tempsMinLower"` // 最低気温の予測上限
		TempsMax      []string `json:"tempsMax"`      // 最大気温
		TempsMaxUpper []string `json:"tempsMaxUpper"` // 最大気温の予測下限
		TempsMaxLower []string `json:"tempsMaxLower"` // 最大気温の予測上限
	}
}

// 地域ごとの向こう一週間（明日から７日先まで）の気温の平年値
type TempAverage struct {
	Areas []struct {
		Area Area   `json:"area"` // 地域
		Min  string `json:"min"`  // 最低気温
		Max  string `json:"max"`  // 最大気温
	} `json:"areas"`
}

// 地域ごとの向こう一週間（明日から７日先まで）の降水量の７日間合計の平年値
type PrecipAverage struct {
	Areas []struct {
		Area Area   `json:"area"` // 地域
		Min  string `json:"min"`  // 最低値
		Max  string `json:"max"`  // 最大値
	} `json:"areas"`
}

func (d ForcastData7Days) TimeSeries7DaysWeathers() (TimeSeries7DaysWeathers, error) {
	var timeSeries7DaysWeathers TimeSeries7DaysWeathers
	if err := mapToStruct(d.TimeSeries[0], &timeSeries7DaysWeathers); err != nil {
		return timeSeries7DaysWeathers, fmt.Errorf("failed to convert TimeSeries7DaysWeathers from map to struct: %w", err)
	}
	return timeSeries7DaysWeathers, nil
}

func (d ForcastData7Days) TimeSeries7DaysTemps() (TimeSeries7DaysTemps, error) {
	var timeSeries7DaysTemps TimeSeries7DaysTemps
	if err := mapToStruct(d.TimeSeries[0], &timeSeries7DaysTemps); err != nil {
		return timeSeries7DaysTemps, fmt.Errorf("failed to convert TimeSeries7DaysTemps from map to struct: %w", err)
	}
	return timeSeries7DaysTemps, nil
}
