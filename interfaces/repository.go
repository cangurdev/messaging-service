package interfaces

type UserRepository interface {
	SaveUser(username, password string) error
	ValidateUser(username, password string) error
	SendMessage(username, msg, date string) error
	GetMessages(username string) error
	BlockUser(username, blockedUser string) error
}
