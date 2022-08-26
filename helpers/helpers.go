package helpers

import (
	"context"
	"gitbot/util"
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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

func init() {

	// Load Environment Variable
	// if err := godotenv.Load(); err != nil {
	// 	log.Printf("No .env file found\n")
	// }
	config, err := util.LoadConfig(".")

	// Port = os.Getenv("PORT")
	// HostUrl = os.Getenv("HOST_URL")
	// UrlPath = os.Getenv("URL_PATH")
	// botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	// mongoURI = os.Getenv("MONGO_URI")
	// Load bot instance
	Bot, err = tgbot.NewBotAPI(config.BotToken)
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
