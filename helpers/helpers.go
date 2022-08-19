package helpers

import (
	"log"
	"os"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var (
	Bot *tgbot.BotAPI
	err error
)

func main() {
	// Load Environment Variable
	if err = godotenv.Load(); err != nil {
		log.Printf("No .env file found\n")
	}
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Load bot instance
	Bot, err = tgbot.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
		return
	}
	Bot.Debug = true
}
