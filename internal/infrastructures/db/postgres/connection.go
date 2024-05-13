package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/danzBraham/eniqilo-store/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(env *config.ConfigEnv) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s",
		env.DBUsername,
		env.DBPassword,
		env.DBHost,
		env.DBPort,
		env.DBName,
		env.DBParams,
	)

	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	config.MinConns = 10
	config.MaxConns = 20
	config.MaxConnIdleTime = 10 * time.Minute
	config.MaxConnLifetime = 60 * time.Minute

	ctx := context.Background()
	dbpool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(ctx); err != nil {
		dbpool.Close()
		return nil, err
	}

	log.Println("Connected to the database")
	return dbpool, nil
}
