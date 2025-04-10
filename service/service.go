package service

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humamux"
	"github.com/gorilla/mux"
	v1 "humaversion/api/v1"
	v2 "humaversion/api/v2"
	"log/slog"
)

func Service(router *mux.Router) {

	// v1 of the api
	v1Router := router.PathPrefix("/api/v1").Subrouter()
	v1config := huma.DefaultConfig("Neato API", "0.0.1")
	v1config.Servers = []*huma.Server{
		{
			URL:         "http://localhost:5000/api/v1",
			Description: "local",
		},
	}
	v1api := humamux.New(v1Router, v1config)
	v1h := &v1.Handler{}
	slog.Info("initializing v1 api")
	huma.AutoRegister(v1api, v1h)

	// v2 of the api
	v2Router := router.PathPrefix("/api/v2").Subrouter()
	v2config := huma.DefaultConfig("Neato API", "0.0.2")
	v2config.Servers = []*huma.Server{
		{
			URL:         "http://localhost:5000/api/v2",
			Description: "local",
		},
	}
	v2api := humamux.New(v2Router, v2config)
	v2h := &v2.Handler{}
	slog.Info("initializing v2 api")
	huma.AutoRegister(v2api, v2h)

}
