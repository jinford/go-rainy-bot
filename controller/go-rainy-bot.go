package controller

import (
	"context"

	"github.com/jinford/go-rainy-bot/entity"
	"github.com/jinford/go-rainy-bot/pkg/jma"
	"github.com/jinford/go-rainy-bot/pkg/linebot"
)

type GoRainyBot struct {
	jmaClient     jma.Client
	linebotClient linebot.Client
}

func NewGoRainyBot(jmaClient jma.Client, linebotClient linebot.Client) *GoRainyBot {
	return &GoRainyBot{
		jmaClient:     jmaClient,
		linebotClient: linebotClient,
	}
}

func (b GoRainyBot) Exec(ctx context.Context) error {
	// 天気予報情報を取得
	forcastData, err := b.jmaClient.GetForcastData(context.Background())
	if err != nil {
		return err
	}

	// 今日の天気予報を取得
	todayForcast, err := entity.NewTodayForcast(forcastData)
	if err != nil {
		return err
	}

	// 今日の天気が雨なら通知する
	if !todayForcast.IsRainy() {
		return nil
	}

	if err := b.linebotClient.BroadcastTextMessage(todayForcast.TextMessage()); err != nil {
		return err
	}

	return nil
}
