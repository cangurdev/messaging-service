package repositories

import (
	"context"
	"cvngur/messaging-service/db"
	"cvngur/messaging-service/domain"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type mRepo struct {
}

func NewMessageRepository() domain.MessageRepository {
	return &mRepo{}
}

func (*mRepo) SendMessage(fromUser, toUser, msg, date string) error {

	var message domain.Message
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
func (*mRepo) GetMessages(username string) ([]domain.Message, error) {
	filter := bson.D{{Key: "username", Value: username}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user domain.User
	err := db.Connection().Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, errors.New("user cannot found")
	}

	return user.Messages, nil
}

func (*mRepo) GetBlockedUsers(username string) []string {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := db.Connection().Collection("users").FindOne(ctx, bson.D{{Key: "username", Value: username}}).Decode(&User)

	if err != nil {
		return nil
	}
	return User.BlockedUsers
}
