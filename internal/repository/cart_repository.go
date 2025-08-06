package repository

import "github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"

type CartRepository interface {
	CreateCart(cart *domain.Cart) error
}