package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	c "github.com/tatsuxyz/GitLabHook/controllers"
	r "github.com/tatsuxyz/GitLabHook/routes"
)

func main() {
	port := os.Getenv("PORT")

	// Handle request and endpoints
	r.HandleRoute()
	// Disconnect database at the end of the program
	defer c.Db.Close()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r.R,
	}

	// Listening to interrupt signal
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		log.Println("service interrupt received")

		log.Println("http server shutting down")
		time.Sleep(5 * time.Second)

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("http server shutdown error: %v", err)
		}

		log.Println("shutdown complete")

		close(idleConnsClosed)
	}()

	// Serve
	go c.HandleCommand()
	log.Printf("Listening to port %s.\n", port)
	if err := srv.ListenAndServe(); err != nil {
		if err.Error() != "http: Server closed" {
			log.Printf("HTTP server closed with: %v\n", err)
		}
	}
}
