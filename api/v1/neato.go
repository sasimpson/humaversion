package v1

import (
	"context"
	"github.com/google/uuid"
	v1 "humaversion/models/v1"
	"net/http"
)

// Handler for the v1 api
type Handler struct{}

// GetNeatoHandlerResponse is the response struct for our handler.
type GetNeatoHandlerResponse struct {
	Status int
	Body   v1.Neat
}

// GetNeatoHandler is a handler which demonstrates how huma handlers work
func (h Handler) GetNeatoHandler(_ context.Context, _ *struct{}) (*GetNeatoHandlerResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	return &GetNeatoHandlerResponse{
		Status: http.StatusOK,
		Body: v1.Neat{
			Name:    "Neato API",
			Id:      id.String(),
			Version: "1",
		},
	}, nil
}
