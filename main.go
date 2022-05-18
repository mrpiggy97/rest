package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/mrpiggy97/rest/handlers"
	"github.com/mrpiggy97/rest/server"
)

func BindRoutes(appServer server.IServer, router *mux.Router) {
	router.HandleFunc("/", handlers.HomeHandler(appServer)).Methods(http.MethodGet)
	router.HandleFunc("/signup", handlers.SignUpHandler(appServer)).Methods(http.MethodPost)
	router.HandleFunc("/login", handlers.LoginHandler(appServer)).Methods(http.MethodPost)
	router.HandleFunc("/posts", handlers.InsertPostHandler(appServer)).Methods(http.MethodPost)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	PORT := os.Getenv("PORT")
	DATABASEURL := os.Getenv("DATABASE_URL")
	JWTSECRET := os.Getenv("JWT_SECRET")
	var serverConfig *server.Config = &server.Config{
		Port:        PORT,
		DatabaseUrl: DATABASEURL,
		JWTSecret:   JWTSECRET,
	}
	app, serverErr := server.NewServer(context.Background(), serverConfig)
	if serverErr != nil {
		log.Fatal(serverErr.Error())
	}

	app.Start(BindRoutes)
}
