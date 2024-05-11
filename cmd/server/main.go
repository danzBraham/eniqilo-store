package main

import (
	"log"

	"github.com/danzBraham/eniqilo-store/internal/http/api"
)

func main() {
	server := api.NewAPIServer("localhost:8080", nil)
	if err := server.Run(); err != nil {
		log.Panic(err)
	}
}
