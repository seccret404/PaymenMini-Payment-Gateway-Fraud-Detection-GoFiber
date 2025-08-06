package usecase

import "github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"

type AuthUsecase interface {
	Register(user *domain.User)error
	Login(email, password string)(*domain.User, error)
}