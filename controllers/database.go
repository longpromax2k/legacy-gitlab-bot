package controllers

import (
	"context"
	"gitbot/configs"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db  *mongo.Client
	err error
)

func LoadDatabase() (*mongo.Client, error) {
	cfg := configs.GetConfig()
	if cfg.MongoURI == "" {
		log.Fatalf("You must set your 'MONGO_URI' environmental variable.\n")
	}

	db, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.MongoURI))
	return db, err
}

func CloseDatabase(ctx context.Context) error {
	checkStatus := configs.GetCheckStatus()
	if checkStatus != "000000000000000000000000" {
		f := bson.D{{Key: "_id", Value: checkStatus}}
		if _, err := GetCol().DeleteOne(context.TODO(), f); err != nil {
			log.Panic(err)
		}
	}

	if err := db.Disconnect(ctx); err != nil {
		log.Printf("database shutdown error: %v", err)
	}

	return err
}

func GetCol() *mongo.Collection {
	return db.Database("app").Collection("data")
}
