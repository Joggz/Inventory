package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"vending-machine/internal/services"
)

type VendingHandler struct {
	svc services.VendingService
}

func NewVendingHandler(s services.VendingService) *VendingHandler {
	return &VendingHandler{svc: s}
}

func (h *VendingHandler) Health(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "ok"})
}

func (h *VendingHandler) Inventory(c *fiber.Ctx) error {
	items, err := h.svc.Inventory()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(items)
}

type purchaseRequest struct {
	ItemID    string `json:"item_id"`
	Quantity  int    `json:"quantity"`
	PaidCents int    `json:"paid_cents"`
}

func (h *VendingHandler) Purchase(c *fiber.Ctx) error {
	var req purchaseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	res, err := h.svc.Purchase(&services.PurchaseInput{
		ItemID:    req.ItemID,
		Quantity:  req.Quantity,
		PaidCents: req.PaidCents,
	})
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"item_id":            res.Item.ID,
		"name":               res.Item.Name,
		"price_cents":        res.Item.PriceCents,
		"total_cents":        res.TotalCents,
		"change_cents":       res.ChangeCents,
		"remaining_quantity": res.RemainingQuantity,
	})
}

type restockRequest struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

func (h *VendingHandler) Restock(c *fiber.Ctx) error {
	var req restockRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	item, err := h.svc.Restock(req.ItemID, req.Quantity)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(item)
}

