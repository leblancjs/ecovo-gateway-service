package main

import (
	"log"
	"net/http"
	"os"

	"azure.com/ecovo/gateway-service/cmd/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()

	r.Handle("/", handler.HelloWorld()).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r)))
}
