package interfaces

type UserService interface {
	Register(username, password string) error
	Login(username, password string) (bool, error)
	SendMessage(username string) error
	ViewMessages(username string) error
	BlockUser(username string) error
}
