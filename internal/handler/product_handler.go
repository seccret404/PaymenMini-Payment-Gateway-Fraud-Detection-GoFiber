package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/usecase"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(product usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{product}
}

func (p *ProductHandler) CreateProducts(c *fiber.Ctx) error {
	var product domain.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := p.productUsecase.CreateProduct(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Add product success",
		"data":    product,
	})
}

func (p *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	product, err := p.productUsecase.GetAllProduct()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cant find your product",
		})
	}
	return c.JSON(product)
}

func (p *ProductHandler) GetByIdProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := p.productUsecase.GetByIdProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "product not found",
		})
	}

	return c.JSON(product)
}

func (p *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var input domain.Product
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	product, err := p.productUsecase.UpdateProduct(uint(id), &input)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(product)
}

func (p *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := p.productUsecase.DeleteProduct(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "product deleted!",
	})
}
