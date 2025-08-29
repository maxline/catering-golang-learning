package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create new Echo instance
	e := echo.New()

	// Add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize controllers
	menuController := NewMenuController()
	orderController := NewOrderController()
	paymentController := NewPaymentController()

	// Link controllers (for access to orders in payments)
	paymentController.SetOrderReference(orderController.orders)

	// Menu routes
	e.GET("/api/menu", menuController.GetMenu)
	e.GET("/api/menu/categories", menuController.GetCategories)

	// Order routes
	e.POST("/api/orders", orderController.CreateOrder)
	e.GET("/api/orders/:id", orderController.GetOrder)
	e.GET("/api/orders", orderController.GetAllOrders)
	e.PUT("/api/orders/:id/status", orderController.UpdateOrderStatus)

	// Payment routes
	e.POST("/api/payments", paymentController.ProcessPayment)
	e.GET("/api/payments/:id", paymentController.GetPaymentStatus)

	// Start server
	log.Println("Catering service starting on :8080")
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting server:", err)
	}
}
