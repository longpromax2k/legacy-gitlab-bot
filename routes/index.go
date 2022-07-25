package routes

import (
	"net/http"

	"github.com/tatsuxyz/GitLabHook/controllers"
)

func Routes() {
	http.HandleFunc("/webhook", controllers.HandleWebHook)
}
