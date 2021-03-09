package services

import (
	"cvngur/messaging-service/domain"
)

type uService struct{}

var uRepo domain.UserRepository

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	uRepo = userRepository
	return &uService{}
}

func (*uService) BlockUser(username, blockedUser string) error {
	err := uRepo.BlockUser(username, blockedUser)
	if err != nil {
		return err
	}
	return nil
}
