package domain

type Message struct {
	FromUser string `json:"fromUser"`
	ToUser   string `json:"toUser"`
	Msg      string `json:"msg"`
	Date     string `json:"date"`
}
type MessageRepository interface {
	SendMessage(fromUser, toUser, msg, date string) error
	GetMessages(username string) ([]Message, error)
	GetBlockedUsers(username string) []string
}
type MessageService interface {
	SendMessage(fromUser, toUser, msg, date string) error
	ViewMessages(username string) ([]Message, error)
}
