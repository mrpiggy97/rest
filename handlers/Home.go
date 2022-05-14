package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrpiggy97/rest/server"
)

type HomeResponse struct {
	Message           string `json:"message"`
	Status            bool   `json:"status"`
	UserAuthenticated bool   `json:"userAuthenticated"`
}

func HomeHandler(s server.IServer) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Add("Content-type", "application/json")
		var key server.Key = "isAuthenticated"
		var value interface{} = req.Context().Value(key)
		var boolValue bool = value.(bool)
		json.NewEncoder(writer).Encode(HomeResponse{
			Message:           "welcome to the api",
			Status:            true,
			UserAuthenticated: boolValue,
		})
	}
}
