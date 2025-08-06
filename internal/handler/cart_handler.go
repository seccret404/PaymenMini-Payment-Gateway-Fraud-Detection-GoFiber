package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/usecase"
)

type CartHandler struct {
	cartusecase usecase.CartUsecase
}

func NewCartHandler(cart usecase.CartUsecase) *CartHandler{
	return &CartHandler{cart}
}

func (t *CartHandler) CreateCarts(c *fiber.Ctx) error{
	var cart domain.Cart

	if err := c.BodyParser(&cart); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "beh cemana nih",
		})
	}

	if err := t.cartusecase.CreateCart(&cart); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "add to cart success",
		"data": cart,
	})
}