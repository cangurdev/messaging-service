package handler

import (
	"cvngur/messaging-service/models"
	"cvngur/messaging-service/repositories/messageRepository"
	"cvngur/messaging-service/services/messageService"
	"encoding/json"
	"net/http"
)

var mService = messageService.NewMessageService(messageRepository.NewMessageRepository())

func SendMessageHandler(w http.ResponseWriter, r *http.Request) {
	var m models.Message
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
	var u User
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
