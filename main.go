package main

import (
	"github.com/mrpiggy97/rest/server"
)

func main() {
	server.Runserver(server.GetConfig())
}
