package tests

import (
	"cvngur/messaging-service/domain"
	"github.com/stretchr/testify/mock"
)

type MockMessageRepository struct {
	mock.Mock
}

func (m *MockMessageRepository) SendMessage(fromUser, toUser, msg, date string) error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockMessageRepository) GetMessages(username string) ([]domain.Message, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]domain.Message), args.Error(1)
}
func (m *MockMessageRepository) GetBlockedUsers(username string) []string {
	args := m.Called()
	result := args.Get(0)
	return result.([]string)
}
