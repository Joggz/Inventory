package routes

import (
	"github.com/gofiber/fiber/v2"
	"vending-machine/internal/handlers"
)

func InventoryStockRoutes(app *fiber.App, handler *handlers.InventoryStockHandler)  {
	group := app.Group("/stock")
		group.Get("/inventory/:id/variant/:variant", handler.GetProductVariantByInventoryID)
		group.Get("/inventory/:id/stocks", handler.GetStocksByInventoryID)	
		group.Post("/inventory/:id", handler.AddMultiInventoryStock)
		group.Get("/inventory/stocks", handler.GetStocks)
}