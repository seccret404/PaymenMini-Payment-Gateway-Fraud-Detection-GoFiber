package usecase

import "github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"

type ProductUsecase interface {
	CreateProduct(product *domain.Product)error
	GetAllProduct() ([]domain.Product,error)
	GetByIdProduct(id uint)(*domain.Product, error)
	UpdateProduct(id uint, input *domain.Product)(*domain.Product, error)
	DeleteProduct(id uint)error
}
