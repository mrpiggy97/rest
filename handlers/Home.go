package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrpiggy97/rest/models"
	"github.com/mrpiggy97/rest/repository"
	"github.com/mrpiggy97/rest/utils"
)

type HomeResponse struct {
	Message           string            `json:"message"`
	Status            bool              `json:"status"`
	UserAuthenticated bool              `json:"userAuthenticated"`
	User              *models.AppClaims `json:"user"`
}

func HomeHandler(writer http.ResponseWriter, req *http.Request) {
	user, isAuthenticated := utils.GetUserFromRequest(req)
	writer.Header().Add("Content-type", "application/json")
	var HomeMessage models.WebsocketMessage = models.WebsocketMessage{
		Type:    "home_message",
		Payload: "this is the message",
	}
	repository.AppHub.BroadCast(HomeMessage, nil)
	json.NewEncoder(writer).Encode(HomeResponse{
		Message:           "welcome to the api",
		Status:            true,
		UserAuthenticated: isAuthenticated,
		User:              user,
	})
}
