package userService

type UserService interface {
	BlockUser(username, blockedUser string) error
}
