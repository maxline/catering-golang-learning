# Catering API Testing
Write-Host "=== Catering API Testing ===" -ForegroundColor Green

# Base URL
$baseUrl = "http://localhost:8080"

# 1. Get menu
Write-Host "`n1. Getting menu..." -ForegroundColor Yellow
try {
    $menuResponse = Invoke-RestMethod -Uri "$baseUrl/api/menu" -Method GET
    Write-Host "Menu retrieved successfully. Number of dishes: $($menuResponse.data.Count)" -ForegroundColor Green
    Write-Host "First dish: $($menuResponse.data[0].name) - $($menuResponse.data[0].price) rub." -ForegroundColor Cyan
} catch {
    Write-Host "Error getting menu: $($_.Exception.Message)" -ForegroundColor Red
}

# 2. Get categories
Write-Host "`n2. Getting categories..." -ForegroundColor Yellow
try {
    $categoriesResponse = Invoke-RestMethod -Uri "$baseUrl/api/menu/categories" -Method GET
    Write-Host "Categories retrieved successfully. Count: $($categoriesResponse.data.Count)" -ForegroundColor Green
    foreach ($category in $categoriesResponse.data) {
        Write-Host "  - $($category.name)" -ForegroundColor Cyan
    }
} catch {
    Write-Host "Error getting categories: $($_.Exception.Message)" -ForegroundColor Red
}

# 3. Create order
Write-Host "`n3. Creating order..." -ForegroundColor Yellow
$orderData = @{
    customer_name = "Test Customer"
    customer_phone = "+7-999-123-45-67"
    items = @(
        @{
            menu_item_id = "1"
            name = "Bruschetta with Tomatoes"
            quantity = 2
            price = 350.0
            total = 700.0
        },
        @{
            menu_item_id = "3"
            name = "Ribeye Steak"
            quantity = 1
            price = 1200.0
            total = 1200.0
        }
    )
} | ConvertTo-Json -Depth 3

try {
    $orderResponse = Invoke-RestMethod -Uri "$baseUrl/api/orders" -Method POST -Body $orderData -ContentType "application/json"
    $orderId = $orderResponse.data.id
    Write-Host "Order created successfully. ID: $orderId" -ForegroundColor Green
    Write-Host "Order total: $($orderResponse.data.total_amount) rub." -ForegroundColor Cyan
    Write-Host "Status: $($orderResponse.data.status)" -ForegroundColor Cyan
} catch {
    Write-Host "Error creating order: $($_.Exception.Message)" -ForegroundColor Red
    exit
}

# 4. Get created order
Write-Host "`n4. Getting order by ID..." -ForegroundColor Yellow
try {
    $getOrderResponse = Invoke-RestMethod -Uri "$baseUrl/api/orders/$orderId" -Method GET
    Write-Host "Order retrieved successfully" -ForegroundColor Green
    Write-Host "Customer: $($getOrderResponse.data.customer_name)" -ForegroundColor Cyan
    Write-Host "Number of items: $($getOrderResponse.data.items.Count)" -ForegroundColor Cyan
} catch {
    Write-Host "Error getting order: $($_.Exception.Message)" -ForegroundColor Red
}

# 5. Process payment
Write-Host "`n5. Processing payment..." -ForegroundColor Yellow
$paymentData = @{
    order_id = $orderId
    method = "card"
} | ConvertTo-Json

try {
    $paymentResponse = Invoke-RestMethod -Uri "$baseUrl/api/payments" -Method POST -Body $paymentData -ContentType "application/json"
    $paymentId = $paymentResponse.data.payment.id
    $paymentSuccess = $paymentResponse.data.success
    Write-Host "Payment processed. ID: $paymentId" -ForegroundColor Green
    Write-Host "Status: $($paymentResponse.data.payment.status)" -ForegroundColor Cyan
    Write-Host "Success: $paymentSuccess" -ForegroundColor Cyan
} catch {
    Write-Host "Error processing payment: $($_.Exception.Message)" -ForegroundColor Red
}

# 6. Get payment status
Write-Host "`n6. Getting payment status..." -ForegroundColor Yellow
try {
    $paymentStatusResponse = Invoke-RestMethod -Uri "$baseUrl/api/payments/$paymentId" -Method GET
    Write-Host "Payment status retrieved successfully" -ForegroundColor Green
    Write-Host "Status: $($paymentStatusResponse.data.status)" -ForegroundColor Cyan
    Write-Host "Method: $($paymentStatusResponse.data.method)" -ForegroundColor Cyan
} catch {
    Write-Host "Error getting payment status: $($_.Exception.Message)" -ForegroundColor Red
}

# 7. Update order status
Write-Host "`n7. Updating order status..." -ForegroundColor Yellow
$statusData = @{
    status = "confirmed"
} | ConvertTo-Json

try {
    $statusResponse = Invoke-RestMethod -Uri "$baseUrl/api/orders/$orderId/status" -Method PUT -Body $statusData -ContentType "application/json"
    Write-Host "Order status updated successfully" -ForegroundColor Green
    Write-Host "New status: $($statusResponse.data.status)" -ForegroundColor Cyan
} catch {
    Write-Host "Error updating order status: $($_.Exception.Message)" -ForegroundColor Red
}

# 8. Get all orders
Write-Host "`n8. Getting all orders..." -ForegroundColor Yellow
try {
    $allOrdersResponse = Invoke-RestMethod -Uri "$baseUrl/api/orders" -Method GET
    Write-Host "All orders retrieved successfully. Count: $($allOrdersResponse.data.Count)" -ForegroundColor Green
} catch {
    Write-Host "Error getting all orders: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n=== Testing completed ===" -ForegroundColor Green
