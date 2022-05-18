package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpiggy97/rest/handlers"
)

func GetRouter() *mux.Router {
	var serverRouter *mux.Router = mux.NewRouter()
	serverRouter.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)
	serverRouter.HandleFunc("/signup", handlers.SignUpHandler).Methods(http.MethodPost)
	serverRouter.HandleFunc("/login", handlers.LoginHandler).Methods(http.MethodPost)
	serverRouter.HandleFunc("/posts", handlers.InsertPostHandler).Methods(http.MethodPost)
	return serverRouter
}