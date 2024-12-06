package routes

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateMStack(middlewares ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			h = middleware(h)
		}
		return h
	}
}
