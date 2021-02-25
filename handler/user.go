package handler

import (
	"cvngur/messaging-service/repositories"
	"cvngur/messaging-service/services"
	"encoding/json"
	"net/http"
)

type User struct {
	Username     string   `json:"username"`
	Password     string   `json:"password"`
	Messages     []string `json:"messages"`
	blockedUsers []string `json:"blockedUsers"`
}
type Response struct {
	StatusCode int
	Msg        string
	Method     string
	Name       string
}
type Message struct {
	FromUser string `json:"fromUser"`
	ToUser   string `json:"toUser"`
	Msg      string `json:"msg"`
	date     string `json:"date"`
}
type Block struct {
	Username    string
	BlockedUser string
}

var service = services.NewUserService(repositories.NewUserRepository())

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		resp := Response{400, "Hata", r.Method, err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			return
		}
		return
	}

	err = service.Register(u.Username, u.Password)

	if err != nil {
		resp := Response{400, "Hata", r.Method, err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.Login(u.Username, u.Password)

	if err != nil {
		http.Error(w, "User Cannot Found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.SendMessage(m.FromUser, m.ToUser, m.Msg, m.date)

	if err != nil {
		http.Error(w, err.Error(), 400)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
func ViewMessages(w http.ResponseWriter, r *http.Request) {

}

func BlockUserHandler(w http.ResponseWriter, r *http.Request) {
	var b Block
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.BlockUser(b.Username, b.BlockedUser)
	if err != nil {
		return
	}
}
