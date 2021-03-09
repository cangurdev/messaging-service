package domain

type UserRepository interface {
	BlockUser(username, blockedUser string) error
}
type UserService interface {
	BlockUser(username, blockedUser string) error
}
type User struct {
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Messages     []Message `json:"messages"`
	BlockedUsers []string  `json:"blockedUsers"`
}
