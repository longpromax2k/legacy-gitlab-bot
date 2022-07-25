package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tatsuxyz/GitLabHook/routes"
)

func main() {
	// Load Environment Variable
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		fmt.Println("[GitLabHook] Failed to load environment variable")
	}

	// Handle request and endpoints
	routes.Routes()

	// Serve
	fmt.Println("[GitLabHook] Listening to port " + os.Getenv("PORT"))
	http.ListenAndServe("localhost:"+os.Getenv("PORT"), nil)
}
