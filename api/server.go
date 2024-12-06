package api

import (
	"net/http"

	_ "github.com/heisenberg8055/capi/api/routes"
)

func Start(mux http.Handler) {

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
