package main

import (
	"fmt"
	"log"

	timegin "github.com/codescalersinternships/Datetime-Server-Marwan-Radwan/pkg/timeGin"
	"github.com/fvbock/endless"
)

func main() {
	PORT := ":8081"

	e := timegin.StartServer()
	if err := e.Run(PORT); err != nil {
		log.Printf("can't run server correctly %v\n", err)
	}

	err := endless.ListenAndServe(PORT, e)

	if err != nil {
		log.Fatalf("cannot shutdown server gracefully %v\n", err)
	}

	fmt.Println("\nshutting down HTTP server gracefully...")
}
