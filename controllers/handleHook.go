package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

	// Check request body
	body, _ := ioutil.ReadAll(r.Body)
	token := r.Header.Get("X-GitLab-Token")
	if token != os.Getenv("SECRET_TOKEN") {
		log.Print("Secret token mismatch")
	}

	// Check object kind and send message
	var pay mdl.ObjectKind
	json.Unmarshal(body, &pay)
	SendTelegramMessage(pay, body)

	// Serve response
	w.Header().Set("Content-Type", "application/json")
	data := struct {
		Message string `json:"message"`
	}{
		Message: "ok",
	}
	json.NewEncoder(w).Encode(data)
}
