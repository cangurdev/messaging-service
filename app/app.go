package app

import (
	"cvngur/messaging-service/db"
	"cvngur/messaging-service/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *mux.Router
	Db     *mongo.Collection
}

func (a *App) Initialize() {
	a.Db = db.Connection().Collection("User")
	a.Router = mux.NewRouter()
	routes.SetupRoutes(a.Router)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
