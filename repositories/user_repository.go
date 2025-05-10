package repositories

import (
	"github.com/BerniceZTT/goadmin/config"
	"github.com/BerniceZTT/goadmin/models"
)

type UserRepository struct{}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)
	return &user, result.Error
}

func (r *UserRepository) Create(user *models.User) error {
	result := config.DB.Create(user)
	return result.Error
}
