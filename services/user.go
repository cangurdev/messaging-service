package services

import (
	"cvngur/messaging-service/interfaces"
	"errors"
)

type service struct{}

var repository interfaces.UserRepository

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	repository = userRepository
	return &service{}
}

func (*service) Register(username, password string) error {
	err := repository.SaveUser(username, password)
	if err != nil {
		return err
	}
	return nil
}

func (*service) Login(username, password string) error {
	err := repository.ValidateUser(username, password)
	if err != nil {
		return err
	}
	return nil
}

func (*service) SendMessage(fromUser, toUser, msg, date string) error {
	err := isBlockedUser(fromUser, toUser)
	if err != nil {
		return err
	}

	err = repository.SendMessage(fromUser, toUser, msg, date)
	if err != nil {
		return err
	}

	return nil
}

func isBlockedUser(fromUser, toUser string) error {
	blockedUsers := repository.GetBlockedUsers(toUser)

	for _, user := range blockedUsers {
		if user == fromUser {
			return errors.New("Cannot message to user")
		}
	}
	return nil
}
func (*service) ViewMessages(username string) error {
	return nil
}
func (*service) BlockUser(username, blockedUser string) error {
	err := repository.BlockUser(username, blockedUser)
	if err != nil {
		return err
	}
	return nil
}
