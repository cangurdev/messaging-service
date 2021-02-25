package handler

import (
	"cvngur/messaging-service/repositories"
	"cvngur/messaging-service/services"
	"encoding/json"
	"net/http"
)

type User struct {
	Username string
	Password string
}
type Response struct {
	StatusCode int
	Msg        string
	Method     string
	Name       string
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		resp := Response{400, "Hata", r.Method, err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	service := services.NewUserService(repositories.NewUserRepository())

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

	service := services.NewUserService(repositories.NewUserRepository())

	err = service.Login(u.Username, u.Password)

	if err != nil {
		http.Error(w, "User Cannot Found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
