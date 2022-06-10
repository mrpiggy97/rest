package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrpiggy97/rest/repository"
	"github.com/segmentio/ksuid"
)

type ChatHandlerRequest struct {
	Message string `json:"message"`
}

type ChatHandlerResponse struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
}

func ChatHandler(writer http.ResponseWriter, req *http.Request) {
	var request *ChatHandlerRequest = new(ChatHandlerRequest)
	json.NewDecoder(req.Body).Decode(request)
	var uuid string = ksuid.New().String()
	var message *ChatHandlerResponse = &ChatHandlerResponse{
		Message: request.Message,
		UUID:    uuid,
	}
	repository.AppHub.BroadCast(message, "")
}
