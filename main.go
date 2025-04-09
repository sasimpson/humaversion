package main

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humamux"
	"github.com/gorilla/mux"
	v1 "humaversion/api/v1"
	v2 "humaversion/api/v2"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	v1Router := router.PathPrefix("/api/v1").Subrouter()
	v1config := huma.DefaultConfig("Neato API", "0.0.1")
	v1config.Servers = []*huma.Server{
		{
			URL: "http://localhost:5000/api/v1",
		},
	}
	v1api := humamux.New(v1Router, v1config)

	v1h := &v1.Handler{}
	huma.AutoRegister(v1api, v1h)

	v2Router := router.PathPrefix("/api/v2").Subrouter()
	v2config := huma.DefaultConfig("Neato API", "0.0.2")
	v2config.Servers = []*huma.Server{
		{
			URL: "http://localhost:5000/api/v2",
		},
	}
	v2api := humamux.New(v2Router, v2config)

	v2h := &v2.Handler{}
	huma.AutoRegister(v2api, v2h)

	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
