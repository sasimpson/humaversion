package v2

import (
	"context"
	"github.com/google/uuid"
	v2 "humaversion/models/v2"
	"net/http"
)

// Handler for the v2 api
type Handler struct{}

// GetNeatoHandlerResponse is the response struct for our handler.
type GetNeatoHandlerResponse struct {
	Status int
	Body   v2.Neat
}

// GetNeatoHandler is a handler which demonstrates how huma handlers work
func (h Handler) GetNeatoHandler(_ context.Context, _ *struct{}) (*GetNeatoHandlerResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return &GetNeatoHandlerResponse{
		Status: http.StatusOK,
		Body: v2.Neat{
			Id:     id.String(),
			Name:   "Neato API v2!",
			MD5Sum: uuid.NewMD5(id, []byte("neato v2")).String(),
		},
	}, nil
}
