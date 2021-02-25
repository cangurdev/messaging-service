package interfaces

type UserService interface {
	Register(username, password string) error
	Login(username, password string) error
	SendMessage(fromUser, toUser, msg, date string) error
	ViewMessages(username string) error
	BlockUser(username, blockedUser string) error
}
