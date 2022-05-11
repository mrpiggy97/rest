package repository

import (
	"context"

	"github.com/mrpiggy97/rest/models"
)

type IUserRepository interface {
	InsertUser(cxt context.Context, user *models.User) error
	GetUserById(cxt context.Context, id string) (*models.User, error)
	GetUserByEmail(cxt context.Context, email string) (*models.User, error)
	Close() error
}

var implementation IUserRepository

func SetRepository(repository IUserRepository) {
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
