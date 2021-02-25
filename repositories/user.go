package repositories

import (
	"context"
	"cvngur/messaging-service/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	SaveUser(username, password string) error
	ValidateUser(username, password string) (bool, error)
}

type repository struct {
}

func NewUserRepository() UserRepository {
	return &repository{}
}

func (*repository) SaveUser(username, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := db.Connection().Collection("users").InsertOne(ctx, bson.D{{Key: "username", Value: username}, {
		Key: "password", Value: password,
	}})

	if err != nil {
		return err
	}

	return nil
}
func (*repository) ValidateUser(username, password string) (bool, error) {
	return true, nil
}
