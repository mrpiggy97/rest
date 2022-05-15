package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrpiggy97/rest/models"
	"github.com/mrpiggy97/rest/server"
	"github.com/mrpiggy97/rest/utils"
)

type HomeResponse struct {
	Message           string            `json:"message"`
	Status            bool              `json:"status"`
	UserAuthenticated bool              `json:"userAuthenticated"`
	User              *models.AppClaims `json:"user"`
}

func HomeHandler(s server.IServer) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		user, isAuthenticated := utils.GetUserFromRequest(req)
		writer.Header().Add("Content-type", "application/json")
		json.NewEncoder(writer).Encode(HomeResponse{
			Message:           "welcome to the api",
			Status:            true,
			UserAuthenticated: isAuthenticated,
			User:              user,
		})
	}
}
