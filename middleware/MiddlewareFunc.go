package middleware

import (
	"net/http"
)

type MiddlewareFunc func(writer http.ResponseWriter, req *http.Request) MiddlewareResponse
