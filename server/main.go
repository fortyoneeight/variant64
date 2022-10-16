package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/variant64/server/api"
)

func main() {
	r := mux.NewRouter()
	api.AttachRESTRoutes(r)
	api.AttachWebsocketRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
