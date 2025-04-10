package main

import (
	"github.com/gorilla/mux"
	"humaversion/service"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	service.Service(router)

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
