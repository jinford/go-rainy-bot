package jma

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinford/go-rainy-bot/pkg/jma/dto"
)

func (c *client) GetForcastData(ctx context.Context) (dto.ForcastData, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, forcastDateURL, http.NoBody)
	if err != nil {
		return dto.ForcastData{}, fmt.Errorf("failed to create request instance: %w", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return dto.ForcastData{}, fmt.Errorf("failed to do request: %w", err)
	}

	var forcastData dto.ForcastData
	if err := json.NewDecoder(res.Body).Decode(&forcastData); err != nil {
		return forcastData, fmt.Errorf("failed to decode response body: %w", err)
	}

	return forcastData, nil
}
