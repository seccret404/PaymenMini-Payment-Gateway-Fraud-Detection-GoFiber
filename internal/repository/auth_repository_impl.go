package repository

import (
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/pkg/database"
)

type AuthRepositoryImpl struct{}



func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

 
func (a *AuthRepositoryImpl) CreateUser(user *domain.User) error {
	return database.DB.Create(user).Error
}

 
func (a *AuthRepositoryImpl) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User

	result:=database.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}