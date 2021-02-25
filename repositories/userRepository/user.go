package userRepository

import (
	"context"
	"cvngur/messaging-service/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type repository struct {
}

func NewUserRepository() UserRepository {
	return &repository{}
}
func (*repository) BlockUser(username, blockedUser string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := db.Connection().Collection("users").UpdateOne(ctx, bson.M{"username": username}, bson.D{{"$push", bson.D{{"blockedUsers", blockedUser}}}})
	if err != nil {
		return err
	}
	return nil
}
