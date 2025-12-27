package handlers

import (
	"net/http"
	"strconv"
	"vending-machine/internal/migrations"
	"vending-machine/internal/services"

	"github.com/gofiber/fiber/v2"
	
)

type InventoryStockHandler struct {
	ish *services.InventoryStockService
}

func NewInventoryStockHandler(iss *services.InventoryStockService) *InventoryStockHandler  {
	return &InventoryStockHandler{ish: iss}
}


func (h *InventoryStockHandler) AddMultiInventoryStock(c *fiber.Ctx) error {
	inventoryID, _ := strconv.ParseInt(
			c.Params("id"), 10, 64,
		)
	var payload migrations.AddMultipleStockPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid payload",
		})
	}

	var items []migrations.AddStockItem;
	for _, item := range payload.Items {
		items = append(items, migrations.AddStockItem{
			ProductVariantID: item.VariantID,
			Quantity:  item.Quantity,
		})
	}

	if err := h.ish.AddMultipleStock(inventoryID, items); err != nil {
		return c.Status(400).JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status" :  http.StatusCreated,
		"message": "Stock added successfully",
	})

}