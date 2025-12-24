package routes

import (
	"github.com/gofiber/fiber/v2"
	"vending-machine/internal/handlers"
)

func Register(app *fiber.App, h *handlers.VendingHandler) {
	app.Get("/health", h.Health)
	api := app.Group("/api")
	api.Get("/inventory", h.Inventory)
	api.Post("/purchase", h.Purchase)
	api.Post("/restock", h.Restock)
}

