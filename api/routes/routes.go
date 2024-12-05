package routes

import (
	"net/http"

	"github.com/heisenberg8055/capi/api/middleware"
	"github.com/heisenberg8055/capi/api/routes/handlers"
)

func Routes() *http.ServeMux {
	router := http.NewServeMux()

	stack := CreateMStack(middleware.LogInfo, middleware.ApplyCors).Then(router)

	router.HandleFunc("POST /add", handlers.Add)

	router.HandleFunc("POST /subtract", handlers.Subtract)

	router.HandleFunc("POST /multiply", handlers.Multiply)

	router.HandleFunc("POST /divide", handlers.Divide)

	return router
}
