package domain

type AuthRepository interface {
	SaveUser(username, password string) error
	GetUser(username string) (string, error)
}
type AuthService interface {
	Register(username, password string) error
	Login(username, password string) error
}
