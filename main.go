package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	ctrl "github.com/tatsuxyz/GitLabHook/controllers"
	"github.com/tatsuxyz/GitLabHook/routes"
)

var (
	botName = os.Getenv("BOT_NAME")
	port    = os.Getenv("PORT")
)

func main() {
	// Load Environment Variable
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		fmt.Printf("[%s] Failed to load environment variable\n", botName)
	}

	// Load bot instance
	bot, err := tgbot.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
		return
	}
	ctrl.Bot = bot

	// Handle request and endpoints
	routes.Routes()

	// Serve
	fmt.Printf("[%s] Listening to port %s\n", botName, port)
	http.ListenAndServe("localhost:"+port, nil)
}
