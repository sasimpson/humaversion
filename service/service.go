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
	config.ApiConfig.Servers = config.Servers
	serviceApi := humamux.New(serviceRouter, config.ApiConfig)
	h := config.Handler
	slog.Info("initializing service api", "path", config.Path)
	huma.AutoRegister(serviceApi, h)
}

// Service sets up sub routers for each version of the api, and sets up those apis
// the key is using the Servers config, without that, the doc links do not work and it's not great
func Service(router *mux.Router) {
	AddService(router, Config{
		Path:      "/api/v1",
		ApiConfig: huma.DefaultConfig("Neato API", "0.0.1"),
		Servers: []*huma.Server{
			{
				URL:         "http://localhost:5000/api/v1",
				Description: "local",
			},
		},
		Handler: v1.Handler{},
	})
	AddService(router, Config{
		Path:      "/api/v2",
		ApiConfig: huma.DefaultConfig("Neato API", "0.0.2"),
		Servers: []*huma.Server{
			{
				URL:         "http://localhost:5000/api/v2",
				Description: "local",
			},
		},
		Handler: v2.Handler{},
	})
}
