package v1

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"humaversion/models/v1"
	"net/http"
	"testing"
)

func TestHandler_GetV1NeatoHandler(t *testing.T) {
	type args struct {
		in0 context.Context
		in1 *struct{}
	}
	tests := []struct {
		name    string
		args    args
		want    *GetNeatoHandlerResponse
		wantErr bool
	}{
		{
			name: "Valid Test",
			args: args{
				in0: context.Background(),
				in1: nil,
			},
			want: &GetNeatoHandlerResponse{
				Status: http.StatusOK,
				Body: v1.Neat{
					Id:      "",
					Version: "1",
					Name:    "Neato API",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{}
			got, err := h.GetNeatoHandler(tt.args.in0, tt.args.in1)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.Status, got.Status)
			assert.Equal(t, tt.want.Body.Name, got.Body.Name)
			assert.Equal(t, tt.want.Body.Version, got.Body.Version)
			assert.NotEmpty(t, got.Body.Id)
			id, err := uuid.Parse(got.Body.Id)
			assert.Nil(t, err)
			assert.IsType(t, uuid.UUID{}, id)
		})
	}
}
