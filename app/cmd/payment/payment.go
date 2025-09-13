package main

import (
	"fmt"
	paymentRouter "kafka-marketplace/modules/payment/router"
	"kafka-marketplace/packages/database"
	"kafka-marketplace/packages/queue"
	"os"

	"context"
	consumerRepository "kafka-marketplace/modules/consumer/repository"
	consumerService "kafka-marketplace/modules/consumer/service"
	productRepository "kafka-marketplace/modules/product/repository"
	productService "kafka-marketplace/modules/product/service"
	"log"
	"os/signal"
	"syscall"
	"time"


	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Payment service main
	_, err := os.Stat(".env")
	if err == nil {
		err := godotenv.Load()
		if err != nil {
			panic("failed to load env file")
		}
	}

	e := echo.New()

	db := database.ConnectionPostgres()

	kafkaCfg := &queue.KafkaConfig{
		Address:       []string{"localhost:9092"},
		Topic:         "payment-topic",
		GroupID:       "payment-consumer-group",
		Batch:         10,
		BatchTimeout:  100,
		RequiredAcks:  1,
	}
	
	paymentRouter.SetupRoutes(e, db, kafkaCfg)
	
	consumerRepository := consumerRepository.NewConsumerRepository(db)
	productRepository := productRepository.NewProductRepository(db)
	consumerService := consumerService.NewConsumerService(consumerRepository)
	productService := productService.NewProductService(productRepository)

	ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // 4. Run Kafka consumer di background
    go func() {
        log.Println("Kafka consumer started...")
        kafkaCfg.Read(ctx, consumerService, productService)
    }()

	go func() {
        if err := e.Start(":8002"); err != nil {
			panic(fmt.Sprintf("error starting server: %v", err))
		}
    }()

	// Wait for SIGINT / SIGTERM
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	<-sigchan
	log.Println("Shutdown signal received... closing Kafka")
	cancel()
	time.Sleep(2 * time.Second) // kasih waktu goroutine stop
}