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

// CreateOrder creates a new order
func (oc *OrderController) CreateOrder(c echo.Context) error {
	var req CreateOrderRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid data format",
		})
	}

	// Validation
	if req.CustomerName == "" {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Customer name is required",
		})
	}

	if req.CustomerPhone == "" {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Customer phone is required",
		})
	}

	if len(req.Items) == 0 {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Order must contain at least one item",
		})
	}

	// Create order
	order := NewOrder(req.CustomerName, req.CustomerPhone, req.Items)
	oc.orders[order.ID] = order

	return c.JSON(http.StatusCreated, APIResponse{
		Success: true,
		Message: "Order created successfully",
		Data:    order,
	})
}

// GetOrder returns an order by ID
func (oc *OrderController) GetOrder(c echo.Context) error {
	orderID := c.Param("id")

	order, exists := oc.orders[orderID]
	if !exists {
		return c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Order not found",
		})
	}

	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    order,
	})
}

// GetAllOrders returns all orders
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

// UpdateOrderStatus updates order status
func (oc *OrderController) UpdateOrderStatus(c echo.Context) error {
	orderID := c.Param("id")

	order, exists := oc.orders[orderID]
	if !exists {
		return c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Order not found",
		})
	}

	var req UpdateOrderStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid data format",
		})
	}

	// Status validation
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
			Message: "Invalid order status",
		})
	}

	// Update status
	order.Status = req.Status
	order.UpdatedAt = time.Now()

	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Order status updated",
		Data:    order,
	})
}
