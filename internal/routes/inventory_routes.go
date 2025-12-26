package routes

import (
	"github.com/gofiber/fiber/v2"
	"vending-machine/internal/handlers"
)


func InventoryRoutes(app *fiber.App, handler *handlers.InventoryService)  {
	group := app.Group("/inventory");
	group.Post("/create", handler.CreateInventory);
	group.Get("/inventories", handler.GetAllInventory);
	group.Get("/:id", handler.GetInventoryByID);
}
