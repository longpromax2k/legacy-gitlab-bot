package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	lib "github.com/tatsuxyz/GitLabHook/library"
	"github.com/tatsuxyz/GitLabHook/model"
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
	token := r.Header.Get("X-GitLab-Token")
	if token != os.Getenv("SECRET_TOKEN") {
		log.Print("Secret token mismatch")
	}

	// JSON parses
	var pay model.Gitlab
	err := json.Unmarshal(body, &pay)
	if err != nil {
		fmt.Printf("[GitLabHook] Json unmarshal error, %v", err)
	}

	// Response to push action
	if pay.ObjectKind == "push" {
		lib.PostMessage(pay)
	}
}
