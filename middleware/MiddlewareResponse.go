package middleware

import "net/http"

type MiddlewareResponse struct {
	Request    *http.Request
	Err        error
	StatusCode int
}
