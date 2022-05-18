package server

import "github.com/mrpiggy97/rest/middleware"

func GetMiddlewareFuncs() []middleware.MiddlewareFunc {
	// middleware.RequestIsAuthenticatedMiddleware should
	// be the first middleware of all
	return []middleware.MiddlewareFunc{
		middleware.RequestIsAuthenticatedMiddleware,
		middleware.AuthorizationMiddleware,
	}
}
