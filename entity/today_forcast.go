package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinford/go-rainy-bot/pkg/jma/dto"
	"github.com/samber/lo"
)

type TodayForcast struct {
	areaName string
	weather  string
	pops     Pops
}

func NewTodayForcast(forcastData dto.ForcastData) (*TodayForcast, error) {
	forcastData3Days, err := forcastData.ForcastData3Days()
	if err != nil {
		return nil, err
	}

	threedaysWeathers, err := forcastData3Days.TimeSeries3DaysWeathers()
	if err != nil {
		return nil, err
	}

	threedaysPops, err := forcastData3Days.TimeSeries3DaysPops()
	if err != nil {
		return nil, err
	}

	pops, err := NewPops(threedaysPops.TimeDefines, threedaysPops.Areas[0].Pops)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	todayPops := lo.Filter(pops, func(x PopByTime, index int) bool {
		nowInLoc := now.In(x.Location())
		return x.Year() == nowInLoc.Year() &&
			x.Month() == nowInLoc.Month() &&
			x.Day() == nowInLoc.Day()
	})

	return &TodayForcast{
		areaName: threedaysWeathers.Areas[0].Area.Name,
		weather:  threedaysWeathers.Areas[0].Weathers[0],
		pops:     todayPops,
	}, err
}

func (tf *TodayForcast) IsRainy() bool {
	// 雨か雪ならtrueを返す
	return strings.ContainsAny(tf.weather, "雨雪")
}

func (tf *TodayForcast) TextMessage() string {
	message := ""
	message += fmt.Sprintf("今日の%sは%sの予報です。\n", tf.areaName, tf.weather)

	message += "\n"
	message += "■ 降水確率☔\n"
	for _, p := range tf.pops {
		message += fmt.Sprintf("%d時 %s%%\n", p.Hour(), p.Value())
	}
	message = strings.TrimSuffix(message, "\n")
	return message
}
