package utils

import (
	"net/http"

	"github.com/mrpiggy97/rest/models"
	"github.com/mrpiggy97/rest/server"
)

// GetUserFromRequest will return user from request context
func GetUserFromRequest(req *http.Request) (*models.AppClaims, bool) {
	var isAuthenticatedKey server.Key = "isAuthenticated"
	var isAuthenticated interface{} = req.Context().Value(isAuthenticatedKey)
	var authenticated bool = isAuthenticated.(bool)
	if authenticated {
		var userKey server.Key = "user"
		var user interface{} = req.Context().Value(userKey)
		var userClaims *models.AppClaims = user.(*models.AppClaims)
		return userClaims, true
	}
	return nil, false
}
