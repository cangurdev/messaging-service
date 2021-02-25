package tests

import "github.com/stretchr/testify/mock"

type MockUserRepository struct {
	mock.Mock
}

func (m MockUserRepository) BlockUser(username, blockedUser string) error {
	args := m.Called()
	return args.Error(0)
}
