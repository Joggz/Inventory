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
			Price:     item.Price,
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
		// "items": items,
	})

}


func (h *InventoryStockHandler) GetStocks(c *fiber.Ctx)error{
	stock, err := h.ish.GetAllStocks()
	if err != nil {
		return  c.Status(http.StatusBadGateway).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "list of stocks",
		"status": http.StatusOK,
		"stocks": stock,
	})

 }

 func (h *InventoryStockHandler) GetStocksByInventoryID(c *fiber.Ctx) error {

		inventoryID, err := strconv.ParseInt(c.Params("id"),10, 64,)
		if err != nil {
			return  err
		}

	stock, err :=h.ish.GetStocksByInventoryID(inventoryID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message":  "Internal Server Error",
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"stock": stock,
			"status" : http.StatusOK,
		})

 }