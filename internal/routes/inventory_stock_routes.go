package routes

import (
	"github.com/gofiber/fiber/v2"
	"vending-machine/internal/handlers"
)

func InventoryStockRoutes(app *fiber.App, handler *handlers.InventoryStockHandler)  {
	group := app.Group("/stock")
	group.Post("/inventory/:id", handler.AddMultiInventoryStock)
}