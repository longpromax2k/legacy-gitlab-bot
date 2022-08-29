package controllers

import (
	"context"
	"encoding/json"
	"gitbot/models"
	"io"
	"log"
	"net/http"

	lib "gitbot/libraries"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleWebHook(w http.ResponseWriter, r *http.Request) {
	chiParam := chi.URLParam(r, "id")

	var findRes bson.M
	uid, err := primitive.ObjectIDFromHex(chiParam)
	if err != nil {
		w.WriteHeader(500)
		log.Panic(err)
	}

	err = GetCol().FindOne(context.TODO(), bson.D{{Key: "_id", Value: uid}}).Decode(&findRes)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(404)
		return
	}
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var res models.GroupDocument
	bb, _ := bson.Marshal(findRes)
	bson.Unmarshal(bb, &res)

	body, _ := io.ReadAll(r.Body)
	var pay models.ObjectKind
	json.Unmarshal(body, &pay)
	lib.SendTelegramMessage(pay, body, res.ChatId)
}
