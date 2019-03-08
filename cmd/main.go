package main

import (
	"log"
	"net/http"
	"os"

	"azure.com/ecovo/gateway-service/cmd/handler"
	"github.com/gorilla/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	userServiceHost := os.Getenv("USER_SERVICE_HOST")
	if userServiceHost == "" {
		log.Fatal("missing user service host")
	}

	tripServiceHost := os.Getenv("TRIP_SERVICE_HOST")
	if tripServiceHost == "" {
		log.Fatal("missing trip service host")
	}

	tripSearchServiceHost := os.Getenv("TRIP_SEARCH_SERVICE_HOST")
	if tripSearchServiceHost == "" {
		log.Fatal("missing trip search service host")
	}

	routes := make(map[string]string)
	routes["users"] = userServiceHost
	routes["trips"] = tripServiceHost
	routes["search"] = tripSearchServiceHost

	mux := http.NewServeMux()
	mux.Handle("/", handler.RequestID(handler.ReverseProxy(routes)))

	log.Fatal(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, mux)))
}
