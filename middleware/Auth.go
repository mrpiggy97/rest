package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mrpiggy97/rest/models"
	"github.com/mrpiggy97/rest/server"
)

var NO_AUTH_NEEDED []string = []string{"login", "signup"}

func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(appServer server.IServer) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			if !shouldCheckToken(req.URL.Path) {
				next.ServeHTTP(writer, req)
				return
			} else {
				tokenString := strings.TrimSpace(req.Header.Get("Authorization"))
				_, err := jwt.ParseWithClaims(tokenString, models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
					return []byte(appServer.Config().JWTSecret), nil
				})
				if err != nil {
					http.Error(writer, err.Error(), http.StatusUnauthorized)
					return
				}
				next.ServeHTTP(writer, req)
			}
		})
	}
}
