package messageRepository

import "cvngur/messaging-service/models"

type MessageRepository interface {
	SendMessage(fromUser, toUser, msg, date string) error
	GetMessages(username string) ([]models.Message, error)
	GetBlockedUsers(username string) []string
}
