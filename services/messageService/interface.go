package messageService

import "cvngur/messaging-service/models"

type MessageService interface {
	SendMessage(fromUser, toUser, msg, date string) error
	ViewMessages(username string) ([]models.Message, error)
}
