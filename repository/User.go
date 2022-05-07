package repository

import (
	"context"

	"github.com/mrpiggy97/rest/models"
)

type IUserRepository interface {
	InsertUser(cxt context.Context, user *models.User)
	GetUserById(cxt context.Context, id int64) (*models.User, error)
}

var implementation IUserRepository

func SetRepository(repository IUserRepository) {
	implementation = repository
}

func InsertUser(cxt context.Context, user *models.User) {
	implementation.InsertUser(cxt, user)
}

func GetUserById(cxt context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(cxt, id)
}
