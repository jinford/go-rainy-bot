package entity

import (
	"errors"
	"time"
)

type Pops []PopByTime

type PopByTime struct {
	timeDefine time.Time
	value      string
}

func NewPops(timeDefineList []time.Time, popList []string) (Pops, error) {
	if len(timeDefineList) != len(popList) {
		return nil, errors.New("timeDefineList and popList should be same length")
	}

	list := make([]PopByTime, len(timeDefineList))
	for i, timeDefine := range timeDefineList {
		list = append(list, NewPopByTime(timeDefine, popList[i]))
	}

	return list, nil
}

func NewPopByTime(timeDefine time.Time, value string) PopByTime {
	return PopByTime{
		timeDefine: timeDefine,
		value:      value,
	}
}

func (p PopByTime) Location() *time.Location {
	return p.timeDefine.Location()
}

func (p PopByTime) Year() int {
	return p.timeDefine.Year()
}

func (p PopByTime) Month() time.Month {
	return p.timeDefine.Month()
}

func (p PopByTime) Day() int {
	return p.timeDefine.Day()
}

func (p PopByTime) Hour() int {
	return p.timeDefine.Hour()
}

func (p PopByTime) Value() string {
	return p.value
}
