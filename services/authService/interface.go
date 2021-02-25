package authService

type AuthService interface {
	Register(username, password string) error
	Login(username, password string) error
}
