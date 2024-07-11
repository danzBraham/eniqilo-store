package main

import (
	"log"

	"github.com/danzBraham/eniqilo-store/internal/http"
)

func main() {
	addr := ":8080"

	server := http.NewServer(addr, nil)
	if err := server.Launch(); err != nil {
		log.Fatal(err)
	}
}
