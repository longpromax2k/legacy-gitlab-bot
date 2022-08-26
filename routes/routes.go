package routes

import (
	c "gitbot/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	R = chi.NewRouter()
)

func HandleRoute() {
	R.Use(middleware.Logger)

	R.Post("/"+config.URL_PATHr+"/{id}", c.HandleHook)
}
