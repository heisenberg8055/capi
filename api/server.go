package api

import (
	"net/http"
)

func Start(mux http.Handler) {
	server := http.Server{
		Addr: ":8080",
	}
	server.ListenAndServe()
}
