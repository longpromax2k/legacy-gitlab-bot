package configs

import (
	"gitbot/models"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	lock        = &sync.Mutex{}
	env         *models.Config
	checkStatus string
)

func GetConfig() *models.Config {
	if env == nil {
		lock.Lock()
		defer lock.Unlock()

		if env == nil {
			if err := godotenv.Load(); err != nil {
				log.Printf(".env file not found.\n")
			}

			env = &models.Config{
				Port:     os.Getenv("PORT"),
				HostURL:  os.Getenv("HOST_URL"),
				PathURL:  os.Getenv("URL_PATH"),
				BotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
				MongoURI: os.Getenv("MONGO_URI"),
			}
		}
	}

	return env
}

func GetCheckStatus() string {
	return checkStatus
}
func SetCheckStatus(id string) {
	checkStatus = id
}
