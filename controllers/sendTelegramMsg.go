package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	mdl "github.com/tatsuxyz/GitLabHook/model"
	webhook "github.com/tatsuxyz/GitLabHook/model/webhook"
)

func SendTelegramMessage(pay mdl.ObjectKind, body []byte) {
	cid, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	var chatId = int64(cid)
	var dt, url, text string
	var err error

	switch pay.ObjectKind {
	case "push":
		var p webhook.PushEventPayload
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.PushEventMsg, p.UserUsername, p.Ref, p.UserUsername, p.Project.Name, p.Project.Homepage, p.Commits[0].Message)
		url, text = p.Commits[0].URL, "Open Commit"
	default:
		log.Fatalf("Invalid Event\n")
		return
	}

	if err != nil {
		log.Fatalf("Json unmarshal error, %v\n", err)
		return
	}

	msg := tgbot.NewMessage(chatId, dt)
	msg.ParseMode = "markdown"
	msg.ReplyMarkup = tgbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbot.InlineKeyboardButton{
			{
				tgbot.InlineKeyboardButton{
					Text: text,
					URL:  &url,
				},
			},
		},
	}
	Bot.Send(msg)
}
