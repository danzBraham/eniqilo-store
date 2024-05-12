package main

import (
	"log"

	"github.com/danzBraham/eniqilo-store/config"
	"github.com/danzBraham/eniqilo-store/internal/db"
	"github.com/danzBraham/eniqilo-store/internal/http/api"
)

func main() {
	env := config.LoadEnv()

	dbpool, err := db.ConnectDB(env)
	if err != nil {
		log.Panic(err)
	}
	defer dbpool.Close()

	server := api.NewAPIServer("localhost:8080", dbpool)
	if err := server.Run(); err != nil {
		log.Panic(err)
	}
}
