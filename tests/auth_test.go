package tests

import (
	"cvngur/messaging-service/services/authService"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_return_nil_when_user_register(t *testing.T) {
	//Arrange
	mockRepo := new(MockAuthRepository)
	aService := authService.NewAuthService(mockRepo)
	mockRepo.On("GetUser").Return("", errors.New(""))
	mockRepo.On("SaveUser").Return(nil)

	//Act
	sut := aService.Register("Can", "123456")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.Equal(t, nil, sut)
}
func Test_should_return_error_when_register_with_existed_username(t *testing.T) {
	//Arrange
	mockRepo := new(MockAuthRepository)
	aService := authService.NewAuthService(mockRepo)
	mockRepo.On("GetUser").Return("", nil)

	//Act
	sut := aService.Register("Can", "123456")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, sut, "username is not available")
}

func Test_should_return_error_when_login_with_invalid_username(t *testing.T) {
	//Arrange
	mockRepo := new(MockAuthRepository)
	aService := authService.NewAuthService(mockRepo)
	mockRepo.On("GetUser").Return("", errors.New("user cannot found"))

	//Act
	sut := aService.Login("Can", "123456")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, sut, "user cannot found")
}
func Test_should_return_error_when_login_with_wrong_password(t *testing.T) {
	//Arrange
	mockRepo := new(MockAuthRepository)
	aService := authService.NewAuthService(mockRepo)
	mockRepo.On("GetUser").Return("", nil)

	//Act
	sut := aService.Login("Can", "123456")

	//Assert
	mockRepo.AssertExpectations(t)
	assert.EqualError(t, sut, "invalid username or password")
}
