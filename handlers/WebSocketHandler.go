package handlers

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mrpiggy97/rest/repository"
	"github.com/mrpiggy97/rest/websockets"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(writer http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(writer, req, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var newClient *websockets.Client = websockets.NewClient(repository.AppHub, socket)
	repository.AppHub.Register <- newClient
	go newClient.Write()
}
