package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mrpiggy97/rest/models"
	"github.com/mrpiggy97/rest/repository"
)

// RequestIsAuthenticated will pass to request context
// either 1 or 2 values: isAuthenticated which will be a bool
// and if the case is true for that value we will pass a *models.User
// request can only be considered as authenticated if token is valid,
// and can only be considered as not authenticated if no token was provided
// in case token was provided but errors ocurred in its verification we will
// send an error
func RequestIsAuthenticatedMiddleware(writer http.ResponseWriter, req *http.Request) MiddlewareResponse {

	var tokenString string = strings.TrimSpace(req.Header.Get("Authorization"))
	// no authorization token was given so request is not authenticated
	if len(tokenString) == 0 {
		var isAuthenticated Key = "isAuthenticated"
		var newContext context.Context = context.WithValue(
			req.Context(),
			isAuthenticated,
			false,
		)
		var newRequest *http.Request = req.Clone(newContext)
		return MiddlewareResponse{Request: newRequest, Err: nil, StatusCode: 0}
	}
	token, parsinErr := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(repository.GetJWTSecret()), nil
	})
	// error trying to verify token, we send error in middlewareResponse
	if parsinErr != nil {
		return MiddlewareResponse{Request: nil, Err: parsinErr, StatusCode: http.StatusBadRequest}
	}
	claims, ok := token.Claims.(*models.AppClaims)
	// another error trying to verify token, we send error
	// in middlewareResponse
	if !ok {
		return MiddlewareResponse{
			Request:    nil,
			Err:        errors.New("error with token claims"),
			StatusCode: http.StatusBadGateway,
		}
	}
	// if token is not valid it most likely expired
	// we send error in middlewareResponse
	if !token.Valid {
		return MiddlewareResponse{
			Request:    nil,
			Err:        errors.New("token has expired"),
			StatusCode: http.StatusBadRequest,
		}
	}
	// token is assumed to be valid we will proceed to
	// get user from claims
	var userKey Key = "user"
	var isAuthenticatedKey Key = "isAuthenticated"
	var newContext context.Context = context.WithValue(
		req.Context(),
		userKey,
		claims,
	)
	newContext = context.WithValue(
		newContext,
		isAuthenticatedKey,
		true,
	)
	var newRequest *http.Request = req.Clone(newContext)
	return MiddlewareResponse{
		Request:    newRequest,
		Err:        nil,
		StatusCode: 0,
	}
}
