package main

import (
	"fmt"
	"kafka-marketplace/packages/database"
	consumerRouter "kafka-marketplace/modules/consumer/router"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Consumer service main
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			panic("failed to load env file")
		}
	}

	e := echo.New()

	db := database.ConnectionPostgres()

	consumerRouter.SetupRoutes(e, db)

	if err := e.Start(":8001"); err != nil {
		panic(fmt.Sprintf("error starting server: %v", err))
	}
}