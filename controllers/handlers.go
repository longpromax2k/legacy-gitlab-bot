package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	h "github.com/tatsuxyz/GitLabHook/helpers"
	lib "github.com/tatsuxyz/GitLabHook/libraries"
	"github.com/tatsuxyz/GitLabHook/model"
	"go.etcd.io/bbolt"
)

func HandleHook(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")
	s := strings.Split(urlId, ".")

	chatId, token := s[0], s[1]

	Db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("gitlabhook"))
		v := b.Get([]byte(chatId))

		if v != nil {
			if string(v[:]) == token {
				body, _ := ioutil.ReadAll(r.Body)
				var pay model.ObjectKind
				json.Unmarshal(body, &pay)
				lib.SendTelegramMessage(pay, body, chatId)
			} else {
				w.WriteHeader(404)
			}
		} else {
			w.WriteHeader(404)
		}
		return nil
	})
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
