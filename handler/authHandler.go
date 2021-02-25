package handler

import (
	"cvngur/messaging-service/repositories/authRepository"
	"cvngur/messaging-service/services/authService"
	"encoding/json"
	"net/http"
)

var aService = authService.NewAuthService(authRepository.NewUserRepository())

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		response := Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	err = aService.Register(u.Username, u.Password)

	if err != nil {
		response := Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var u User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		response := Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	err = aService.Login(u.Username, u.Password)

	if err != nil {
		response := Response{StatusCode: http.StatusNotFound, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
