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
	case "merge_request":
		var p webhook.MergeRequestEventsLoad
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.MergeRequestEventsMsg, p.User.Username, p.ObjectAttributes.SourceBranch, p.User.Username, p.Project.Name, p.Project.Homepage)
		url, text = p.ObjectAttributes.URL, "Open Request"
	case "pipeline":
		var p webhook.PipelineEventsLoad
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.PipelineEventsMsg, p.User.Username, p.ObjectAttributes.Ref, p.User.Username, p.Project.Name, p.Project.DefaultBranch, p.ObjectAttributes.Status)
		url, text = p.Project.WebURL, "Open Request"
	case "deployment":
		var p webhook.DeploymentEventsLoad
		err = json.Unmarshal(body, &p)
		dt = fmt.Sprintf(mdl.DeployEventsMesg, p.User.Username, p.Project.DefaultBranch, p.User.Username, p.Project.Name, p.Project.Homepage, p.Status)
		url, text = p.Project.WebURL, "Open Deloyment"
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
