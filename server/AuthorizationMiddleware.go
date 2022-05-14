package server

import (
	"errors"
	"net/http"
)

// AuthoriationMiddleware will check if user has permission to perform action
func AuthorizationMiddleware(appServer IServer, writer http.ResponseWriter, req *http.Request) MiddlewareResponse {
	// if request method is different than Get verify if user
	// is authenticated, if so do not return error
	// only login and signup routes can use POST method and not have its
	// user authenticated
	if req.Method != "GET" {
		if req.URL.Path == "login" || req.URL.Path == "signup" {
			return MiddlewareResponse{Request: req, Err: nil, StatusCode: 0}
		}
		var isAuthenticated Key = "isAuthenticated"
		switch req.Context().Value(isAuthenticated) {
		case true:
			return MiddlewareResponse{Request: req, Err: nil, StatusCode: 0}
		case false:
			status, err := http.StatusUnauthorized, errors.New("request not authorized to perform action")
			return MiddlewareResponse{Request: nil, Err: err, StatusCode: status}
		default:
			status, err := http.StatusUnauthorized, errors.New("request not authorized to perform action")
			return MiddlewareResponse{Request: nil, Err: err, StatusCode: status}
		}
	} else {
		return MiddlewareResponse{Request: req, Err: nil, StatusCode: 0}
	}
}
