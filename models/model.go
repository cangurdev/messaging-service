package models

type Message struct {
	FromUser string `json:"fromUser"`
	ToUser   string `json:"toUser"`
	Msg      string `json:"msg"`
	Date     string `json:"date"`
}
