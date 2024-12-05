package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func ApplyCors(a http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://test.com", "https://test.com:8080"},
		AllowCredentials: true,
		Debug:            true,
	})
	return c.Handler(a)
}
