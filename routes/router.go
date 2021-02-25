package routes

import (
	"cvngur/messaging-service/handler"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {

	router.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
	router.HandleFunc("/send", handler.SendMessageHandler).Methods("POST")
	router.HandleFunc("/block", handler.BlockUserHandler).Methods("POST")
	router.HandleFunc("/view", handler.ViewMessages).Methods("GET")
}
