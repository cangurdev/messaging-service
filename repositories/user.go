package repositories

import (
	"context"
	"cvngur/messaging-service/db"
	"cvngur/messaging-service/interfaces"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type repository struct {
}

func NewUserRepository() interfaces.UserRepository {
	return &repository{}
}

func (*repository) SaveUser(username, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var messages []string
	var blockedUsers []string

	_, err := db.Connection().Collection("users").InsertOne(ctx, bson.D{{Key: "username", Value: username}, {
		Key: "password", Value: password,
	}, {Key: "messages", Value: messages}, {Key: "blockedUsers", Value: blockedUsers}})

	if err != nil {
		return err
	}

	return nil
}

var User struct {
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Messages     []string `json:"messages"`
	blockedUsers []string `json:"blockedUsers"`
}

func (*repository) ValidateUser(username, password string) error {

	filter := bson.D{{Key: "username", Value: username}, {Key: "password", Value: password}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.Connection().Collection("users").FindOne(ctx, filter).Decode(&User)
	if err != nil {
		return err
	}
	return nil
}

func (*repository) SendMessage(username, msg, date string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Connection().Collection("users").UpdateOne(ctx, bson.M{"username": username}, bson.D{{"$push", bson.D{{"messages", msg}}}})
	if err != nil {
		return err
	}
	return nil
}
func (*repository) GetMessages(username string) error {
	return nil
}
func (*repository) BlockUser(username string) error {
	return nil
}
