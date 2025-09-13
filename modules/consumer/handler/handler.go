package handler

import (
	"net/http"
	"strconv"
	"kafka-marketplace/modules/consumer/model"
	"kafka-marketplace/modules/consumer/service"
	"github.com/labstack/echo/v4"
)

type ConsumerHandler struct {
	service service.ConsumerService
}

type ConsumerHandlerInterface interface {
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

func NewConsumerHandler(service service.ConsumerService) ConsumerHandlerInterface {
	return &ConsumerHandler{service}
}

func (h *ConsumerHandler) GetAll(c echo.Context) error {
	consumers, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, consumers)
}
func (h *ConsumerHandler) GetByID(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid consumer ID"})
		}
		consumer, err := h.service.GetByID(id)
		if err != nil {
			if err.Error() == "consumer not found" {
				return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, consumer)
	
}
func (h *ConsumerHandler) Create(c echo.Context) error {
	var consumer model.Consumer
	if err := c.Bind(&consumer); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

		createdConsumer, err := h.service.Create(consumer)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, createdConsumer)
}
func (h *ConsumerHandler) Update(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid consumer ID"})
		}
		var consumer model.Consumer
		if err := c.Bind(&consumer); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}
		consumer.ID = id
		updatedConsumer, err := h.service.Update(consumer)
		if err != nil {
			if err.Error() == "consumer not found" {
				return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, updatedConsumer)
}
func (h *ConsumerHandler) Delete(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid consumer ID"})
		}
		err = h.service.Delete(id)
		if err != nil {
			if err.Error() == "consumer not found" {
				return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.NoContent(http.StatusNoContent)
}