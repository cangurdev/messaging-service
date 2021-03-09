package handler

import (
	"cvngur/messaging-service/app/repositories"
	"cvngur/messaging-service/app/services"
	"cvngur/messaging-service/domain"
	"encoding/json"
	"net/http"
)

var mService = services.NewMessageService(repositories.NewMessageRepository())

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var m domain.Message
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		response := Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}

	err = mService.SendMessage(m.FromUser, m.ToUser, m.Msg, m.Date)

	if err != nil {
		response := Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
func ViewMessages(w http.ResponseWriter, r *http.Request) {
	var u domain.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response := Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
		errorRespond(w, response)
		return
	}
	messages, err := mService.ViewMessages(u.Username)

	if err != nil {
		response := Response{StatusCode: http.StatusBadRequest, Msg: "Hata", Method: r.Method, Name: err.Error()}
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
