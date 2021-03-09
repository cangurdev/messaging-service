package handler

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
