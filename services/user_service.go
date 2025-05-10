package services

import (
	"github.com/BerniceZTT/goadmin/models"
	"github.com/BerniceZTT/goadmin/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}
