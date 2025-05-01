package service

import (
	"errors"
	"go-backend-boilerplate/internal/models"
	"go-backend-boilerplate/internal/repo"
	"go-backend-boilerplate/internal/utils"
)

type UserService interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
	GetById(id uint) (*models.User, error)
	List() ([]models.User, error)
	Update(id uint, email string) error
	Delete(id uint) error
}

type userService struct {
	repo   repo.UserRepository
	secret string
}

func NewUserService(repo repo.UserRepository, secret string) UserService {
	return &userService{
		repo:   repo,
		secret: secret,
	}
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

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, s.secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetById(id uint) (*models.User, error) {
	return s.repo.FindById(id)
}

func (s *userService) List() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) Update(id uint, email string) error {
	user, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	user.Email = email
	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}
