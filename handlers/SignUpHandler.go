package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mrpiggy97/rest/models"
	"github.com/mrpiggy97/rest/repository"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Token          string `json:"token"`
	ExpirationDate *jwt.NumericDate
}

func SignUpHandler(writer http.ResponseWriter, req *http.Request) {
	var request *SignUpRequest = new(SignUpRequest)
	var expirationDate *jwt.NumericDate = new(jwt.NumericDate)
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
			return
		} else {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	// if user created successfuly then we create token
	// and send it
	expirationDate = jwt.NewNumericDate(time.Now().Add(2 * time.Hour * 24))
	var newClaims models.AppClaims = models.AppClaims{
		UserId: user.Id,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationDate,
		},
	}
	var token *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	signedToken, signingErr := token.SignedString([]byte(repository.GetJWTSecret()))
	if signingErr != nil {
		http.Error(writer, signingErr.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Add("Content-type", "application/json")
	json.NewEncoder(writer).Encode(SignUpResponse{
		Token:          signedToken,
		ExpirationDate: expirationDate,
	})
}
