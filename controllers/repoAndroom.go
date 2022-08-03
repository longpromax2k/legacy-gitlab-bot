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
// Get repository list
// In type mongo.DB connection
//               Repo name    Room ID's
// out map, map[Repository1:[1111 2222] Repository2:[1111 2222]]

// func getRepoList(db* mongoDb) map[string][]string{
// 	repository := make (map[string][]string)
// 	rows, err := db.Query()
// 	if err != nil {
// 		log.Panic(err)
// 		return
// 	}
// 	defer row.Close()
// 	for rows.Next(){
// 		var id_URLWebhook string
// 		var list string
// 		err = rows.Scan(&id_URLWebhook,&list)
// 		if err != nil {
// 			log.Panic(err)
// 			return 
// 		}
// 			repository[id_URLWebhook] = strings.Split(list,",")
// 		--
// 		err = row.Err()
// 		if err != nil {
// 			log.Panic(err)
// 			return 
// 		}
// 		return repository
// }
// add one new document into database
func addNewUrlAndIDchat(db* mongoDb,URLWebhook string, Idchat string){
	doc := bson.D{{"ID_chat",id_chat},{"URL_Webhook",URLWebhook}}

	result, err := db.InsertOne(context.TODO(), doc)
	if err != nil{
		panic(err)
	}
}

// update new URLwebhook after generate a new link
func updateUrl(db* mongoDb,URLWebhook string, Idchat string){
	filter := bson.D{{"id_chat",idchat}}
	replacement := bson.D{{"ID_chat",idchat},{"URL_Webhook",URLWebhook}}

	result, err := db.ReplaceOne(context.TODO(),filter, replacement)
	if err != nil {
		panic(err)
	}
}

// remove URLwebhook after stop command
func remove(db* mongoDb,URLWebhook string){
	filter := bson.D{{"URL_Webhook",URLWebhook}}

	result, err := db.DeleteOne(context.TODO(),filter)
	if err != nil {
		panic(err)
	}
}
// check available Idchat
func checkIDchatAndURL(db* mongoDB,Idchat string ) bool{
	var result bson.M
	err := db.FindOne(context.TODO(), bson.D{{"ID_chat",Idchat}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocument {
			return false
		}
		return true
	}
	panic (err)
}