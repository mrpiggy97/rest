package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrpiggy97/rest/models"
	"github.com/mrpiggy97/rest/repository"
	"github.com/mrpiggy97/rest/utils"
	"github.com/segmentio/ksuid"
)

type InsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type InsertPostRespone struct {
	Message string `json:"message"`
}

func InsertPostHandler(writer http.ResponseWriter, req *http.Request) {
	var decodedBody *InsertPostRequest = new(InsertPostRequest)
	requestUser, requestAuthenticated := utils.GetUserFromRequest(req)
	var newPost *models.Post = new(models.Post)
	var err error
	var statusCode int

	if !requestAuthenticated {
		err = errors.New("user not authorized to perform action")
		statusCode = http.StatusUnauthorized
	}
	if err == nil {
		err = json.NewDecoder(req.Body).Decode(decodedBody)
		statusCode = http.StatusBadRequest
	}
	if err == nil {
		var postId string = ksuid.New().String()
		newPost.Id = postId
		newPost.PostContent = decodedBody.PostContent
		newPost.UserId = requestUser.UserId
		err = repository.InsertPost(context.Background(), newPost)
		statusCode = http.StatusInternalServerError
	}
	if err == nil {
		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		var response InsertPostRespone = InsertPostRespone{
			Message: "successfuly created post",
		}
		json.NewEncoder(writer).Encode(response)
	} else {
		http.Error(writer, err.Error(), statusCode)
	}
}

func GetPostById(writer http.ResponseWriter, req *http.Request) {
	var err error
	var post *models.Post = new(models.Post)
	params := mux.Vars(req)
	post, err = repository.GetPostById(context.Background(), params["id"])
	if err == nil {
		writer.Header().Add("Content-type", "application/json")
		json.NewEncoder(writer).Encode(post)
	} else {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
