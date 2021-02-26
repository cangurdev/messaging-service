package messageService

import (
	"cvngur/messaging-service/models"
	"cvngur/messaging-service/repositories/messageRepository"
	"errors"
)

type service struct{}

var repository messageRepository.MessageRepository

func NewMessageService(messageRepository messageRepository.MessageRepository) MessageService {
	repository = messageRepository
	return &service{}
}

func (*service) SendMessage(fromUser, toUser, msg, date string) error {
	err := isBlockedUser(fromUser, toUser)
	if err != nil {
		return err
	}

	err = repository.SendMessage(fromUser, toUser, msg, date)
	if err != nil {
		return errors.New("cannot send message")
	}
	return nil
}
func (*service) ViewMessages(username string) ([]models.Message, error) {
	messages, err := repository.GetMessages(username)
	if err != nil {
		return nil, errors.New("cannot view messages")
	}
	return messages, nil
}

func isBlockedUser(fromUser, toUser string) error {
	blockedUsers := repository.GetBlockedUsers(toUser)

	for _, user := range blockedUsers {
		if user == fromUser {
			return errors.New("cannot message to user")
		}
	}
	return nil
}
