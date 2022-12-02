package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"

	_ "github.com/variant64/server/docs"
	"github.com/variant64/server/pkg/api"
)

//	@title		Variant64 Server
//	@version	1.0

// @host		localhost:8000
// @BasePath	/api
func main() {
	r := mux.NewRouter()

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	api.AttachRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
