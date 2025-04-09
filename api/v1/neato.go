package v1

import (
	"context"
	"net/http"
)

type Handler struct{}

type GetNeatoHandlerResponse struct {
	Status int
	Body   GetNeatoBody
}

type GetNeatoBody struct {
	Message string `json:"message" doc:"message to the user"`
}

func (h Handler) GetNeatoHandler(_ context.Context, _ *struct{}) (*GetNeatoHandlerResponse, error) {
	return &GetNeatoHandlerResponse{
		Status: http.StatusOK,
		Body: GetNeatoBody{
			Message: "neato v1",
		},
	}, nil
}
