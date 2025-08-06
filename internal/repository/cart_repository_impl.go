package repository

import (
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/pkg/database"
)

type CartRepositoryImpl struct{}

// CreateCart implements CartRepository.
func (c *CartRepositoryImpl) CreateCart(cart *domain.Cart) error {
	if err := database.DB.Create(cart).Error; err != nil{
		return err
	}
	//preload constraint relations
	return database.DB.Preload("User").Preload("Product").First(cart, cart.ID).Error
}

func NewCartRepository() CartRepository {
	return &CartRepositoryImpl{}
}
