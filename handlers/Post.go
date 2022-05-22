package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mrpiggy97/rest/models"
	"github.com/mrpiggy97/rest/repository"
	"github.com/mrpiggy97/rest/utils"
	"github.com/segmentio/ksuid"
)

type InsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type InsertPostResponse struct {
	Message string `json:"message"`
}

func InsertPostHandler(writer http.ResponseWriter, req *http.Request) {
	var decodedBody *InsertPostRequest = new(InsertPostRequest)
	requestUser, _ := utils.GetUserFromRequest(req)
	var newPost *models.Post = new(models.Post)
	var err error
	var statusCode int

	// decode request
	if err == nil {
		err = json.NewDecoder(req.Body).Decode(decodedBody)
		statusCode = http.StatusBadRequest
	}
	// insert post into db
	if err == nil {
		var postId string = ksuid.New().String()
		newPost.Id = postId
		newPost.PostContent = decodedBody.PostContent
		newPost.UserId = requestUser.UserId
		err = repository.InsertPost(context.Background(), newPost)
		statusCode = http.StatusInternalServerError
	}
	//send response to client
	if err == nil {
		var postMessage models.WebsocketMessage = models.WebsocketMessage{
			Type:    "post_created",
			Payload: newPost,
		}
		repository.AppHub.BroadCast(postMessage, nil)
		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		var response InsertPostResponse = InsertPostResponse{
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

func UpdatePost(writer http.ResponseWriter, req *http.Request) {
	var decodedBody *InsertPostRequest = new(InsertPostRequest)
	requestUser, _ := utils.GetUserFromRequest(req)
	var post *models.Post = new(models.Post)
	var err error
	var statusCode int
	var params map[string]string = mux.Vars(req)
	var postId string = params["id"]

	// decode request
	if err == nil {
		err = json.NewDecoder(req.Body).Decode(decodedBody)
		if err != nil {
			statusCode = http.StatusBadRequest
		}
	}
	// get post to edit
	if err == nil {
		post, err = repository.GetPostById(context.Background(), postId)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	}
	// verify post is owned by user request
	if err == nil {
		if requestUser.UserId != post.UserId {
			err = errors.New("unauthorized to perform action on object")
			statusCode = http.StatusUnauthorized
		}
	}
	//update post
	if err == nil {
		post.PostContent = decodedBody.PostContent
		err = repository.UpdatePost(context.Background(), post)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	}
	// send response to client
	if err == nil {
		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		var response InsertPostResponse = InsertPostResponse{
			Message: "successfuly updated post post",
		}
		json.NewEncoder(writer).Encode(response)
	} else {
		http.Error(writer, err.Error(), statusCode)
	}
}

func DeletePost(writer http.ResponseWriter, req *http.Request) {
	var err error
	var post *models.Post = new(models.Post)
	var statusCode int
	requestUser, _ := utils.GetUserFromRequest(req)
	var params map[string]string = mux.Vars(req)
	var postId string = params["id"]

	// get post to delete
	post, err = repository.GetPostById(context.Background(), postId)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	// check that post is owned by request user
	if err == nil {
		if post.UserId != requestUser.UserId {
			err = errors.New("unauthorized to perform action on object")
			statusCode = http.StatusUnauthorized
		}
	}

	// delete post
	if err == nil {
		err = repository.DeletePost(context.Background(), post)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	}

	// send response to client
	if err == nil {
		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		var response map[string]string = make(map[string]string)
		response["message"] = "successfuly deleted post"
		json.NewEncoder(writer).Encode(response)
	} else {
		http.Error(writer, err.Error(), statusCode)
	}
}

func ListPostHandler(writer http.ResponseWriter, req *http.Request) {
	var err error
	var posts []*models.Post = []*models.Post{}
	var pageString string = req.URL.Query().Get("page")
	var page int
	var statusCode int

	if pageString == "" {
		err = errors.New("page query is empty")
		statusCode = http.StatusBadRequest
	}

	if err == nil {
		page, err = strconv.Atoi(pageString)
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	}

	if err == nil {
		posts, err = repository.ListPost(context.Background(), uint64(page))
		if err != nil {
			statusCode = http.StatusInternalServerError
		}
	}

	if err == nil {
		writer.Header().Add("Content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode(posts)
	} else {
		http.Error(writer, err.Error(), statusCode)
	}
}
