package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mrpiggy97/rest/database"
	"github.com/mrpiggy97/rest/repository"
)

func Runserver() {
	appServer, err := NewServer(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	db, dbErr := database.NewPostgresqlRespository(appServer.Config.DatabaseUrl)
	if dbErr != nil {
		log.Fatal(dbErr.Error())
	}
	repository.SetDatabase(db)
	repository.SetConfig(appServer.Config)
	fmt.Println("starting appServer at port ", appServer.Config.Port)
	if err := http.ListenAndServe(appServer.Config.Port, appServer); err != nil {
		log.Fatal(err.Error())
	}
}