package service

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humamux"
	"github.com/gorilla/mux"
	v1 "humaversion/api/v1"
	v2 "humaversion/api/v2"
	"log/slog"
)

type VersionHandler interface {
	RegisterRoutes(api huma.API)
}

// AddService sets up a sub router for the given path and registers the API
func AddService(router *mux.Router, config Config) {
	serviceRouter := router.PathPrefix(config.Path).Subrouter()
	serviceApi := humamux.New(serviceRouter, config.ApiConfig)
	serviceHandler := config.Handler
	slog.Info("initializing service api", "path", config.Path)
	huma.AutoRegister(serviceApi, serviceHandler)
}

// Service sets up sub routers for each version of the api, and sets up those apis
// the key is using the Servers config, without that, the doc links do not work and it's not great
func Service(router *mux.Router) {
	v1Config := huma.DefaultConfig("Neato API", "0.0.1")
	v1Config.Servers = []*huma.Server{
		{
			URL:         "http://localhost:5000/api/v1",
			Description: "local",
		},
	}
	AddService(router, Config{
		Path:      "/api/v1",
		ApiConfig: v1Config,
		Handler:   v1.Handler{},
	})

	v2Config := huma.DefaultConfig("Neato API", "0.0.2")
	v2Config.Servers = []*huma.Server{
		{
			URL:         "http://localhost:5000/api/v2",
			Description: "local",
		},
	}
	AddService(router, Config{
		Path:      "/api/v2",
		ApiConfig: v2Config,
		Handler:   v2.Handler{},
	})
}
