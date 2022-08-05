package command

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

func CreateLinkAtStart(up *tgbot.Update, msg *tgbot.MessageConfig) {
	cid := strconv.Itoa(int(up.Message.Chat.ID))

	// check existed data
	var re bson.M
	err := h.Col.FindOne(context.TODO(), bson.D{{Key: "chatId", Value: cid}}).Decode(&re)
	if err != mongo.ErrNoDocuments {
		log.Printf("Document existed.\n")
		msg.Text = model.ChatExistMsg
		return
	}

	// insert new value into database
	doc := bson.D{{Key: "chatId", Value: cid}}

	result, _ := h.Col.InsertOne(context.TODO(), doc)
	oid := result.InsertedID.(primitive.ObjectID)
	log.Printf("Inserted doc with id: %v.\n", result.InsertedID)

	msg.Text = fmt.Sprintf(model.ChatInsertMsg, os.Getenv("HOST_URL"), oid.Hex())
}
