package handler

import (
	"cvngur/messaging-service/models"
	"cvngur/messaging-service/repositories"
	"cvngur/messaging-service/services"
	"encoding/json"
	"net/http"
)

var service = services.NewUserService(repositories.NewUserRepository())

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		response := models.Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	err = service.Register(u.Username, u.Password)

	if err != nil {
		response := models.Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		response := models.Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	err = service.Login(u.Username, u.Password)

	if err != nil {
		response := models.Response{StatusCode: http.StatusNotFound, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}
func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var m models.Message
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		response := models.Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	err = service.SendMessage(m.FromUser, m.ToUser, m.Msg, m.Date)

	if err != nil {
		response := models.Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
func ViewMessages(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response := models.Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}
	messages, err := service.ViewMessages(u.Username)

	if err != nil {
		response := models.Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}
	w.WriteHeader(http.StatusOK)
	messagesJson, err := json.Marshal(messages)
	if err != nil {
		return
	}
	w.Write(messagesJson)
	return
}

func BlockUserHandler(w http.ResponseWriter, r *http.Request) {
	var b models.Block
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		response := models.Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}
	err = service.BlockUser(b.Username, b.BlockedUser)
	if err != nil {
		return
	}
}
func errorRespond(w http.ResponseWriter, r models.Response) {
	w.WriteHeader(r.StatusCode)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	json.NewEncoder(w).Encode(r)
}
