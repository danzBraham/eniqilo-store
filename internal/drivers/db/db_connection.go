package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

func GetConnectionPool() *pgxpool.Pool {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic("File.env not found: ", err)
	}

	dbUser := viper.GetString("DB_USERNAME")
	dbPass := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	dbParams := viper.GetString("DB_PARAMS")

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", dbUser, dbPass, dbHost, dbPort, dbName, dbParams)
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Panic(err)
	}

	config.MinConns = 10
	config.MaxConns = 20

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Panic(err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Panic(err)
		pool.Close()
	}

	return pool
}
