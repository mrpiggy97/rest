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
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(writer http.ResponseWriter, req *http.Request) {
	// decode request
	var request *LoginRequest = new(LoginRequest)
	var user *models.User = new(models.User)
	var tokenString string = ""
	var status int

	// decode request
	err := json.NewDecoder(req.Body).Decode(request)
	if err != nil {
		fmt.Println(err.Error())
		err = errors.New("error decoding request")
		status = http.StatusBadRequest
	}
	// get user from database
	if err == nil {
		user, err = repository.GetUserByEmail(context.Background(), request.Email)
		if err != nil {
			fmt.Println(err.Error())
			err = errors.New("failed to retrieve user with that password and email")
		}
		status = http.StatusInternalServerError
	}
	// check that password sent by
	// client is password stored at db
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
		if err != nil {
			fmt.Println(err.Error())
			err = errors.New("error with password")
			status = http.StatusBadRequest
		}
	}
	// create and sign jwt token for client
	if err == nil {
		// create jwt token and sign it
		var claims models.AppClaims = models.AppClaims{
			UserId: user.Id,
			Email:  user.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(2 * time.Hour * 24).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err = token.SignedString([]byte(repository.GetJWTSecret()))
		if err != nil {
			fmt.Println(err.Error())
			err = errors.New("error signing token")
			status = http.StatusInternalServerError
		}
	}
	// send response to client
	if err == nil {
		// finalize response to requests
		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusAccepted)
		json.NewEncoder(writer).Encode(LoginResponse{
			Token: tokenString,
		})
	} else {
		http.Error(writer, err.Error(), status)
	}
}
