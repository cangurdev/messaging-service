package routes

import (
	"cvngur/messaging-service/handler"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {

	router.HandleFunc("/register", handler.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handler.LoginHandler).Methods("POST")
}
