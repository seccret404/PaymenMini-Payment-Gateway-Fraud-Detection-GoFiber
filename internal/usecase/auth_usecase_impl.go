package usecase

import (
	"errors"

	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/domain"
	"github.com/seccret404/PaymenMini-Payment-Gateway-Fraud-Detection-GoFiber/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecaseImpl struct {
	authRepo repository.AuthRepository
}

func NewAuthUsecase(repo repository.AuthRepository) AuthUsecase {
	return &AuthUsecaseImpl{repo}
}

func (a *AuthUsecaseImpl) Login(email string, password string) (*domain.User, error) {
    user, err := a.authRepo.GetUserByEmail(email)
    if err != nil {
        return nil, errors.New("invalid email or password")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid email or password")
    }

    return user, nil
    
}


func (a *AuthUsecaseImpl) Register(user *domain.User) error {
	hshPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hshPassword)

	return a.authRepo.CreateUser(user)
}