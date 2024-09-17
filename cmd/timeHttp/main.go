package main

import (
	"log"

	timehttp "github.com/codescalersinternships/Datetime-Server-Marwan-Radwan/pkg/timeHttp"
)

func main() {
	PORT := "8080"

	log.Printf("Http Server is running on %s", PORT)

	err := timehttp.StartServer(PORT)
	if err != nil {
		log.Fatalf("error starting the http server \n%s", err)
	}
}
