package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orders map[string]*Order
}

func NewOrderController() *OrderController {
	return &OrderController{
		orders: make(map[string]*Order),
	}
}

// CreateOrder создает новый заказ
func (oc *OrderController) CreateOrder(c echo.Context) error {
	var req CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Неверный формат данных",
		})
	}

	// Валидация
	if req.CustomerName == "" {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Имя клиента обязательно",
		})
	}

	if req.CustomerPhone == "" {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Телефон клиента обязателен",
		})
	}

	if len(req.Items) == 0 {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Заказ должен содержать хотя бы одну позицию",
		})
	}

	// Создаем заказ
	order := NewOrder(req.CustomerName, req.CustomerPhone, req.Items)
	oc.orders[order.ID] = order

	return c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: "Заказ успешно создан",
		Data:    order,
	})
}

// GetOrder возвращает заказ по ID
func (oc *OrderController) GetOrder(c echo.Context) error {
	orderID := c.Param("id")

	order, exists := oc.orders[orderID]
	if !exists {
		return c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Заказ не найден",
		})
	}

	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    order,
	})
}

// GetAllOrders возвращает все заказы
func (oc *OrderController) GetAllOrders(c echo.Context) error {
	orders := make([]*Order, 0, len(oc.orders))
	for _, order := range oc.orders {
		orders = append(orders, order)
	}

	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    orders,
	})
}

// UpdateOrderStatus обновляет статус заказа
func (oc *OrderController) UpdateOrderStatus(c echo.Context) error {
	orderID := c.Param("id")

	order, exists := oc.orders[orderID]
	if !exists {
		return c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Заказ не найден",
		})
	}

	var req UpdateOrderStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Неверный формат данных",
		})
	}

	// Валидация статуса
	validStatuses := []string{
		OrderStatusPending,
		OrderStatusConfirmed,
		OrderStatusPreparing,
		OrderStatusReady,
		OrderStatusDelivered,
		OrderStatusCancelled,
	}

	isValidStatus := false
	for _, status := range validStatuses {
		if req.Status == status {
			isValidStatus = true
			break
		}
	}

	if !isValidStatus {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Неверный статус заказа",
		})
	}

	// Обновляем статус
	order.Status = req.Status
	order.UpdatedAt = time.Now()

	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Статус заказа обновлен",
		Data:    order,
	})
}
