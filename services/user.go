package services

import (
	"cvngur/messaging-service/interfaces"
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

func (*service) SendMessage(username, msg, date string) error {
	err := repository.SendMessage(username, msg, date)
	if err != nil {
		return err
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
