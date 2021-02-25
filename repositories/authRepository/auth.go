package authRepository

import (
	"context"
	"cvngur/messaging-service/db"
	"cvngur/messaging-service/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type repository struct {
}

func NewUserRepository() AuthRepository {
	return &repository{}
}

var User struct {
	Username     string           `json:"username"`
	Password     string           `json:"password"`
	Messages     []models.Message `json:"messages"`
	BlockedUsers []string         `json:"blockedUsers"`
}

func (*repository) SaveUser(username, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	messages := make([]models.Message, 0)
	blockedUsers := make([]models.Message, 0)

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
	err := db.Connection().Collection("users").FindOne(ctx, filter).Decode(&User)
	if err != nil {
		return "", errors.New("user cannot found")
	}
	return User.Password, nil
}
