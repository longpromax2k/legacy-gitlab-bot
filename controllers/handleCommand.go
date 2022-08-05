package controllers

import (
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	cmd "github.com/tatsuxyz/GitLabHook/controllers/command"
	h "github.com/tatsuxyz/GitLabHook/helpers"
)

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
			cmd.CreateLinkAtStart(&up, &msg)
		case "genlink":
		case "stop":
		default:
			msg.Text = "I dunno bro"
		}

		msg.ParseMode = "markdown"
		if _, err := h.Bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
