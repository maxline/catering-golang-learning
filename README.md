# Catering Service API

Простое приложение для кейтеринга на Go с использованием Echo framework.

## Возможности

- Просмотр меню и категорий блюд
- Создание и управление заказами
- Мок-оплата заказов
- Отслеживание статуса заказов

## Запуск приложения

1. Установите зависимости:
```bash
go mod tidy
```

2. Запустите сервер:
```bash
go run .
```

Сервер запустится на порту `:8080`

## API Endpoints

### Меню

#### GET /api/menu
Получить полное меню
```json
{
  "success": true,
  "data": [
    {
      "id": "1",
      "name": "Брускетта с томатами",
      "description": "Итальянская закуска с помидорами и базиликом",
      "price": 350.0,
      "category": "1",
      "available": true,
      "image_url": "https://example.com/bruschetta.jpg"
    }
  ]
}
```

#### GET /api/menu/categories
Получить категории блюд
```json
{
  "success": true,
  "data": [
    {
      "id": "1",
      "name": "Закуски"
    }
  ]
}
```

### Заказы

#### POST /api/orders
Создать новый заказ
```json
{
  "customer_name": "Иван Иванов",
  "customer_phone": "+7-999-123-45-67",
  "items": [
    {
      "menu_item_id": "1",
      "name": "Брускетта с томатами",
      "quantity": 2,
      "price": 350.0,
      "total": 700.0
    }
  ]
}
```

#### GET /api/orders/:id
Получить заказ по ID
```json
{
  "success": true,
  "data": {
    "id": "uuid-here",
    "customer_name": "Иван Иванов",
    "customer_phone": "+7-999-123-45-67",
    "items": [...],
    "total_amount": 700.0,
    "status": "pending",
    "created_at": "2024-01-01T12:00:00Z",
    "updated_at": "2024-01-01T12:00:00Z"
  }
}
```

#### GET /api/orders
Получить все заказы
```json
{
  "success": true,
  "data": [...]
}
```

#### PUT /api/orders/:id/status
Обновить статус заказа
```json
{
  "status": "confirmed"
}
```

Доступные статусы:
- `pending` - ожидает подтверждения
- `confirmed` - подтвержден
- `preparing` - готовится
- `ready` - готов
- `delivered` - доставлен
- `cancelled` - отменен

### Платежи

#### POST /api/payments
Обработать платеж
```json
{
  "order_id": "uuid-here",
  "method": "card"
}
```

Доступные методы оплаты:
- `card` - карта
- `cash` - наличные
- `online` - онлайн

#### GET /api/payments/:id
Получить статус платежа
```json
{
  "success": true,
  "data": {
    "id": "payment-uuid",
    "order_id": "order-uuid",
    "amount": 700.0,
    "status": "success",
    "method": "card",
    "created_at": "2024-01-01T12:00:00Z"
  }
}
```

## Тестирование API

### Пример создания заказа и оплаты

1. Создайте заказ:
```bash
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer_name": "Тест Клиент",
    "customer_phone": "+7-999-123-45-67",
    "items": [
      {
        "menu_item_id": "1",
        "name": "Брускетта с томатами",
        "quantity": 1,
        "price": 350.0,
        "total": 350.0
      },
      {
        "menu_item_id": "3",
        "name": "Стейк Рибай",
        "quantity": 1,
        "price": 1200.0,
        "total": 1200.0
      }
    ]
  }'
```

2. Оплатите заказ (используйте order_id из предыдущего ответа):
```bash
curl -X POST http://localhost:8080/api/payments \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": "ORDER_ID_FROM_PREVIOUS_RESPONSE",
    "method": "card"
  }'
```

3. Проверьте статус заказа:
```bash
curl http://localhost:8080/api/orders/ORDER_ID_FROM_PREVIOUS_RESPONSE
```

## Структура проекта

```
catering-service/
├── main.go              # Основной файл приложения
├── models.go            # Модели данных
├── menu_controller.go   # Контроллер меню
├── order_controller.go  # Контроллер заказов
├── payment_controller.go # Контроллер платежей
├── go.mod              # Зависимости Go
└── README.md           # Документация
```

## Особенности

- Все данные хранятся в памяти (при перезапуске сервера данные теряются)
- Платежи обрабатываются мок-логикой (90% успешных, 10% неудачных)
- Валидация входных данных
- CORS поддержка для фронтенда
- Логирование запросов
