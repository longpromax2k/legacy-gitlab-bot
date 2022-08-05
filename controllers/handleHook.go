package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	h "github.com/tatsuxyz/GitLabHook/helpers"
	mdl "github.com/tatsuxyz/GitLabHook/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDocument struct {
	ChatId string `json:"chatId"`
}

func HandleHook(w http.ResponseWriter, r *http.Request) {
	urlId := chi.URLParam(r, "id")

	var result bson.M
	id, err := primitive.ObjectIDFromHex(urlId)
	if err != nil {
		log.Panic(err)
	}

	err = h.Col.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Printf("No document was found with the id %s\n", urlId)
		w.WriteHeader(404)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		log.Panic(err)
	}

	var d MongoDocument
	bb, _ := bson.Marshal(result)
	bson.Unmarshal(bb, &d)

	log.Printf("%s\n", d.ChatId)

	body, _ := ioutil.ReadAll(r.Body)
	// Check object kind and send message
	var pay mdl.ObjectKind
	json.Unmarshal(body, &pay)
	SendTelegramMessage(pay, body, d.ChatId)
}

func HandleWebHook(w http.ResponseWriter, r *http.Request) {

}
