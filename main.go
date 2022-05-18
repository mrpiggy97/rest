package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mrpiggy97/rest/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	server.Runserver()
}
