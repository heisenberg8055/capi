package routes

import (
	"net/http"

	"github.com/heisenberg8055/capi/api/middleware"
	"github.com/heisenberg8055/capi/api/routes/handlers"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	stack := CreateMStack(middleware.LogInfo, middleware.RateLimit, middleware.ApplyCors)

	finalHandler := http.HandlerFunc(handlers.DecodeJSONRequest)

	router.Handle("POST /add", stack(finalHandler))

	router.Handle("POST /subtract", stack(finalHandler))

	router.Handle("POST /multiply", stack(finalHandler))

	router.Handle("POST /divide", stack(finalHandler))

	return router
}
