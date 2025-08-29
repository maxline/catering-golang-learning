package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	payments map[string]*Payment
	orders   map[string]*Order // Reference to orders for getting amount
}

func NewPaymentController() *PaymentController {
	return &PaymentController{
		payments: make(map[string]*Payment),
		orders:   make(map[string]*Order),
	}
}

// SetOrderReference sets reference to orders
func (pc *PaymentController) SetOrderReference(orders map[string]*Order) {
	pc.orders = orders
}

// ProcessPayment processes payment (mock implementation)
func (pc *PaymentController) ProcessPayment(c echo.Context) error {
	var req PaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid data format",
		})
	}

	// Validation
	if req.OrderID == "" {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Order ID is required",
		})
	}

	if req.Method == "" {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Payment method is required",
		})
	}

	// Check if order exists
	order, exists := pc.orders[req.OrderID]
	if !exists {
		return c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Order not found",
		})
	}

	// Payment method validation
	validMethods := []string{PaymentMethodCard, PaymentMethodCash, PaymentMethodOnline}
	isValidMethod := false
	for _, method := range validMethods {
		if req.Method == method {
			isValidMethod = true
			break
		}
	}

	if !isValidMethod {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid payment method",
		})
	}

	// Create payment
	payment := NewPayment(req.OrderID, order.TotalAmount, req.Method)
	pc.payments[payment.ID] = payment

	// Mock payment processing (in real application there would be payment system integration)
	// Simulate successful payment in 90% of cases
	success := true
	if time.Now().UnixNano()%10 == 0 { // 10% failure probability
		success = false
	}

	if success {
		payment.Status = PaymentStatusSuccess
		// Update order status to "confirmed"
		order.Status = OrderStatusConfirmed
		order.UpdatedAt = time.Now()
	} else {
		payment.Status = PaymentStatusFailed
	}

	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Payment processed",
		Data: map[string]interface{}{
			"payment": payment,
			"success": success,
		},
	})
}

// GetPaymentStatus returns payment status
func (pc *PaymentController) GetPaymentStatus(c echo.Context) error {
	paymentID := c.Param("id")

	payment, exists := pc.payments[paymentID]
	if !exists {
		return c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Payment not found",
		})
	}

	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    payment,
	})
}
