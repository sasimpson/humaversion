package v2

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
	Code    int    `json:"code" doc:"system response code"`
}

func (h Handler) GetNeatoHandler(_ context.Context, _ *struct{}) (*GetNeatoHandlerResponse, error) {
	return &GetNeatoHandlerResponse{
		Status: http.StatusOK,
		Body: GetNeatoBody{
			Message: "neato v2",
			Code:    999,
		},
	}, nil
}
