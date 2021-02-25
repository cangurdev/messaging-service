package tests

import "github.com/stretchr/testify/mock"

type MockAuthRepository struct {
	mock.Mock
}

func (m MockAuthRepository) SaveUser(username, password string) error {
	args := m.Called()
	return args.Error(0)
}

func (m MockAuthRepository) GetUser(username string) (string, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(string), args.Error(1)
}
