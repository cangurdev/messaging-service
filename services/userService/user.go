package userService

import "cvngur/messaging-service/repositories/userRepository"

type service struct{}

var repository userRepository.UserRepository

func NewUserService(userRepository userRepository.UserRepository) UserService {
	repository = userRepository
	return &service{}
}

func (*service) BlockUser(username, blockedUser string) error {
	err := repository.BlockUser(username, blockedUser)
	if err != nil {
		return err
	}
	return nil
}
