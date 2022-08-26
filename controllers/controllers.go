// package controllers
// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	h "gitbot/helpers"
// 	lib "gitbot/libraries"
// 	"gitbot/model"

// 	"github.com/go-chi/chi/v5"
// 	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// var chatId string
// var CheckUpOid primitive.ObjectID

// func HandleHook(w http.ResponseWriter, r *http.Request) {
// 	chiParam := chi.URLParam(r, "id")

// 	var findRes bson.M
// 	uid, err := primitive.ObjectIDFromHex(chiParam)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	err = h.GroupCol.FindOne(context.TODO(), bson.D{{Key: "_id", Value: uid}}).Decode(&findRes)
// 	if err == mongo.ErrNoDocuments {
// 		w.WriteHeader(404)
// 		return
// 	}
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}

// 	var res model.GroupDocument
// 	bb, _ := bson.Marshal(findRes)
// 	bson.Unmarshal(bb, &res)

// 	body, _ := io.ReadAll(r.Body)
// 	var pay model.ObjectKind
// 	json.Unmarshal(body, &pay)
// 	lib.SendTelegramMessage(pay, body, res.ChatId)
// }
//