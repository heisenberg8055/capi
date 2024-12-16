package middleware

import (
	"fmt"
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimit(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(2, 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
			return
		} else {

			fmt.Println(limiter.Tokens())
			next.ServeHTTP(w, r)
		}
	})
}
