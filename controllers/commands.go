package controllers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"gitbot/configs"
	"gitbot/models"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	lock   = &sync.Mutex{}
	chatId string
)

type User struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}

func CommandStart(up *tgbot.Update, msg *tgbot.MessageConfig) {
	chatId = strconv.Itoa(int(up.Message.Chat.ID))
	cfg := configs.GetConfig()
	var r User

	err := GetCol().FindOne(context.TODO(), bson.D{{Key: "chatId", Value: chatId}}).Decode(&r)
	log.Println(err)
	if err != mongo.ErrNoDocuments {
		msg.Text = fmt.Sprintf(models.ChatExistMsg, cfg.HostURL, cfg.PathURL, r.ID.Hex())
		return
	}

	doc := bson.D{{Key: "chatId", Value: chatId}}
	res, err := GetCol().InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	oid := res.InsertedID.(primitive.ObjectID)
	msg.Text = fmt.Sprintf(models.ChatInsertMsg, cfg.HostURL, cfg.PathURL, oid.Hex())
}

func CommandDrop(up *tgbot.Update, msg *tgbot.MessageConfig) {
	chatId = strconv.Itoa(int(up.Message.Chat.ID))

	f := bson.D{{Key: "chatId", Value: chatId}}
	if _, err := GetCol().DeleteOne(context.TODO(), f); err != nil {
		log.Panic(err)
	}

	msg.Text = models.ChatDropMsg
}

func HandleCommand() {
	lock.Lock()
	defer lock.Unlock()

	var r bson.M

	for {
		err := GetCol().FindOne(context.TODO(), bson.D{{Key: "status", Value: true}}).Decode(&r)
		if err != mongo.ErrNoDocuments {
			time.Sleep(6 * time.Second)
			continue
		}
		break
	}

	doc := bson.D{{Key: "status", Value: true}}
	res, err := GetCol().InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	configs.SetCheckStatus(res.InsertedID.(primitive.ObjectID))

	u := tgbot.NewUpdate(0)
	u.Timeout = 60

	bot, err := LoadBot()
	if err != nil {
		log.Fatalln(err)
	}

	ups := bot.GetUpdatesChan(u)

	for up := range ups {
		if up.Message == nil {
			continue
		}
		if !up.Message.IsCommand() {
			continue
		}

		msg := tgbot.NewMessage(up.Message.Chat.ID, "")

		switch up.Message.Command() {
		case "start":
			CommandStart(&up, &msg)
		case "drop":
			CommandDrop(&up, &msg)
		default:
			msg.Text = models.ChatNotCmdMsg
		}

		msg.ParseMode = "markdown"
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
