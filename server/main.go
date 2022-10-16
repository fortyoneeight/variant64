package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, world! UPDATE PLACEHOLDER\n")
}

func main() {
	r := mux.NewRouter()
	AttachRoutes(r)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}
