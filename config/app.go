package config

import (
	"fmt"
	"log"

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
	if err := envconfig.Process("DB", &dbConfig); err != nil {
		log.Fatal("error processing DB configuration:", err)
	}

	conn := connectToDb(&dbConfig)

	return &AppConfig{
		DbConn: conn,
	}
}
