package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	c "github.com/tatsuxyz/GitLabHook/controllers"
)

var (
	R = chi.NewRouter()
)

func HandleRoute() {
	R.Use(middleware.Logger)

	R.Post("/webhook", c.HandleWebHook)
	R.Post("/wh/{id}", c.HandleHook)
}
