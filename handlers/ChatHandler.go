package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrpiggy97/rest/repository"
)

type ChatHandlerRequest struct {
	Message string `json:"message"`
}

func ChatHandler(writer http.ResponseWriter, req *http.Request) {
	var message *ChatHandlerRequest = new(ChatHandlerRequest)
	json.NewDecoder(req.Body).Decode(message)
	repository.AppHub.BroadCast(message, nil)
}
