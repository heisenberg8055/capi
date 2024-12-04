package api

import (
	"net/http"

	"github.com/heisenberg8055/capi/api/middleware"
)

func Start(mux http.Handler) {
	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.LogInfo(mux),
	}
	server.ListenAndServe()
}
