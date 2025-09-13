package database

import (
	"fmt"

	consumerModel "kafka-marketplace/modules/consumer/model"
	paymentModel "kafka-marketplace/modules/payment/model"
	productModel "kafka-marketplace/modules/product/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionPostgres() *gorm.DB {
	cfg, err := loadPostgreSQLConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load PostgreSQL config: %v", err))
	}
	fmt.Println(cfg)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.POSTGRESQL_HOST,
		cfg.POSTGRESQL_USER,
		cfg.POSTGRESQL_PASS,
		cfg.POSTGRESQL_NAME,
		cfg.POSTGRESQL_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect PostgreSQL: %v", err))
	}

	db.AutoMigrate(
		consumerModel.Consumer{},
		productModel.Product{},
		paymentModel.Payment{},
	)

	fmt.Println("Connected to PostgreSQL with GORM successfully")
	return db
}