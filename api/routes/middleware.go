package routes

import "net/http"

type Middleware func(http.Handler) http.Handler

type Mstack []Middleware

func CreateMStack(middlewares ...Middleware) Mstack {
	var stack Mstack
	return append(stack, middlewares...)
}

func (s Mstack) Then(originalHandler http.Handler) http.Handler {
	if originalHandler == nil {
		originalHandler = http.DefaultServeMux
	}

	for i := range s {
		originalHandler = s[len(s)-i-1](originalHandler)
	}
	return originalHandler
}
