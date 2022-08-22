package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	c "github.com/tatsuxyz/GitLabHook/controllers"
	h "github.com/tatsuxyz/GitLabHook/helpers"
)

var (
	R = chi.NewRouter()
)

func HandleRoute() {
	R.Use(middleware.Logger)

	R.Post("/"+h.UrlPath+"/{id}", c.HandleHook)
}
