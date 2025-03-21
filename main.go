package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	c "github.com/tatsuxyz/GitLabHook/controllers"
	r "github.com/tatsuxyz/GitLabHook/routes"
)

func main() {
	port := os.Getenv("PORT")
	var wg sync.WaitGroup
	wg.Add(1)

	// Handle request and endpoints
	r.HandleRoute()

	// server config
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r.R,
	}

	// Listening to interrupt signal
	idleConnsClosed := make(chan struct{})
	go func() {
		defer wg.Done()
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		log.Printf("service interrupt received\n")
		log.Printf("http server shutting down\n")

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("http server shutdown error: %v", err)
		}
		if err := c.Db.Close(); err != nil {
			log.Printf("database shutdown error: %v", err)
		}
		close(idleConnsClosed)
	}()

	// Handle Telegram Command
	c.HandleCommand()

	// Serve
	log.Printf("Listening to port %s.\n", port)
	if err := srv.ListenAndServe(); err != nil {
		if err.Error() != "http: Server closed" {
			log.Printf("HTTP server closed with: %v\n", err)
		}
	}

	wg.Wait()
}
