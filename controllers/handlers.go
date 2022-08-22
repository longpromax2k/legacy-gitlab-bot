package controllers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	h "github.com/tatsuxyz/GitLabHook/helpers"
	lib "github.com/tatsuxyz/GitLabHook/libraries"
	"github.com/tatsuxyz/GitLabHook/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleHook(w http.ResponseWriter, r *http.Request) {
	chiParam := chi.URLParam(r, "id")

	var findRes bson.M
	uid, err := primitive.ObjectIDFromHex(chiParam)
	if err != nil {
		log.Panic(err)
	}

	err = h.GroupCol.FindOne(context.TODO(), bson.D{{Key: "_id", Value: uid}}).Decode(&findRes)
	if err == mongo.ErrNoDocuments {
		log.Fatal(err)
		return
	}
	if err != nil {
		log.Panic(err)
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
