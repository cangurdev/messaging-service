package messageRepository

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

func NewMessageRepository() MessageRepository {
	return &repository{}
}

var User struct {
	Username     string           `json:"username"`
	Password     string           `json:"password"`
	Messages     []models.Message `json:"messages"`
	BlockedUsers []string         `json:"blockedUsers"`
}

func (*repository) SendMessage(fromUser, toUser, msg, date string) error {

	var message models.Message
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	message.FromUser = fromUser
	message.ToUser = toUser
	message.Msg = msg
	message.Date = date

	_, err := db.Connection().Collection("users").UpdateOne(ctx, bson.M{"username": fromUser}, bson.D{{"$push", bson.D{{"messages", message}}}})
	if err != nil {
		return err
	}
	_, err = db.Connection().Collection("users").UpdateOne(ctx, bson.M{"username": toUser}, bson.D{{"$push", bson.D{{"messages", message}}}})
	if err != nil {
		return err
	}
	return nil
}
func (*repository) GetMessages(username string) ([]models.Message, error) {
	filter := bson.D{{Key: "username", Value: username}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.Connection().Collection("users").FindOne(ctx, filter).Decode(&User)
	if err != nil {
		return nil, errors.New("user cannot found")
	}

	return User.Messages, nil
}

func (*repository) GetBlockedUsers(username string) []string {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.Connection().Collection("users").FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&User)

	if err != nil {
		return nil
	}
	return User.BlockedUsers
}
