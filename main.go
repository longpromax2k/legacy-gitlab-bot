package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"context"
	"os/signal"
	"syscall"

	c "github.com/tatsuxyz/GitLabHook/controllers"
	r "github.com/tatsuxyz/GitLabHook/routes"
)

func main() {

	timeWait := 15 * time.Second
	signChan := make(chan os.Signal, 1)
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

	// Set up a channel to hear system signal
	signal.Notify( signChan, os.Interrupt, syscall.SIGTERM) <-signChan
	lo.Printf("Shutting down")
	//Set up Timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer func(){
		log.Printf("Close another connection")
		cancel()
	}()
	log.Printf( "Stop http server")
	if err := http.Shutdown(ctx): err == context.DeadlineExceeded {
		log.Printf("Halted active connections")
	}
	close(signChan)
	log.Print("Completed")
}
