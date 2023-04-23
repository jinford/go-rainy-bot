package jma

import (
	"context"
	"net/http"

	"github.com/jinford/go-rainy-bot/pkg/jma/dto"
)

type Client interface {
	GetForcastData(ctx context.Context) (dto.ForcastData, error)
}

type client struct {
	httpClient *http.Client
}

func NewClient() Client {
	return &client{
		httpClient: http.DefaultClient,
	}
}
