# Тестирование API кейтеринга
Write-Host "=== Тестирование API кейтеринга ===" -ForegroundColor Green

# Базовый URL
$baseUrl = "http://localhost:8080"

# 1. Получить меню
Write-Host "`n1. Получение меню..." -ForegroundColor Yellow
try {
    $menuResponse = Invoke-RestMethod -Uri "$baseUrl/api/menu" -Method GET
    Write-Host "Меню получено успешно. Количество блюд: $($menuResponse.data.Count)" -ForegroundColor Green
    Write-Host "Первое блюдо: $($menuResponse.data[0].name) - $($menuResponse.data[0].price) руб." -ForegroundColor Cyan
} catch {
    Write-Host "Ошибка при получении меню: $($_.Exception.Message)" -ForegroundColor Red
}

# 2. Получить категории
Write-Host "`n2. Получение категорий..." -ForegroundColor Yellow
try {
    $categoriesResponse = Invoke-RestMethod -Uri "$baseUrl/api/menu/categories" -Method GET
    Write-Host "Категории получены успешно. Количество: $($categoriesResponse.data.Count)" -ForegroundColor Green
    foreach ($category in $categoriesResponse.data) {
        Write-Host "  - $($category.name)" -ForegroundColor Cyan
    }
} catch {
    Write-Host "Ошибка при получении категорий: $($_.Exception.Message)" -ForegroundColor Red
}

# 3. Создать заказ
Write-Host "`n3. Создание заказа..." -ForegroundColor Yellow
$orderData = @{
    customer_name = "Тест Клиент"
    customer_phone = "+7-999-123-45-67"
    items = @(
        @{
            menu_item_id = "1"
            name = "Брускетта с томатами"
            quantity = 2
            price = 350.0
            total = 700.0
        },
        @{
            menu_item_id = "3"
            name = "Стейк Рибай"
            quantity = 1
            price = 1200.0
            total = 1200.0
        }
    )
} | ConvertTo-Json -Depth 3

try {
    $orderResponse = Invoke-RestMethod -Uri "$baseUrl/api/orders" -Method POST -Body $orderData -ContentType "application/json"
    $orderId = $orderResponse.data.id
    Write-Host "Заказ создан успешно. ID: $orderId" -ForegroundColor Green
    Write-Host "Сумма заказа: $($orderResponse.data.total_amount) руб." -ForegroundColor Cyan
    Write-Host "Статус: $($orderResponse.data.status)" -ForegroundColor Cyan
} catch {
    Write-Host "Ошибка при создании заказа: $($_.Exception.Message)" -ForegroundColor Red
    exit
}

# 4. Получить созданный заказ
Write-Host "`n4. Получение заказа по ID..." -ForegroundColor Yellow
try {
    $getOrderResponse = Invoke-RestMethod -Uri "$baseUrl/api/orders/$orderId" -Method GET
    Write-Host "Заказ получен успешно" -ForegroundColor Green
    Write-Host "Клиент: $($getOrderResponse.data.customer_name)" -ForegroundColor Cyan
    Write-Host "Количество позиций: $($getOrderResponse.data.items.Count)" -ForegroundColor Cyan
} catch {
    Write-Host "Ошибка при получении заказа: $($_.Exception.Message)" -ForegroundColor Red
}

# 5. Обработать платеж
Write-Host "`n5. Обработка платежа..." -ForegroundColor Yellow
$paymentData = @{
    order_id = $orderId
    method = "card"
} | ConvertTo-Json

try {
    $paymentResponse = Invoke-RestMethod -Uri "$baseUrl/api/payments" -Method POST -Body $paymentData -ContentType "application/json"
    $paymentId = $paymentResponse.data.payment.id
    $paymentSuccess = $paymentResponse.data.success
    Write-Host "Платеж обработан. ID: $paymentId" -ForegroundColor Green
    Write-Host "Статус: $($paymentResponse.data.payment.status)" -ForegroundColor Cyan
    Write-Host "Успешность: $paymentSuccess" -ForegroundColor Cyan
} catch {
    Write-Host "Ошибка при обработке платежа: $($_.Exception.Message)" -ForegroundColor Red
}

# 6. Получить статус платежа
Write-Host "`n6. Получение статуса платежа..." -ForegroundColor Yellow
try {
    $paymentStatusResponse = Invoke-RestMethod -Uri "$baseUrl/api/payments/$paymentId" -Method GET
    Write-Host "Статус платежа получен успешно" -ForegroundColor Green
    Write-Host "Статус: $($paymentStatusResponse.data.status)" -ForegroundColor Cyan
    Write-Host "Метод: $($paymentStatusResponse.data.method)" -ForegroundColor Cyan
} catch {
    Write-Host "Ошибка при получении статуса платежа: $($_.Exception.Message)" -ForegroundColor Red
}

# 7. Обновить статус заказа
Write-Host "`n7. Обновление статуса заказа..." -ForegroundColor Yellow
$statusData = @{
    status = "confirmed"
} | ConvertTo-Json

try {
    $statusResponse = Invoke-RestMethod -Uri "$baseUrl/api/orders/$orderId/status" -Method PUT -Body $statusData -ContentType "application/json"
    Write-Host "Статус заказа обновлен успешно" -ForegroundColor Green
    Write-Host "Новый статус: $($statusResponse.data.status)" -ForegroundColor Cyan
} catch {
    Write-Host "Ошибка при обновлении статуса заказа: $($_.Exception.Message)" -ForegroundColor Red
}

# 8. Получить все заказы
Write-Host "`n8. Получение всех заказов..." -ForegroundColor Yellow
try {
    $allOrdersResponse = Invoke-RestMethod -Uri "$baseUrl/api/orders" -Method GET
    Write-Host "Все заказы получены успешно. Количество: $($allOrdersResponse.data.Count)" -ForegroundColor Green
} catch {
    Write-Host "Ошибка при получении всех заказов: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host "`n=== Тестирование завершено ===" -ForegroundColor Green
