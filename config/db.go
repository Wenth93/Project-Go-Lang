package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DbConfig struct {
	Url string `envconfig:"DB_URL"`
}

func connectToDb(dbConfig *DbConfig) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(dbConfig.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		fmt.Println("Connection to db established.")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(fmt.Errorf("connection to db failed: %v", err))
	}

	return pool
}
