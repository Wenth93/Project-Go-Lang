package config

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DbConn *pgxpool.Pool
}

func Load() *AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("couldn't load env file")
	}

	dbConfig := DbConfig{}
	envconfig.MustProcess("DB", &dbConfig)

	conn := connectToDb(&dbConfig)

	return &AppConfig{
		DbConn: conn,
	}
}
