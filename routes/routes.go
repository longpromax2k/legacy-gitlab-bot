package routes

import (
	c "gitlabhook/controllers"
	h "gitlabhook/helpers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	R = chi.NewRouter()
)

func HandleRoute() {
	R.Use(middleware.Logger)

	R.Post("/"+h.UrlPath+"/{id}", c.HandleHook)
}
