package usecase

import (
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/repository"
)

type ProductUsecaseImpl struct {
	productrepo repository.ProductRepository
}

func NewProductUsecase(product repository.ProductRepository) ProductUsecase {
	return &ProductUsecaseImpl{product}
}

func (p *ProductUsecaseImpl) CreateProduct(product *domain.Product) error {
	return p.productrepo.CreateProduct(product)
}

func (p *ProductUsecaseImpl) GetAllProduct() ([]domain.Product, error) {
	return p.productrepo.GetProduct()
}

func (p *ProductUsecaseImpl) GetByIdProduct(id uint) (*domain.Product, error) {
	return p.productrepo.GetByIdProduct(id)
}

func (p *ProductUsecaseImpl) UpdateProduct(id uint, input *domain.Product) (*domain.Product, error) {
	product, err := p.productrepo.GetByIdProduct(id)
	if err != nil{
		return nil, err
	}

	if input.Name != ""{
		product.Name = input.Name
	}
	if input.Description != ""{
		product.Description = input.Description
	}	
	if input.Price != ""{
		product.Price = input.Price
	}
	if input.Quantity != ""{
		product.Quantity = input.Quantity
	}
	err = p.productrepo.UpdateProduct(product)

	return product, err
}
func (p *ProductUsecaseImpl) DeleteProduct(id uint) error {
	return p.productrepo.DeleteProduct(id)
}