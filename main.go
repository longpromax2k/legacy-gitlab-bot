package main

import (
	"context"
	"log"
	"net/http"
	"os"

	c "github.com/tatsuxyz/GitLabHook/controllers"
	h "github.com/tatsuxyz/GitLabHook/helpers"
	r "github.com/tatsuxyz/GitLabHook/routes"
)

func main() {
	// Handle request and endpoints
	r.HandleRoute()

	// Disconnect MongoDB at the end of the program
	defer func() {
		if err := h.Client.Disconnect(context.TODO()); err != nil {
			log.Panic(err)
			return
		}
	}()

	// Serve
	go func() {
		port := os.Getenv("PORT")
		log.Printf("Listening to port %s.\n", port)
		http.ListenAndServe("localhost:"+port, r.R)
	}()
	c.HandleCommand()
}
