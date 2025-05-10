package services

import "github.com/BerniceZTT/goadmin/models"

type UserServiceInterface interface {
	GetUser(id uint) (*models.User, error)
	CreateUser(user *models.User) error
}
