package handlers

import (
	"net/http"
	"vending-machine/internal/migrations"
	"vending-machine/internal/services"

	"github.com/gofiber/fiber/v2"
)

type InventoryService struct{
	svc *services.InventoryService
 }

 func NewInventoryHandler(svc *services.InventoryService ) *InventoryService{
	return  &InventoryService{svc: svc}
 }


 func (h *InventoryService) CreateInventory(c *fiber.Ctx) error {
	var payload []migrations.Inventory
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.svc.Create(payload); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"message": "inventory Created",
	})
 }


 func (h *InventoryService) GetAllInventory(c *fiber.Ctx) error {
	invs, err := h.svc.GetAllInventory()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"inventories": invs,
	})
 }

 func (h *InventoryService) GetInventoryByID(c *fiber.Ctx) error  {
	id, err := c.ParamsInt("id");
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	inv, err := h.svc.GetInventoryByID(int64(id));
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": http.StatusOK,
		"inventory": inv,
	})
 }