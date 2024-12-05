package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func LogInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		next.ServeHTTP(w, r)
		slog.SetDefault(logger)
		slog.LogAttrs(
			context.Background(),
			slog.LevelInfo,
			"API Call Method: "+r.Method,
			slog.String("Status Code", "200"),
			slog.String("IP", r.RemoteAddr),
			slog.String("Endpoint", r.RequestURI),
			slog.String("Time to process:", time.Since(start).String()),
		)
	})
}
