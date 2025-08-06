package repository

import "github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"

type ProductRepository interface {
	CreateProduct(product *domain.Product) error
	GetProduct() ([]domain.Product, error)
	GetByIdProduct(id uint)(*domain.Product, error)
	UpdateProduct(product *domain.Product) error
	DeleteProduct(id uint)error	
}