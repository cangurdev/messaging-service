package services

import (
	"cvngur/messaging-service/domain"
	"errors"
)

type mService struct{}

var mRepo domain.MessageRepository

func NewMessageService(messageRepository domain.MessageRepository) domain.MessageService {
	mRepo = messageRepository
	return &mService{}
}

func (*mService) SendMessage(fromUser, toUser, msg, date string) error {
	err := isBlockedUser(fromUser, toUser)
	if err != nil {
		return err
	}

	err = mRepo.SendMessage(fromUser, toUser, msg, date)
	if err != nil {
		return errors.New("cannot send message")
	}
	return nil
}
func (*mService) ViewMessages(username string) ([]domain.Message, error) {
	messages, err := mRepo.GetMessages(username)
	if err != nil {
		return nil, errors.New("cannot view messages")
	}
	return messages, nil
}

func isBlockedUser(fromUser, toUser string) error {
	blockedUsers := mRepo.GetBlockedUsers(toUser)

	for _, user := range blockedUsers {
		if user == fromUser {
			return errors.New("cannot message to user")
		}
	}
	return nil
}
