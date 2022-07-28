package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	mdl "github.com/tatsuxyz/GitLabHook/model"
)

var (
	Bot *tgbot.BotAPI
)

func HandleWebHook(w http.ResponseWriter, r *http.Request) {
	// Only allow POST request
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	// Serve response
	w.Header().Set("Content-Type", "application/json")
	data := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}
	json.NewEncoder(w).Encode(data)

	// Request body
	body, _ := ioutil.ReadAll(r.Body)
	// token := r.Header.Get("X-GitLab-Token")
	// if token != os.Getenv("SECRET_TOKEN") {
	// 	log.Print("Secret token mismatch")
	// 	return
	// }

	// JSON parses
	var pay mdl.Gitlab
	err := json.Unmarshal(body, &pay)
	if err != nil {
		fmt.Printf("[GitLabHook] Json unmarshal error, %v\n", err)
		return
	}

	// Send message
	cid, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	var chatId = int64(cid)

	switch pay.ObjectKind {
	case "push":
		dt := fmt.Sprintf(mdl.PushEventMsg, pay.UserUsername, pay.Ref, pay.UserUsername, pay.Project.Name, pay.Project.Homepage, pay.Commits[0].Message)
		msg := tgbot.NewMessage(chatId, dt)
		msg.ParseMode = "markdown"
		msg.ReplyMarkup = tgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]tgbot.InlineKeyboardButton{
				{
					tgbot.InlineKeyboardButton{
						Text: "Open Commit",
						URL:  &pay.Commits[0].URL,
					},
				},
			},
		}
		Bot.Send(msg)
	default:
		log.Fatalf("Invalid Event")
		return
	}
}
