package router

import (
	"kafka-marketplace/modules/consumer/handler"
	"kafka-marketplace/modules/consumer/repository"
	"kafka-marketplace/modules/consumer/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// Setup your routes here, for example:
	consumer := e.Group("/api/v1/consumers")
	consumer.GET("/test", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "List of consumers",
		})
	})

	consumerRepository := repository.NewConsumerRepository(db)
	consumerService := service.NewConsumerService(consumerRepository)
	consumerHandler := handler.NewConsumerHandler(consumerService)

	consumer.GET("", consumerHandler.GetAll)
	consumer.POST("", consumerHandler.Create)
	consumer.GET("/:id", consumerHandler.GetByID)
	consumer.PUT("/:id", consumerHandler.Update)
	consumer.DELETE("/:id", consumerHandler.Delete)
}