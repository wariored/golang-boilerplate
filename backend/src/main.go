package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"wrapup/api"
	"wrapup/database"
)

func main() {
	// create a new MongoDB client
	client, err := database.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// create a new Echo instance and register routes
	e := api.Router(client)

	// create a new http.Server instance and set its handler to Echo
	server := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	// start the server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// gracefully shutdown the server when signaled to do so
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
