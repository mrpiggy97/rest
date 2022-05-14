package server

import (
	"net/http"
)

type MiddlewareFunc func(appServer IServer, writer http.ResponseWriter, req *http.Request) MiddlewareResponse
