package service

import (
	"errors"
	"go-backend-boilerplate/internal/models"
	"go-backend-boilerplate/internal/repo"
	"go-backend-boilerplate/internal/utils"
)

type UserService interface {
	Register(email, password string) error
}

type userService struct {
	repo repo.UserRepository
}

func NewUserService(r repo.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Register(email, password string) error {
	existing, _ := s.repo.FindByEmail(email)
	if existing != nil {
		return errors.New("user already exists")
	}

	hashedPwd, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    email,
		Password: hashedPwd,
	}

	return s.repo.Create(user)
}
