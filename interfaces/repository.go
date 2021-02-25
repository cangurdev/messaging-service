package interfaces

import (
	"cvngur/messaging-service/models"
)

type UserRepository interface {
	SaveUser(username, password string) error
	GetUser(username string) (string, error)
	SendMessage(fromUser, toUser, msg, date string) error
	GetMessages(username string) ([]models.Message, error)
	BlockUser(username, blockedUser string) error
	GetBlockedUsers(username string) []string
}
