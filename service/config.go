package service

import "github.com/danielgtaylor/huma/v2"

// Config for the service API
type Config struct {
	Path      string
	ApiConfig huma.Config
	Handler   VersionHandler
}
