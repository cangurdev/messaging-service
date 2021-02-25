package services

import (
	"cvngur/messaging-service/interfaces"
	"cvngur/messaging-service/repositories"
)

type service struct{}

var repository repositories.UserRepository

func NewUserService(userRepository repositories.UserRepository) interfaces.UserService {
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

func (*service) Login(username, password string) (bool, error) {
	result, err := repository.ValidateUser(username, password)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (*service) SendMessage(username string) error {
	return nil
}
func (*service) ViewMessages(username string) error {
	return nil
}
func (*service) BlockUser(username string) error {
	return nil
}
