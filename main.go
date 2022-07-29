package main

import (
	"log"
	"net/http"
	"os"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	dotenv "github.com/joho/godotenv"
	ctrl "github.com/tatsuxyz/GitLabHook/controllers"
	r "github.com/tatsuxyz/GitLabHook/routes"
)

func main() {
	// Load Environment Variable
	dotEnvErr := dotenv.Load()
	if dotEnvErr != nil {
		log.Printf("Failed to load environment variable\n")
	}

	// Load bot instance
	bot, err := tgbot.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
		return
	}
	ctrl.Bot = bot

	// Handle request and endpoints
	r.Routes()

	// Serve
	port := os.Getenv("PORT")
	log.Printf("Listening to port %s\n", port)
	http.ListenAndServe("localhost:"+port, nil)
}
