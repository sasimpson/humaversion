package v2

import (
	"context"
	"github.com/google/uuid"
	v2 "humaversion/models/v2"
	"net/http"
)

type Handler struct{}

type GetNeatoHandlerResponse struct {
	Status int
	Body   v2.Neat
}

type GetNeatoBody struct {
	Message string `json:"message" doc:"message to the user"`
	Code    int    `json:"code" doc:"system response code"`
}

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
			MD5Sum: uuid.NewMD5(id, []byte("neato v1")).String(),
		},
	}, nil
}
