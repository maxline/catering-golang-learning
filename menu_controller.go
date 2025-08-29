package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuController struct {
	menu       []MenuItem
	categories []Category
}

func NewMenuController() *MenuController {
	// Initialize test data
	categories := []Category{
		{ID: "1", Name: "Appetizers"},
		{ID: "2", Name: "Main Dishes"},
		{ID: "3", Name: "Desserts"},
		{ID: "4", Name: "Beverages"},
		{ID: "5", Name: "Salads"},
	}

	menu := []MenuItem{
		{
			ID:          "1",
			Name:        "Bruschetta with Tomatoes",
			Description: "Italian appetizer with tomatoes and basil",
			Price:       350.0,
			Category:    "1",
			Available:   true,
			ImageURL:    "https://example.com/bruschetta.jpg",
		},
		{
			ID:          "2",
			Name:        "Beef Carpaccio",
			Description: "Thinly sliced raw beef with olive oil",
			Price:       650.0,
			Category:    "1",
			Available:   true,
			ImageURL:    "https://example.com/carpaccio.jpg",
		},
		{
			ID:          "3",
			Name:        "Ribeye Steak",
			Description: "Juicy marbled beef steak",
			Price:       1200.0,
			Category:    "2",
			Available:   true,
			ImageURL:    "https://example.com/ribeye.jpg",
		},
		{
			ID:          "4",
			Name:        "Grilled Salmon",
			Description: "Salmon fillet with grilled vegetables",
			Price:       850.0,
			Category:    "2",
			Available:   true,
			ImageURL:    "https://example.com/salmon.jpg",
		},
		{
			ID:          "5",
			Name:        "Caesar Salad with Chicken",
			Description: "Classic salad with chicken breast",
			Price:       450.0,
			Category:    "5",
			Available:   true,
			ImageURL:    "https://example.com/caesar.jpg",
		},
		{
			ID:          "6",
			Name:        "Greek Salad",
			Description: "Fresh vegetables with feta cheese and olives",
			Price:       380.0,
			Category:    "5",
			Available:   true,
			ImageURL:    "https://example.com/greek.jpg",
		},
		{
			ID:          "7",
			Name:        "Tiramisu",
			Description: "Italian dessert with coffee and mascarpone",
			Price:       420.0,
			Category:    "3",
			Available:   true,
			ImageURL:    "https://example.com/tiramisu.jpg",
		},
		{
			ID:          "8",
			Name:        "Cheesecake",
			Description: "Classic New York cheesecake",
			Price:       380.0,
			Category:    "3",
			Available:   true,
			ImageURL:    "https://example.com/cheesecake.jpg",
		},
		{
			ID:          "9",
			Name:        "Latte",
			Description: "Coffee with milk",
			Price:       180.0,
			Category:    "4",
			Available:   true,
			ImageURL:    "https://example.com/latte.jpg",
		},
		{
			ID:          "10",
			Name:        "Fresh Orange Juice",
			Description: "Natural juice from oranges",
			Price:       220.0,
			Category:    "4",
			Available:   true,
			ImageURL:    "https://example.com/orange-juice.jpg",
		},
	}

	return &MenuController{
		menu:       menu,
		categories: categories,
	}
}

// GetMenu returns the full menu
func (mc *MenuController) GetMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    mc.menu,
	})
}

// GetCategories returns dish categories
func (mc *MenuController) GetCategories(c echo.Context) error {
	return c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Data:    mc.categories,
	})
}
