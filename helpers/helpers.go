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
	Bot *tgbot.BotAPI
	Db  *mongo.Client
	err error
)

var (
	GroupCol   *mongo.Collection
	CheckUpCol *mongo.Collection
)

var (
	Port    string
	HostUrl string
	UrlPath string

	botToken string
	mongoURI string
)

func init() {
	// Load Environment Variable
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found\n")
	}

	Port = os.Getenv("PORT")
	HostUrl = os.Getenv("HOST_URL")
	UrlPath = os.Getenv("URL_PATH")
	botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	mongoURI = os.Getenv("MONGO_URI")

	// Load bot instance
	Bot, err = tgbot.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
		return
	}
	// Debug
	Bot.Debug = true

	// Load database
	if mongoURI == "" {
		log.Fatalf("You must set your 'MONGODB_URI' environmental variable.\n")
		return
	}
	Db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Panic(err)
	}

	GroupCol = Db.Database("app").Collection("group")
	CheckUpCol = Db.Database("app").Collection("checkup")
}
