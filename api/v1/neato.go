package v1

import (
	"context"
	"net/http"
)

type Handler struct{}

type GetNeatoHandlerResponse struct {
	Status int
	Body   struct {
		Message string `json:"message"`
	}
}

func (h Handler) GetNeatoHandler(ctx context.Context, in *struct{}) (*GetNeatoHandlerResponse, error) {
	return &GetNeatoHandlerResponse{
		Status: http.StatusOK,
		Body: struct {
			Message string `json:"message"`
		}{
			Message: "neato v1",
		},
	}, nil
}
