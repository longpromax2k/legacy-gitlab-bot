package controllers

import (
	"context"

	h "github.com/tatsuxyz/GitLabHook/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddNewIDchat(IDchat string) {
	doc := bson.D{{Key: "chatId", Value: IDchat}}

	_, err := h.Col.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}
}

// remove URLwebhook after stop command
func removeIDchat(db *mongo.Collection, IDchat string) {
	filter := bson.D{{Key: "chatId", Value: IDchat}}

	_, err := db.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
}

//check available Idchat
func checkIDchat(db *mongo.Collection, IDchat string) bool {
	var result bson.M
	err := db.FindOne(context.TODO(), bson.D{{Key: "chatId", Value: IDchat}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false
		}
		return true
	}
	panic(err)
}
