package usecase

import (
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/repository"
)

type CartUsecaseImpl struct {
	cartrepo repository.CartRepository
}

// CreateCart implements CartUsecase.
func (c *CartUsecaseImpl) CreateCart(cart *domain.Cart) error {
	return c.cartrepo.CreateCart(cart)
}

func NewCartUsecase(cart repository.CartRepository) CartUsecase {
	return &CartUsecaseImpl{cart}
}
