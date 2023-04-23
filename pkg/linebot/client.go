package linebot

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Client interface {
	BroadcastTextMessage(content string) error
}

type client struct {
	sdkClient *linebot.Client
}

func NewClient(channelSecret string, channelToken string) (Client, error) {
	sdkClient, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, fmt.Errorf("error create linebot client instance: %w", err)
	}

	return &client{
		sdkClient: sdkClient,
	}, nil
}

func (c *client) BroadcastTextMessage(content string) error {
	_, err := c.sdkClient.BroadcastMessage(linebot.NewTextMessage(content)).Do()
	if err != nil {
		return fmt.Errorf("failed to do broadcast message: %w", err)
	}
	return nil
}
