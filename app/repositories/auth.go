package repositories

import (
	"context"
	"cvngur/messaging-service/db"
	"cvngur/messaging-service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type repository struct {
}

func NewAuthRepository() domain.AuthRepository {
	return &repository{}
}

func (*repository) SaveUser(username, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	messages := make([]domain.Message, 0)
	blockedUsers := make([]domain.Message, 0)

	_, err := db.Connection().Collection("users").InsertOne(ctx, bson.D{{Key: "username", Value: username}, {
		Key: "password", Value: password,
	}, {Key: "messages", Value: messages}, {Key: "blockedUsers", Value: blockedUsers}})

	if err != nil {
		return err
	}

	return nil
}

func (*repository) GetUser(username string) (string, error) {

	filter := bson.D{{Key: "username", Value: username}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := db.Connection().Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return "", errors.New("user cannot found")
	}
	return user.Password, nil
}
