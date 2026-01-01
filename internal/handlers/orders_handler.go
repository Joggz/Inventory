package handlers

import (
	"net/http"

	"vending-machine/internal/migrations"
	"vending-machine/internal/repository"


	// "vending-machine/internal/services"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	orderService *repository.PurchaseOrderRepository
}

func NewOrderHandler(orderService *repository.PurchaseOrderRepository) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}




func (oh *OrderHandler) CreateOrder(c *fiber.Ctx) error {

	var payload migrations.Orders
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	
	}

	orderID, err := oh.orderService.CreatePurchaseOrder(payload.InventoryID, payload.ProductVariantID, int64(payload.Quantity), payload.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"order_id": orderID})
}