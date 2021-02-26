package authService

import (
	"cvngur/messaging-service/logs"
	"cvngur/messaging-service/repositories/authRepository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var repository authRepository.AuthRepository

type service struct{}

func NewAuthService(authRepository authRepository.AuthRepository) AuthService {
	repository = authRepository
	return &service{}
}

func (*service) Register(username, password string) error {
	if isAvailableUsername(username) {
		return errors.New("username is not available")
	}
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("cannot create a new user")
	}
	err = repository.SaveUser(username, string(bytePassword))
	if err != nil {
		return errors.New("cannot create a new user")
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
		logs.Write("Invalid Login")
		return errors.New("invalid username or password")
	}
	logs.Write("Logged")
	return nil
}
func isAvailableUsername(username string) bool {
	_, err := repository.GetUser(username)
	if err != nil {
		return false
	}
	return true
}
