package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	h "gitbot/helpers"
	lib "gitbot/libraries"
	"gitbot/model"

	"github.com/go-chi/chi/v5"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var chatId string
var CheckUpOid primitive.ObjectID

func HandleHook(w http.ResponseWriter, r *http.Request) {
	chiParam := chi.URLParam(r, "id")

	var findRes bson.M
	uid, err := primitive.ObjectIDFromHex(chiParam)
	if err != nil {
		log.Panic(err)
	}

	err = h.GroupCol.FindOne(context.TODO(), bson.D{{Key: "_id", Value: uid}}).Decode(&findRes)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(404)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var res model.GroupDocument
	bb, _ := bson.Marshal(findRes)
	bson.Unmarshal(bb, &res)

	body, _ := io.ReadAll(r.Body)
	var pay model.ObjectKind
	json.Unmarshal(body, &pay)
	lib.SendTelegramMessage(pay, body, res.ChatId)
}

func HandleCommand() {
	var r bson.M
	err := h.CheckUpCol.FindOne(context.TODO(), bson.D{{Key: "status", Value: true}}).Decode(&r)
	if err != mongo.ErrNoDocuments {
		log.Printf("There's an existed instance running, no check needed.")
		for {
			err := h.CheckUpCol.FindOne(context.TODO(), bson.D{{Key: "status", Value: true}}).Decode(&r)
			if err != mongo.ErrNoDocuments {
				continue
			}
			break
		}
	}
	doc := bson.D{{Key: "status", Value: true}}
	res, err := h.CheckUpCol.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	CheckUpOid = res.InsertedID.(primitive.ObjectID)
	u := tgbot.NewUpdate(0)
	u.Timeout = 60

	ups := h.Bot.GetUpdatesChan(u)

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
			msg.Text = model.ChatNotCmdMsg
		}

		msg.ParseMode = "markdown"
		if _, err := h.Bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}

type User struct {
	ID primitive.ObjectID `bson:"_id" json:"id,omitempty"`
}

func CommandStart(up *tgbot.Update, msg *tgbot.MessageConfig) {
	chatId = strconv.Itoa(int(up.Message.Chat.ID))

	var r User

	err := h.GroupCol.FindOne(context.TODO(), bson.D{{Key: "chatId", Value: chatId}}).Decode(&r)
	log.Println(err)
	if err != mongo.ErrNoDocuments {
		msg.Text = fmt.Sprintf(model.ChatExistMsg, h.HostUrl, h.UrlPath, r.ID.Hex())
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
