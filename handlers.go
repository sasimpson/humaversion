package main

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
	"net/http"
)

type v1Handler struct{}

type v1GetHandlerResponse struct {
	Message string `json:"message"`
}

func (h v1Handler) v1GetHandler(_ context.Context, _ *struct{}) (*v1GetHandlerResponse, error) {
	return &v1GetHandlerResponse{Message: "neato v1"}, nil
}

func (h v1Handler) RegisterRoutes(api huma.API) {
	// GET /v1/neato
	huma.Register(api, huma.Operation{
		OperationID: "GetNeato",
		Method:      http.MethodGet,
		Path:        "/neato",
		Tags:        []string{"v1"},
	}, h.v1GetHandler)
}

type v2Handler struct{}

type v2GetHandlerResponse struct {
	Message string `json:"message"`
}

func (h v2Handler) v2GetHandler(_ context.Context, _ *struct{}) (*v2GetHandlerResponse, error) {
	return &v2GetHandlerResponse{Message: "neato v2"}, nil
}

func (h v2Handler) RegisterRoutes(api huma.API) {
	// GET /neato
	huma.Register(api, huma.Operation{
		OperationID: "GetNeato",
		Method:      "GET",
		Path:        "/neato",
		Tags:        []string{"v2"},
	}, h.v2GetHandler)
}
