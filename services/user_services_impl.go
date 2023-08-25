package services

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
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

	_, err := u.usercollection.InsertOne(u.ctx, u)
	return err
}

func (u *userServiceImpl) GetUsers(name *string) (*models.User, error) {

	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)

	return user, err
}

func (u *userServiceImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *userServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_age", Value: user.Age}, bson.E{Key: "user_address", Value: user.Address}}}}

	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No Matched Found For Update")
	}
	return nil
}

func (u *userServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: &name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No Matched Found For Delete")
	}
	return nil
}
