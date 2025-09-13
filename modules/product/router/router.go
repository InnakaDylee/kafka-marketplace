package router

import (
	"kafka-marketplace/modules/product/handler"
	"kafka-marketplace/modules/product/repository"
	"kafka-marketplace/modules/product/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(e *echo.Echo, db *gorm.DB) {
	// Setup your routes here, for example:
	product := e.Group("/api/v1/products")

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	handlerProduct := handler.NewProductHandler(productService)

	product.GET("/", handlerProduct.GetAll)
	product.GET("/:id", handlerProduct.GetByID)
	product.POST("/", handlerProduct.Create)
	product.PUT("/:id", handlerProduct.Update)
	product.DELETE("/:id", handlerProduct.Delete)
}