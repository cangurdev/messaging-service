package authRepository

type AuthRepository interface {
	SaveUser(username, password string) error
	GetUser(username string) (string, error)
}
