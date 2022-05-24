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

func Runserver() {
	appServer, err := NewServer(context.Background())
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
	go repository.AppHub.Run()
	fmt.Println("starting appServer at port ", appServer.Config.Port)
	var corsOptions cors.Options = cors.Options{
		AllowedOrigins: []string{"http://localhost:8000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Content-type"},
	}
	var corsHandler *cors.Cors = cors.New(corsOptions)
	if err := http.ListenAndServe(appServer.Config.Port, corsHandler.Handler(appServer)); err != nil {
		log.Fatal(err.Error())
	}
}
