package tests

import (
	"cvngur/messaging-service/app/services"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_return_nil_when_user_blocked_someone(t *testing.T) {
	//Arrange
	mockRepo := new(MockUserRepository)
	uService := services.NewUserService(mockRepo)
	mockRepo.On("BlockUser").Return(nil)

	//Act
	sut := uService.BlockUser("Can", "Ali")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, sut)
}
func Test_should_return_error_when_blocked_nonexistent_user(t *testing.T) {
	//Arrange
	mockRepo := new(MockUserRepository)
	uService := services.NewUserService(mockRepo)
	mockRepo.On("BlockUser").Return(errors.New(""))

	//Act
	sut := uService.BlockUser("Can", "Ali")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.NotNil(t, sut)
}
