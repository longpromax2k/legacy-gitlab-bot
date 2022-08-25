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

	c "gitlabhook/controllers"
	h "gitlabhook/helpers"
	r "gitlabhook/routes"

	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	// Handle request and endpoints
	r.HandleRoute()
	// server config
	srv := &http.Server{
		Addr:        ":" + h.Port,
		Handler:     r.R,
		ReadTimeout: 10 * time.Second,
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

		if c.CheckUpOid.Hex() != "000000000000000000000000" {
			f := bson.D{{Key: "_id", Value: c.CheckUpOid}}
			if _, err := h.CheckUpCol.DeleteOne(context.TODO(), f); err != nil {
				log.Panic(err)
			}

		}
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("http server shutdown error: %v", err)
		}
		if err := h.Db.Disconnect(ctx); err != nil {
			log.Printf("database shutdown error: %v", err)
		}

		log.Printf("shutdown completed\n")
		close(idleConnsClosed)
	}()

	// Handle Telegram Command
	go c.HandleCommand()
	// Serve
	go func() {
		log.Printf("Listening to port %s.\n", h.Port)
		if err := srv.ListenAndServe(); err != nil {
			if err.Error() != "http: Server closed" {
				log.Printf("HTTP server closed with: %v\n", err)
			}
		}

	}()
	wg.Wait()
}
