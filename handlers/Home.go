package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrpiggy97/rest/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.IServer) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(HomeResponse{
			Message: "welcome to the api",
			Status:  true,
		})
	}
}
