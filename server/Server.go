package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpiggy97/rest/database"
	"github.com/mrpiggy97/rest/repository"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

type IServer interface {
	Config() *Config
	ServeHTTP(writer http.ResponseWriter, req *http.Request)
}

type Broker struct {
	config          *Config
	router          *mux.Router
	middlewareFuncs []MiddlewareFunc
}

func (broker *Broker) Config() *Config {
	return broker.config
}

func (broker *Broker) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	var middlewareResponse MiddlewareResponse
	for _, middlewareFunc := range broker.middlewareFuncs {
		middlewareResponse = middlewareFunc(broker, writer, req)
		req = middlewareResponse.Request
		if middlewareResponse.Err != nil {
			break
		}
	}
	if middlewareResponse.Err != nil {
		http.Error(writer, middlewareResponse.Err.Error(), middlewareResponse.StatusCode)
	} else {
		broker.router.ServeHTTP(writer, req)
	}
}

func (broker *Broker) Start(binder func(appServer IServer, router *mux.Router)) {
	broker.router = mux.NewRouter()
	binder(broker, broker.router)
	repo, repoErr := database.NewPostgresqlRespository(broker.config.DatabaseUrl)
	if repoErr != nil {
		log.Fatal(repoErr.Error())
	}
	repository.SetRepository(repo)
	fmt.Println("starting server at port ", broker.config.Port)
	if err := http.ListenAndServe(broker.config.Port, broker); err != nil {
		log.Fatal(err.Error())
	}
}

func NewServer(cxt context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port must not be empty")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwtSecret must not be empty")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("databaseUrl must not be empty")
	}
	var broker *Broker = &Broker{
		config:          config,
		router:          mux.NewRouter(),
		middlewareFuncs: []MiddlewareFunc{RequestIsAuthenticated, AuthorizationMiddleware},
	}
	return broker, nil
}
