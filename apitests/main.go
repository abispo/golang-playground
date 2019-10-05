package main

import (
	"fmt"
	"golang-playground/apitests/handlers"
	"golang-playground/apitests/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var apiVersion = "v1"

func main() {
	r := mux.NewRouter()
	r.Use(middleware.TokenMiddleware)

	r.HandleFunc(fmt.Sprintf("/%s/message", apiVersion), handlers.OkHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
