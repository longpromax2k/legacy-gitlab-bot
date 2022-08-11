package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	c "github.com/tatsuxyz/GitLabHook/controllers"
	r "github.com/tatsuxyz/GitLabHook/routes"
)

// Our custom handler that holds a wait group used to block the shutdown
// while it's running the jobs.
type CustomHandler struct {
	wg *sync.WaitGroup
}

func NewCustomerHandler(wg *sync.WaitGroup) *CustomHandler {
	return &CustomHandler{wg: wg}
}
func main() {

	wg := &sync.WaitGroup{}
	//customHandler :=NewCustomerHandler(wg)

	// Handle request and endpoints
	r.HandleRoute()

	port := os.Getenv("PORT")
	log.Printf("Listening to port %s.\n", port)

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: r.R,
	}
	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-termChan // Blocks here until interrupted
		log.Print("SIGTERM received. Shutdown process initiated\n")
		httpServer.Shutdown(context.Background())
	}()

	// Disconnect database at the end of the program
	defer c.Db.Close()

	// Serve
	// port := os.Getenv("PORT")
	// log.Printf("Listening to port %s.\n", port)
	if err := httpServer.ListenAndServe(); err != nil {
		if err.Error() != "http: Server closed" {
			log.Printf("HTTP server closed with: %v\n", err)
		}
		log.Printf("HTTP server shut down")
		//os.Exit(0)
	}
	go func() {

		http.ListenAndServe(":"+port, r.R)

	}()

	// This is where, once we're closing the program, we wait for all
	// jobs (they all have been added to this WaitGroup) to `wg.Done()`.
	log.Println("waiting for running jobs to finish")
	wg.Wait()
	c.HandleCommand()
	log.Println("jobs finished. exiting")

}
