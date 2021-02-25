package interfaces

import (
	"cvngur/messaging-service/models"
)

type UserService interface {
	Register(username, password string) error
	Login(username, password string) error
	SendMessage(fromUser, toUser, msg, date string) error
	ViewMessages(username string) ([]models.Message, error)
	BlockUser(username, blockedUser string) error
}
