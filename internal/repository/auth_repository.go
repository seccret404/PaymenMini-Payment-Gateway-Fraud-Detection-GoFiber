package repository

import "github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"

type AuthRepository interface {
	CreateUser(user *domain.User)error
	GetUserByEmail(email string )(*domain.User, error)
}