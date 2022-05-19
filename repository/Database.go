package repository

import (
	"context"

	"github.com/mrpiggy97/rest/models"
)

type IDatabase interface {
	InsertUser(cxt context.Context, user *models.User) error
	GetUserById(cxt context.Context, id string) (*models.User, error)
	GetUserByEmail(cxt context.Context, email string) (*models.User, error)
	InsertPost(cxt context.Context, post *models.Post) error
	GetPostById(cxt context.Context, id string) (*models.Post, error)
	UpdatePost(cxt context.Context, post *models.Post) error
	DeletePost(cxt context.Context, post *models.Post) error
	Close()
}

var implementation IDatabase

func SetDatabase(repository IDatabase) {
	implementation = repository
}

func InsertUser(cxt context.Context, user *models.User) error {
	return implementation.InsertUser(cxt, user)
}

func GetUserById(cxt context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(cxt, id)
}

func GetUserByEmail(cxt context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(cxt, email)
}

func InsertPost(cxt context.Context, post *models.Post) error {
	return implementation.InsertPost(cxt, post)
}

func GetPostById(cxt context.Context, id string) (*models.Post, error) {
	return implementation.GetPostById(cxt, id)
}

func UpdatePost(cxt context.Context, post *models.Post) error {
	return implementation.UpdatePost(cxt, post)
}

func DeletePost(cxt context.Context, post *models.Post) error {
	return implementation.DeletePost(cxt, post)
}
