package main

import (
	"github.com/heisenberg8055/capi/api"
	"github.com/heisenberg8055/capi/api/routes"
)

func main() {

	mux := routes.Routes()
	api.Start(mux)

}
