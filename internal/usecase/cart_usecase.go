package usecase

import "github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"

type CartUsecase interface {
	CreateCart(cart *domain.Cart) error
}