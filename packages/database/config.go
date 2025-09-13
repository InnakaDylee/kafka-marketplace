package database

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

type PostgreSQLConfig struct {
		POSTGRESQL_USER string
		POSTGRESQL_PASS string
		POSTGRESQL_HOST string
		POSTGRESQL_PORT string
		POSTGRESQL_NAME string
}

func loadPostgreSQLConfig() (*PostgreSQLConfig, error) {
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("error loading .env file: %v", err)
		}
	} else {
		fmt.Println(".env file not found, using default values")
	}
	fmt.Println(os.Getenv("POSTGRES_USER"))
	return &PostgreSQLConfig{
			POSTGRESQL_USER: os.Getenv("POSTGRES_USER"),
			POSTGRESQL_PASS: os.Getenv("POSTGRES_PASS"),
			POSTGRESQL_HOST: os.Getenv("POSTGRES_HOST"),
			POSTGRESQL_PORT: os.Getenv("POSTGRES_PORT"),
			POSTGRESQL_NAME: os.Getenv("POSTGRES_NAME"),
	}, nil
}