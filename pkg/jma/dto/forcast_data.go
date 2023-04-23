package dto

import "fmt"

// 0要素目：ForcastData7Days
// 1要素目：ForcastData7Days
type ForcastData [2]map[string]any

func (d ForcastData) ForcastData3Days() (ForcastData3Days, error) {
	var forcastData3Days ForcastData3Days
	if err := mapToStruct(d[0], &forcastData3Days); err != nil {
		return forcastData3Days, fmt.Errorf("failed to convert ForcastData3Days from map to struct: %w", err)
	}
	return forcastData3Days, nil
}

func (d ForcastData) ForcastData7Days() (ForcastData7Days, error) {
	var forcastData7Days ForcastData7Days
	if err := mapToStruct(d[1], &forcastData7Days); err != nil {
		return forcastData7Days, fmt.Errorf("failed to convert ForcastData7Days from map to struct: %w", err)
	}
	return forcastData7Days, nil
}
