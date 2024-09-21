package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	timehttp "github.com/codescalersinternships/Datetime-Server-Marwan-Radwan/pkg/timeHttp"
)

func main() {
	PORT := "8080"

	log.Printf("Http Server is running on %s", PORT)

	err := timehttp.StartServer(PORT)
	if err != nil {
		log.Fatalf("error starting the http server \n%s", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := timehttp.ShutdownServer(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
