package v1

import (
	"context"
	"github.com/google/uuid"
	v1 "humaversion/models/v1"
	"net/http"
)

type Handler struct{}

type GetNeatoHandlerResponse struct {
	Status int
	Body   v1.Neat
}

type GetNeatoBody struct {
	Message string `json:"message" doc:"message to the user"`
}

func (h Handler) GetNeatoHandler(ctx context.Context, _ *struct{}) (*GetNeatoHandlerResponse, error) {
	var version int
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "version", version)
	return &GetNeatoHandlerResponse{
		Status: http.StatusOK,
		Body: v1.Neat{
			Name:    "Neato API",
			Id:      id.String(),
			Version: "1",
		},
	}, nil
}
