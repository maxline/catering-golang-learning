package main

import (
	"time"

	"github.com/google/uuid"
)

// MenuItem represents a dish in the menu
type MenuItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Available   bool    `json:"available"`
	ImageURL    string  `json:"image_url,omitempty"`
}

// Category represents a dish category
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	MenuItemID string  `json:"menu_item_id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Price      float64 `json:"price"`
	Total      float64 `json:"total"`
}

// Order represents an order
type Order struct {
	ID            string      `json:"id"`
	CustomerName  string      `json:"customer_name"`
	CustomerPhone string      `json:"customer_phone"`
	Items         []OrderItem `json:"items"`
	TotalAmount   float64     `json:"total_amount"`
	Status        string      `json:"status"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"`
}

// CreateOrderRequest request to create an order
type CreateOrderRequest struct {
	CustomerName  string      `json:"customer_name"`
	CustomerPhone string      `json:"customer_phone"`
	Items         []OrderItem `json:"items"`
}

// UpdateOrderStatusRequest request to update order status
type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}

// Payment represents a payment
type Payment struct {
	ID        string    `json:"id"`
	OrderID   string    `json:"order_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	Method    string    `json:"method"`
	CreatedAt time.Time `json:"created_at"`
}

// PaymentRequest payment request
type PaymentRequest struct {
	OrderID string `json:"order_id"`
	Method  string `json:"method"`
}

// APIResponse common API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Order statuses
const (
	OrderStatusPending   = "pending"
	OrderStatusConfirmed = "confirmed"
	OrderStatusPreparing = "preparing"
	OrderStatusReady     = "ready"
	OrderStatusDelivered = "delivered"
	OrderStatusCancelled = "cancelled"
)

// Payment statuses
const (
	PaymentStatusPending = "pending"
	PaymentStatusSuccess = "success"
	PaymentStatusFailed  = "failed"
)

// Payment methods
const (
	PaymentMethodCard   = "card"
	PaymentMethodCash   = "cash"
	PaymentMethodOnline = "online"
)

// NewOrder creates a new order
func NewOrder(customerName, customerPhone string, items []OrderItem) *Order {
	now := time.Now()
	total := 0.0
	for _, item := range items {
		total += item.Total
	}

	return &Order{
		ID:            uuid.New().String(),
		CustomerName:  customerName,
		CustomerPhone: customerPhone,
		Items:         items,
		TotalAmount:   total,
		Status:        OrderStatusPending,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
}

// NewPayment creates a new payment
func NewPayment(orderID string, amount float64, method string) *Payment {
	return &Payment{
		ID:        uuid.New().String(),
		OrderID:   orderID,
		Amount:    amount,
		Status:    PaymentStatusPending,
		Method:    method,
		CreatedAt: time.Now(),
	}
}
