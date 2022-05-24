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
	serverRouter.HandleFunc("/posts/{id}", handlers.GetPostById).Methods(http.MethodGet)
	serverRouter.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods(http.MethodPut)
	serverRouter.HandleFunc("/posts/{id}", handlers.DeletePost).Methods(http.MethodDelete)
	serverRouter.HandleFunc("/posts-list", handlers.ListPostHandler).Methods(http.MethodGet)
	serverRouter.HandleFunc("/ws", handlers.WebSocketHandler)
	serverRouter.HandleFunc("/chat", handlers.ChatHandler).Methods(http.MethodPost)
	return serverRouter
}
