package services

import "restapi.com/my-apis/models"

type UserServices interface {
	CreateUser(*models.User) error
	GetUsers(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
