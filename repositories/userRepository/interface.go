package userRepository

type UserRepository interface {
	BlockUser(username, blockedUser string) error
}
