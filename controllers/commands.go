package controllers

import (
	"strconv"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tatsuxyz/GitLabHook/model"
)

var chatId string

func CommandStart(up *tgbot.Update, msg *tgbot.MessageConfig) {
	chatId = strconv.Itoa(int(up.Message.Chat.ID))
	// hostUrl, urlPath := os.Getenv("HOST_URL"), os.Getenv("URL_PATH")

	// TODO: check existed data
	// res := q.FindInGroup(chatId)
	// if res.Token != "" {
	// 	log.Printf("Document existed.\n")
	// 	msg.Text = model.ChatExistMsg
	// 	return
	// }

	// TODO: insert new value into database
	// old code
	// h.Db.Update(func(tx *bbolt.Tx) error {
	// 	b := tx.Bucket([]byte("gitlabhook"))
	// 	v := b.Get([]byte(chatId))

	// 	if v != nil {
	// 		msg.Text = fmt.Sprintf(model.ChatExistMsg, hostUrl, urlPath, chatId, v)
	// 	} else {
	// 		uid := uuid.New()
	// 		err := b.Put([]byte(chatId), []byte(uid.String()))
	// 		if err != nil {
	// 			return err
	// 		}

	// 		msg.Text = fmt.Sprintf(model.ChatInsertMsg, hostUrl, urlPath, chatId, uid)
	// 	}

	// 	return nil
	// })
}

func CommandDrop(up *tgbot.Update, msg *tgbot.MessageConfig) {
	chatId = strconv.Itoa(int(up.Message.Chat.ID))

	// TODO
	// h.Db.Update(func(tx *bbolt.Tx) error {
	// 	b := tx.Bucket([]byte("gitlabhook"))
	// 	b.Delete([]byte(chatId))
	// 	return nil
	// })

	msg.Text = model.ChatDropMsg
}
