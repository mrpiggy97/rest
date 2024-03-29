package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mrpiggy97/rest/database"
	"github.com/mrpiggy97/rest/repository"
	"github.com/rs/cors"
)

func Runserver(config *Config) {
	appServer, err := NewServer(context.Background(), config)
	if err != nil {
		log.Fatal(err.Error())
	}
	db, dbErr := database.NewPostgresqlRepository(appServer.Config.DatabaseUrl)
	if dbErr != nil {
		log.Fatal(dbErr.Error())
	}
	repository.SetDatabase(db)
	repository.SetConfig(appServer.Config)
	repository.SetHub(appServer.Hub)
	repository.RunHub()
	fmt.Println("starting appServer at port ", appServer.Config.Port)
	var corsOptions cors.Options = cors.Options{
		AllowedOrigins: appServer.AllowedOrigins,
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-type"},
	}
	var corsHandler *cors.Cors = cors.New(corsOptions)
	var address string = fmt.Sprintf("0.0.0.0:%v", appServer.Config.Port)
	if err := http.ListenAndServe(address, corsHandler.Handler(appServer)); err != nil {
		log.Fatal(err.Error())
	}
}
