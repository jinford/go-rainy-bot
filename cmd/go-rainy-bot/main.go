package main

import (
	"context"
	"log"
	"os"

	"github.com/jinford/go-rainy-bot/controller"
	"github.com/jinford/go-rainy-bot/pkg/jma"
	"github.com/jinford/go-rainy-bot/pkg/linebot"
	"github.com/urfave/cli/v2"
)

// Flags
var (
	channelSecret string
	channelToken  string
)

func main() {
	app := &cli.App{
		Name: "go-rainy-bot",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "channelSecret",
				Aliases:     []string{"s"},
				Required:    true,
				Usage:       "channel secret for linebot",
				EnvVars:     []string{"LINEBOT_CHANNEL_SECRET"},
				Destination: &channelSecret,
			},
			&cli.StringFlag{
				Name:        "channelToken",
				Aliases:     []string{"t"},
				Required:    true,
				Usage:       "channel token for linebot",
				EnvVars:     []string{"LINEBOT_CHANNEL_TOKEN"},
				Destination: &channelToken,
			},
		},
		Action: mainAction,
	}

	if err := app.RunContext(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func mainAction(cCtx *cli.Context) error {
	jmaClient := jma.NewClient()

	linebotClient, err := linebot.NewClient(channelSecret, channelToken)
	if err != nil {
		return err
	}

	goRainyBot := controller.NewGoRainyBot(jmaClient, linebotClient)

	if err := goRainyBot.Exec(cCtx.Context); err != nil {
		return err
	}

	return nil
}
