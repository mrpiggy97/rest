package utils

import (
	"errors"
	"log"
	"net/http"

	"github.com/mrpiggy97/rest/contextTypes"
	"github.com/mrpiggy97/rest/models"
)

// GetUserFromRequest will return user from request context
func GetUserFromRequest(req *http.Request) (*models.AppClaims, bool) {
	var isAuthenticatedKey contextTypes.Key = "isAuthenticated"
	var isAuthenticated interface{} = req.Context().Value(isAuthenticatedKey)
	authenticated, typeOk := isAuthenticated.(bool)
	//type verification has failed
	if !typeOk {
		var typeError error = errors.New("error with interface type checking")
		log.Fatal(typeError.Error())
	}
	if authenticated {
		var userKey contextTypes.Key = "user"
		var user interface{} = req.Context().Value(userKey)
		var userClaims *models.AppClaims = user.(*models.AppClaims)
		return userClaims, true
	}
	return nil, false
}
