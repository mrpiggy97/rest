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
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (broker *Broker) Config() *Config {
	return broker.config
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
	if err := http.ListenAndServe(broker.config.Port, broker.router); err != nil {
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
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}
