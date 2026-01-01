package routes

import (
	"github.com/gofiber/fiber/v2"
	"vending-machine/internal/handlers"
)

func  PurchaseStockRoutes(app *fiber.App, orderHandler *handlers.OrderHandler)  {
	purchase := app.Group("/buy")
	purchase.Post("/product", orderHandler.CreateOrder)
	purchase.Get("/confirm/:ref", orderHandler.ConfirmPayment)
}