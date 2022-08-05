package helpers

import (
	"context"
	"log"
	"os"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Bot    *tgbot.BotAPI
	Client *mongo.Client
	Col    *mongo.Collection
	err    error
)

func init() {
	// Load Environment Variable
	if err = godotenv.Load(); err != nil {
		log.Printf("No .env file found\n")
	}
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	mongoURI := os.Getenv("MONGO_URI")

	// Telegram: Load bot instance
	Bot, err = tgbot.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
		return
	}

	// MongoDB: Connect to cluster
	if mongoURI == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.\n")
		return
	}
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Panic(err)
	}

	// Use collection
	Col = Client.Database("gitlabhook").Collection("gitlabhook")
}
