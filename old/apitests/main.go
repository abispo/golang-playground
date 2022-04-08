package main

import (
	"fmt"
	"golang-playground/apitests/handlers"
	"golang-playground/apitests/middleware"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/gorilla/mux"
)

var apiVersion = "v1"

func main() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APITESTS_APP")

	r := mux.NewRouter()
	r.Use(middleware.TokenMiddleware)

	r.HandleFunc(fmt.Sprintf("/%s/message", apiVersion), handlers.OkHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", viper.Get("PORT")), r))
}
