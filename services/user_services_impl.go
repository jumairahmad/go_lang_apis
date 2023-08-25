package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"restapi.com/my-apis/models"
)

type userServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserServices {

	return &userServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *userServiceImpl) CreateUser(user *models.User) error {

	return nil
}

func (u *userServiceImpl) GetUsers(name *string) (*models.User, error) {

	return nil, nil
}

func (u *userServiceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *userServiceImpl) UpdateUser(*models.User) error {
	return nil
}

func (u *userServiceImpl) DeleteUser(name *string) error {
	return nil
}
