package controllers

import (
	"gitbot/configs"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LoadBot() (*tgbotapi.BotAPI, error) {
	cfg := configs.GetConfig()
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return bot, err
}
