package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuController struct {
	menu      []MenuItem
	categories []Category
}

func NewMenuController() *MenuController {
	// Инициализируем тестовые данные
	categories := []Category{
		{ID: "1", Name: "Закуски"},
		{ID: "2", Name: "Основные блюда"},
		{ID: "3", Name: "Десерты"},
		{ID: "4", Name: "Напитки"},
		{ID: "5", Name: "Салаты"},
	}

	menu := []MenuItem{
		{
			ID:          "1",
			Name:        "Брускетта с томатами",
			Description: "Итальянская закуска с помидорами и базиликом",
			Price:       350.0,
			Category:    "1",
			Available:   true,
			ImageURL:    "https://example.com/bruschetta.jpg",
		},
		{
			ID:          "2",
			Name:        "Карпаччо из говядины",
			Description: "Тонко нарезанная сырая говядина с оливковым маслом",
			Price:       650.0,
			Category:    "1",
			Available:   true,
			ImageURL:    "https://example.com/carpaccio.jpg",
		},
		{
			ID:          "3",
			Name:        "Стейк Рибай",
			Description: "Сочный стейк из мраморной говядины",
			Price:       1200.0,
			Category:    "2",
			Available:   true,
			ImageURL:    "https://example.com/ribeye.jpg",
		},
		{
			ID:          "4",
			Name:        "Лосось на гриле",
			Description: "Филе лосося с овощами гриль",
			Price:       850.0,
			Category:    "2",
			Available:   true,
			ImageURL:    "https://example.com/salmon.jpg",
		},
		{
			ID:          "5",
			Name:        "Цезарь с курицей",
			Description: "Классический салат с куриным филе",
			Price:       450.0,
			Category:    "5",
			Available:   true,
			ImageURL:    "https://example.com/caesar.jpg",
		},
		{
			ID:          "6",
			Name:        "Греческий салат",
			Description: "Свежие овощи с фетой и оливками",
			Price:       380.0,
			Category:    "5",
			Available:   true,
			ImageURL:    "https://example.com/greek.jpg",
		},
		{
			ID:          "7",
			Name:        "Тирамису",
			Description: "Итальянский десерт с кофе и маскарпоне",
			Price:       420.0,
			Category:    "3",
			Available:   true,
			ImageURL:    "https://example.com/tiramisu.jpg",
		},
		{
			ID:          "8",
			Name:        "Чизкейк",
			Description: "Классический нью-йоркский чизкейк",
			Price:       380.0,
			Category:    "3",
			Available:   true,
			ImageURL:    "https://example.com/cheesecake.jpg",
		},
		{
			ID:          "9",
			Name:        "Латте",
			Description: "Кофе с молоком",
			Price:       180.0,
			Category:    "4",
			Available:   true,
			ImageURL:    "https://example.com/latte.jpg",
		},
		{
			ID:          "10",
			Name:        "Свежевыжатый апельсиновый сок",
			Description: "Натуральный сок из апельсинов",
			Price:       220.0,
			Category:    "4",
			Available:   true,
			ImageURL:    "https://example.com/orange-juice.jpg",
		},
	}

	return &MenuController{
		menu:      menu,
		categories: categories,
	}
}

// GetMenu возвращает полное меню
func (mc *MenuController) GetMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    mc.menu,
	})
}

// GetCategories возвращает категории блюд
func (mc *MenuController) GetCategories(c echo.Context) error {
	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    mc.categories,
	})
}
