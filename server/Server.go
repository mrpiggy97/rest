package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpiggy97/rest/middleware"
)

type IServer interface {
	ServeHTTP(writer http.ResponseWriter, req *http.Request)
}

type Server struct {
	Config          *Config
	Router          *mux.Router
	MiddlewareFuncs []middleware.MiddlewareFunc
}

func (server *Server) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	var middlewareResponse middleware.MiddlewareResponse
	for _, middlewareFunc := range server.MiddlewareFuncs {
		middlewareResponse = middlewareFunc(writer, req)
		req = middlewareResponse.Request
		if middlewareResponse.Err != nil {
			break
		}
	}
	if middlewareResponse.Err != nil {
		http.Error(writer, middlewareResponse.Err.Error(), middlewareResponse.StatusCode)
	} else {
		server.Router.ServeHTTP(writer, req)
	}
}

func NewServer(cxt context.Context) (*Server, error) {
	var server *Server = &Server{
		Config:          GetConfig(),
		MiddlewareFuncs: GetMiddlewareFuncs(),
		Router:          GetRouter(),
	}
	if server.Config.Port == "" {
		return nil, errors.New("port must not be empty")
	}
	if server.Config.JWTSecret == "" {
		return nil, errors.New("jwtSecret must not be empty")
	}
	if server.Config.DatabaseUrl == "" {
		return nil, errors.New("databaseUrl must not be empty")
	}
	return server, nil
}
