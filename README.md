# Catering Service API

Simple catering application in Go using Echo framework.

## Features

- View menu and dish categories
- Create and manage orders
- Mock payment processing
- Track order status

## Running the Application

1. Install dependencies:
```bash
go mod tidy
```

2. Start the server:
```bash
go run .
```

The server will start on port `:8080`

## API Endpoints

### Menu

#### GET /api/menu
Get full menu
```json
{
  "success": true,
  "data": [
    {
      "id": "1",
      "name": "Bruschetta with Tomatoes",
      "description": "Italian appetizer with tomatoes and basil",
      "price": 350.0,
      "category": "1",
      "available": true,
      "image_url": "https://example.com/bruschetta.jpg"
    }
  ]
}
```

#### GET /api/menu/categories
Get dish categories
```json
{
  "success": true,
  "data": [
    {
      "id": "1",
      "name": "Appetizers"
    }
  ]
}
```

### Orders

#### POST /api/orders
Create new order
```json
{
  "customer_name": "John Doe",
  "customer_phone": "+7-999-123-45-67",
  "items": [
    {
      "menu_item_id": "1",
      "name": "Bruschetta with Tomatoes",
      "quantity": 2,
      "price": 350.0,
      "total": 700.0
    }
  ]
}
```

#### GET /api/orders/:id
Get order by ID
```json
{
  "success": true,
  "data": {
    "id": "uuid-here",
    "customer_name": "John Doe",
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
Get all orders
```json
{
  "success": true,
  "data": [...]
}
```

#### PUT /api/orders/:id/status
Update order status
```json
{
  "status": "confirmed"
}
```

Available statuses:
- `pending` - waiting for confirmation
- `confirmed` - confirmed
- `preparing` - being prepared
- `ready` - ready
- `delivered` - delivered
- `cancelled` - cancelled

### Payments

#### POST /api/payments
Process payment
```json
{
  "order_id": "uuid-here",
  "method": "card"
}
```

Available payment methods:
- `card` - card
- `cash` - cash
- `online` - online

#### GET /api/payments/:id
Get payment status
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

## API Testing

### Example of creating an order and payment

1. Create an order:
```bash
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "customer_name": "Test Customer",
    "customer_phone": "+7-999-123-45-67",
    "items": [
      {
        "menu_item_id": "1",
        "name": "Bruschetta with Tomatoes",
        "quantity": 1,
        "price": 350.0,
        "total": 350.0
      },
      {
        "menu_item_id": "3",
        "name": "Ribeye Steak",
        "quantity": 1,
        "price": 1200.0,
        "total": 1200.0
      }
    ]
  }'
```

2. Pay for the order (use order_id from previous response):
```bash
curl -X POST http://localhost:8080/api/payments \
  -H "Content-Type: application/json" \
  -d '{
    "order_id": "ORDER_ID_FROM_PREVIOUS_RESPONSE",
    "method": "card"
  }'
```

3. Check order status:
```bash
curl http://localhost:8080/api/orders/ORDER_ID_FROM_PREVIOUS_RESPONSE
```

## Project Structure

```
catering-service/
├── main.go              # Main application file
├── models.go            # Data models
├── menu_controller.go   # Menu controller
├── order_controller.go  # Order controller
├── payment_controller.go # Payment controller
├── go.mod              # Go dependencies
└── README.md           # Documentation
```

## Features

- All data is stored in memory (data is lost when server restarts)
- Payments are processed with mock logic (90% successful, 10% failed)
- Input data validation
- CORS support for frontend
- Request logging
