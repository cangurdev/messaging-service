package handler

import (
	"cvngur/messaging-service/app/repositories"
	"cvngur/messaging-service/app/services"
	"cvngur/messaging-service/domain"
	"encoding/json"
	"net/http"
)

var aService = services.NewAuthService(repositories.NewAuthRepository())

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var u domain.User
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

	respond(w, http.StatusCreated)
	return
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var u domain.User
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

	respond(w, http.StatusOK)
	return
}
