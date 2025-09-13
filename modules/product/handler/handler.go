package handler

import (
	"net/http"
	"strconv"
	"kafka-marketplace/modules/product/model"
	"kafka-marketplace/modules/product/service"
	"github.com/labstack/echo/v4"
)
type ProductHandler struct {
	service service.ProductService
}
type ProductHandlerInterface interface {
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
func NewProductHandler(service service.ProductService) ProductHandlerInterface {
	return &ProductHandler{service}
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	products, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}
func (h *ProductHandler) GetByID(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid product ID"})
		}
		product, err := h.service.GetByID(id)
		if err != nil {
			if err.Error() == "product not found" {
				return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, product)

}
func (h *ProductHandler) Create(c echo.Context) error {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
		createdProduct, err := h.service.Create(product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, createdProduct)
}
func (h *ProductHandler) Update(c echo.Context) error {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid product ID"})
		}
		product.ID = id
		updatedProduct, err := h.service.Update(product)
		if err != nil {
			if err.Error() == "product not found" {
				return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, updatedProduct)
}
func (h *ProductHandler) Delete(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid product ID"})
		}
		err = h.service.Delete(id)
		if err != nil {
			if err.Error() == "product not found" {
				return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.NoContent(http.StatusNoContent)
}