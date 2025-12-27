package routes

import (
	"github.com/gofiber/fiber/v2"
	"vending-machine/internal/handlers"
)


func ProductRoutes(app *fiber.App, handler *handlers.ProductHandler)  {
	product := app.Group("/products");
	product.Post("/create", handler.CreateProduct);
}