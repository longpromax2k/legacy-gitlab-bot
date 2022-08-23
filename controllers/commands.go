package controllers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	h "github.com/tatsuxyz/GitLabHook/helpers"
	"github.com/tatsuxyz/GitLabHook/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var chatId string

func CommandStart(up *tgbot.Update, msg *tgbot.MessageConfig) {
	chatId = strconv.Itoa(int(up.Message.Chat.ID))
	var r bson.M

	err := h.GroupCol.FindOne(context.TODO(), bson.D{{Key: "chatId", Value: chatId}}).Decode(&r)
	if err != mongo.ErrNoDocuments {
		msg.Text = model.ChatExistMsg
		return
	}

	doc := bson.D{{Key: "chatId", Value: chatId}}
	res, err := h.GroupCol.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	oid := res.InsertedID.(primitive.ObjectID)

	msg.Text = fmt.Sprintf(model.ChatInsertMsg, h.HostUrl, h.UrlPath, oid.Hex())
}

func CommandDrop(up *tgbot.Update, msg *tgbot.MessageConfig) {
	chatId = strconv.Itoa(int(up.Message.Chat.ID))

	f := bson.D{{Key: "chatId", Value: chatId}}
	if _, err := h.GroupCol.DeleteOne(context.TODO(), f); err != nil {
		log.Panic(err)
	}

	msg.Text = model.ChatDropMsg
}
