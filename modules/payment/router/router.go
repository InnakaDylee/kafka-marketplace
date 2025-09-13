package router

import (
	"kafka-marketplace/modules/payment/handler"
	"kafka-marketplace/modules/payment/repository"
	"kafka-marketplace/modules/payment/service"
	"kafka-marketplace/packages/queue"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB, kafkaCfg *queue.KafkaConfig) {
	// Setup your routes here, for example:
	payment := e.Group("/api/v1/payments")
	payment.GET("/test", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "List of payments",
		})
	})

	paymentRepository := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepository, kafkaCfg)
	handlerPayment := handler.NewPaymentHandler(paymentService)

	payment.GET("", handlerPayment.GetAll)
	payment.GET("/:id", handlerPayment.GetByID)
	payment.POST("", handlerPayment.Create)
	payment.PUT("/:id", handlerPayment.Update)
	payment.DELETE("/:id", handlerPayment.Delete)
}