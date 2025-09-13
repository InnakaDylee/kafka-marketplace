package handler

import (
	"net/http"
	"strconv"
	"kafka-marketplace/modules/payment/model"
	"kafka-marketplace/modules/payment/service"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	service service.PaymentService
}
type PaymentHandlerInterface interface {
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func NewPaymentHandler(service service.PaymentService) PaymentHandlerInterface {
	return &PaymentHandler{service}
}

func (h *PaymentHandler) GetAll(c echo.Context) error {
	payments, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, payments)
}
func (h *PaymentHandler) GetByID(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payment ID"})
		}
		payment, err := h.service.GetByID(id)
		if err != nil {
			if err.Error() == "payment not found" {
				return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, payment)
}
func (h *PaymentHandler) Create(c echo.Context) error {
	var payment model.Payment
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
		createdPayment, err := h.service.Create(payment)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, createdPayment)
}
func (h *PaymentHandler) Update(c echo.Context) error {
	var payment model.Payment
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payment ID"})
		}
		payment.ID = id
		updatedPayment, err := h.service.Update(payment)
		if err != nil {
			if err.Error() == "payment not found" {
				return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, updatedPayment)
}
func (h *PaymentHandler) Delete(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid payment ID"})
		}
		err = h.service.Delete(id)
		if err != nil {
			if err.Error() == "payment not found" {
				return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.NoContent(http.StatusNoContent)
}