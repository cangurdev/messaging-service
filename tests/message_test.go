package tests

import (
	"cvngur/messaging-service/models"
	"cvngur/messaging-service/services/messageService"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_return_nil_when_user_send_message(t *testing.T) {

	//Arrange
	mockRepo := new(MockMessageRepository)
	mService := messageService.NewMessageService(mockRepo)
	mockRepo.On("SendMessage").Return(nil)

	var blockedUsers []string
	mockRepo.On("GetBlockedUsers").Return(blockedUsers)

	//Act
	sut := mService.SendMessage("Can", "Ali", "Hello", "01.01.2021")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, sut)
}
func Test_should_return_error_when_blocked_user_send_message(t *testing.T) {

	//Arrange
	mockRepo := new(MockMessageRepository)
	mService := messageService.NewMessageService(mockRepo)

	blockedUsers := []string{"Can", "Burak"}
	mockRepo.On("GetBlockedUsers").Return(blockedUsers)

	//Act
	sut := mService.SendMessage("Can", "Ali", "Hi", "01.01.2021")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, sut, "cannot message to user")
}
func Test_should_return_messages(t *testing.T) {

	//Arrange
	mockRepo := new(MockMessageRepository)
	mService := messageService.NewMessageService(mockRepo)

	messages := []models.Message{{"Can", "Ali", "Hello", "01.01.2021"},
		{"Ali", "Can", "Hi", "01.01.2021"}}
	mockRepo.On("GetMessages").Return(messages, nil)

	//Act
	sut, err := mService.ViewMessages("Can")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, err)
	assert.NotNil(t, sut)
}
