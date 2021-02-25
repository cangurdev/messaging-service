package models

type Message struct {
	FromUser string `json:"fromUser"`
	ToUser   string `json:"toUser"`
	Msg      string `json:"msg"`
	Date     string `json:"date"`
}
type User struct {
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Messages     []Message `json:"messages"`
	BlockedUsers []string  `json:"blockedUsers"`
}
type Response struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
	Method     string `json:"method"`
	Name       string `json:"name"`
}
type Block struct {
	Username    string
	BlockedUser string
}
