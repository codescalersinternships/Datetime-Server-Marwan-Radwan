package main

import (
	"log"

	timegin "github.com/codescalersinternships/Datetime-Server-Marwan-Radwan/pkg/timeGin"
)

func main() {
	PORT := ":8081"

	e := timegin.StartServer()
	if err := e.Run(PORT); err != nil {
		log.Printf("can't run server correctly %v\n", err)
	}

}
