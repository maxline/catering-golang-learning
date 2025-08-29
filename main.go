package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Создаем новый экземпляр Echo
	e := echo.New()

	// Добавляем middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Инициализируем контроллеры
	menuController := NewMenuController()
	orderController := NewOrderController()
	paymentController := NewPaymentController()

	// Связываем контроллеры (для доступа к заказам в платежах)
	paymentController.SetOrderReference(orderController.orders)

	// Роуты для меню
	e.GET("/api/menu", menuController.GetMenu)
	e.GET("/api/menu/categories", menuController.GetCategories)

	// Роуты для заказов
	e.POST("/api/orders", orderController.CreateOrder)
	e.GET("/api/orders/:id", orderController.GetOrder)
	e.GET("/api/orders", orderController.GetAllOrders)
	e.PUT("/api/orders/:id/status", orderController.UpdateOrderStatus)

	// Роуты для оплаты
	e.POST("/api/payments", paymentController.ProcessPayment)
	e.GET("/api/payments/:id", paymentController.GetPaymentStatus)

	// Запускаем сервер
	log.Println("Catering service starting on :8080")
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal("Error starting server:", err)
	}
}
