package middleware

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		limit := rate.Every(60 * time.Second)

		limiter := rate.NewLimiter(limit, 2)

		fmt.Println(limit)

		if !limiter.Allow() {
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
