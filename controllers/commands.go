package controllers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	h "github.com/tatsuxyz/GitLabHook/helpers"
	"github.com/tatsuxyz/GitLabHook/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}

var chatId string

func CommandStart(up *tgbot.Update, msg *tgbot.MessageConfig) {
	chatId = strconv.Itoa(int(up.Message.Chat.ID))
	var re User
	cid := strconv.Itoa(int(up.Message.Chat.ID))
	//TODO: check existed data
	err := h.Col.FindOne(context.TODO(), bson.D{{Key: "chatId", Value: "cid"}}).Decode(&re)
	//err.Decode(&re)
	if err != mongo.ErrNoDocuments {
		log.Printf("Document existed.\n")
		msg.Text = fmt.Sprintf(model.ChatExistMsg, os.Getenv("HOST_URL"), re.ID)
		return
	}

	// TODO: insert new value into database
	doc := bson.D{{Key: "chatId", Value: cid}}
	result, _ := h.Col.InsertOne(context.TODO(), doc)
	oid := result.InsertedID.(primitive.ObjectID)
	log.Printf("Inserted doc with id:%v.\n", result.InsertedID)
	msg.Text = fmt.Sprintf(model.ChatInsertMsg, os.Getenv("HOST_URL"), oid.Hex())

}

func CommandDrop(up *tgbot.Update, msg *tgbot.MessageConfig) {
	chatId = strconv.Itoa(int(up.Message.Chat.ID))

	// TODO
	filter := bson.D{{Key: "chatId", Value: chatId}}
	_, err := h.Col.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	msg.Text = model.ChatDropMsg
}
