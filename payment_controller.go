package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type PaymentController struct {
	payments map[string]*Payment
	orders   map[string]*Order // Ссылка на заказы для получения суммы
}

func NewPaymentController() *PaymentController {
	return &PaymentController{
		payments: make(map[string]*Payment),
		orders:   make(map[string]*Order),
	}
}

// SetOrderReference устанавливает ссылку на заказы
func (pc *PaymentController) SetOrderReference(orders map[string]*Order) {
	pc.orders = orders
}

// ProcessPayment обрабатывает платеж (мок-реализация)
func (pc *PaymentController) ProcessPayment(c echo.Context) error {
	var req PaymentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Неверный формат данных",
		})
	}

	// Валидация
	if req.OrderID == "" {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "ID заказа обязателен",
		})
	}

	if req.Method == "" {
		return c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Метод оплаты обязателен",
		})
	}

	// Проверяем существование заказа
	order, exists := pc.orders[req.OrderID]
	if !exists {
		return c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Заказ не найден",
		})
	}

	// Валидация метода оплаты
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
			Message: "Неверный метод оплаты",
		})
	}

	// Создаем платеж
	payment := NewPayment(req.OrderID, order.TotalAmount, req.Method)
	pc.payments[payment.ID] = payment

	// Мок-обработка платежа (в реальном приложении здесь была бы интеграция с платежной системой)
	// Имитируем успешную оплату в 90% случаев
	success := true
	if time.Now().UnixNano()%10 == 0 { // 10% вероятность неудачи
		success = false
	}

	if success {
		payment.Status = PaymentStatusSuccess
		// Обновляем статус заказа на "подтвержден"
		order.Status = OrderStatusConfirmed
		order.UpdatedAt = time.Now()
	} else {
		payment.Status = PaymentStatusFailed
	}

	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Платеж обработан",
		Data: map[string]interface{}{
			"payment": payment,
			"success": success,
		},
	})
}

// GetPaymentStatus возвращает статус платежа
func (pc *PaymentController) GetPaymentStatus(c echo.Context) error {
	paymentID := c.Param("id")

	payment, exists := pc.payments[paymentID]
	if !exists {
		return c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Платеж не найден",
		})
	}

	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    payment,
	})
}
