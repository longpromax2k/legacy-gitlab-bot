package main

import (
	"log"
	"net/http"
	"os"

	c "github.com/tatsuxyz/GitLabHook/controllers"
	r "github.com/tatsuxyz/GitLabHook/routes"
)

func main() {
	// Handle request and endpoints
	r.HandleRoute()

	// Disconnect database at the end of the program
	defer c.Db.Close()

	// Serve
	go func() {
		port := os.Getenv("PORT")
		log.Printf("Listening to port %s.\n", port)
		http.ListenAndServe(":"+port, r.R)
	}()
	c.HandleCommand()
}
