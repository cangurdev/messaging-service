package services

import (
	"cvngur/messaging-service/interfaces"
	"cvngur/messaging-service/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type service struct{}

var repository interfaces.UserRepository

func NewUserService(userRepository interfaces.UserRepository) interfaces.UserService {
	repository = userRepository
	return &service{}
}

func (*service) Register(username, password string) error {
	if !isAvailableUsername(username) {
		return errors.New("username is not available")
	}
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = repository.SaveUser(username, string(bytePassword))
	if err != nil {
		return err
	}
	return nil
}

func (*service) Login(username, password string) error {

	hashedPassword, err := repository.GetUser(username)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("invalid username or password")
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
			return errors.New("cannot message to user")
		}
	}
	return nil
}
func isAvailableUsername(username string) bool {
	_, err := repository.GetUser(username)
	if err != nil {
		return true
	}
	return false
}
func (*service) ViewMessages(username string) ([]models.Message, error) {
	messages, err := repository.GetMessages(username)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
func (*service) BlockUser(username, blockedUser string) error {
	err := repository.BlockUser(username, blockedUser)
	if err != nil {
		return err
	}
	return nil
}
