package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"vending-machine/internal/services"
)


type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(productService *services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error   {

	var payload services.ProductInfo;
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productId, err := h.productService.CreateProduct(payload);
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
			})
	}

	return  c.Status(http.StatusCreated).JSON(fiber.Map{
		"productId": productId,
		"status": http.StatusCreated,
		"message": "product Created successfully",
	})
}
