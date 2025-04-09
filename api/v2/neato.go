package v2

import (
	"context"
	"net/http"
)

type Handler struct{}

type GetNeatoHandlerResponse struct {
	Status int
	Body   struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
}

func (h Handler) GetNeatoHandler(_ context.Context, _ *struct{}) (*GetNeatoHandlerResponse, error) {
	return &GetNeatoHandlerResponse{
		Status: http.StatusOK,
		Body: struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}{
			Message: "neato v2",
			Code:    999,
		},
	}, nil
}
