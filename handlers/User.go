package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mrpiggy97/rest/models"
	"github.com/mrpiggy97/rest/repository"
	"github.com/mrpiggy97/rest/server"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

type SignUpLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func SignUpHandler(appServer server.IServer) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		var request *SignUpLoginRequest = new(SignUpLoginRequest)
		err := json.NewDecoder(req.Body).Decode(request)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		//encrypt password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 8)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		var user *models.User = &models.User{
			Id:       id.String(),
			Email:    request.Email,
			Password: string(hashedPassword),
		}
		//create user in database
		err = repository.InsertUser(req.Context(), user)
		if err != nil {
			fmt.Println("User handler ", err.Error())
			if err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"` {
				var newErr error = errors.New("user already exists")
				http.Error(writer, newErr.Error(), http.StatusInternalServerError)
			} else {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
		}
		writer.Header().Add("Content-type", "application/json")
		json.NewEncoder(writer).Encode(SignUpResponse{
			Id:    user.Id,
			Email: user.Email,
		})
	}
}

func LoginHandler(appServer server.IServer) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		// decode request
		var request *SignUpLoginRequest = new(SignUpLoginRequest)
		decodingErr := json.NewDecoder(req.Body).Decode(request)
		if decodingErr != nil {
			fmt.Println(decodingErr.Error())
			responseError := "error decoding request"
			http.Error(writer, errors.New(responseError).Error(), http.StatusBadRequest)
			return
		}
		//make query to database
		user, err := repository.GetUserByEmail(context.Background(), request.Email)
		if err != nil {
			fmt.Println(err.Error())
			responseError := "there was an error retrieving user with that email"
			http.Error(writer, errors.New(responseError).Error(), http.StatusInternalServerError)
			return
		}
		// decrypt password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			http.Error(writer, errors.New("error with password").Error(), http.StatusUnauthorized)
			return
		}
		// create jwt token and sign it
		var claims models.AppClaims = models.AppClaims{
			UserId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(appServer.Config().JWTSecret))
		if err != nil {
			fmt.Println(err.Error())
			responseError := "error signing token"
			http.Error(writer, errors.New(responseError).Error(), http.StatusInternalServerError)
			return
		}
		// finalize response to requests
		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusAccepted)
		json.NewEncoder(writer).Encode(LoginResponse{
			Token: tokenString,
		})
	}
}
