package repository

import (
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/pkg/database"
)

type ProductRepositoryImpl struct{}



func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p *ProductRepositoryImpl) CreateProduct(product *domain.Product) error {
	if err := database.DB.Create(product).Error; err != nil{
		return err
	}

	return database.DB.Preload("User").First(product, product.ID).Error 

}

// GetProduct implements ProductRepository.
func (p *ProductRepositoryImpl) GetProduct() ([]domain.Product, error) {
	var products []domain.Product
	err := database.DB.Find(&products).Error
	return products, err
}


func (p *ProductRepositoryImpl) GetByIdProduct(id uint) (*domain.Product, error) {
	var product domain.Product

	if err := database.DB.Find(&product, id).Error; err != nil{
		return nil, err
	}

	return &product, nil
}

func (p *ProductRepositoryImpl) UpdateProduct(product *domain.Product) error {
	return database.DB.Save(product).Error
}

func (p *ProductRepositoryImpl) DeleteProduct(id uint) error {
	return database.DB.Delete(&domain.Product{}, id).Error
}
