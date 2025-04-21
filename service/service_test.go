package service

import (
	"context"
	"github.com/danielgtaylor/huma/v2"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHandler struct{}

type mockResponse struct {
	Name string `json:"name"`
}

func (h mockHandler) RegisterRoutes(api huma.API) {
	huma.Register(
		api,
		huma.Operation{
			OperationID:   "GetTest",
			Method:        "GET",
			Path:          "/test",
			DefaultStatus: http.StatusOK,
		},
		func(ctx context.Context, _ *struct{}) (*mockResponse, error) {
			return &mockResponse{Name: "test"}, nil
		},
	)
}

func TestAddService(t *testing.T) {
	type args struct {
		config Config
	}
	type request struct {
		path   string
		method string
	}
	tests := []struct {
		name       string
		args       args
		req        request
		wantStatus int
	}{
		{
			name: "Test with valid config",
			args: args{
				config: Config{
					Path: "/api/v1",
					ApiConfig: huma.Config{
						OpenAPI: &huma.OpenAPI{
							Info: &huma.Info{
								Title:   "Test API",
								Version: "1.0.0",
							},
							Servers: []*huma.Server{{URL: "http://localhost:5000/api/v1"}},
						},
					},
					Handler: mockHandler{},
				},
			},
			req: request{
				path:   "/api/v1/test",
				method: http.MethodGet,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req := httptest.NewRequest(tt.req.method, tt.req.path, nil)

			router := mux.NewRouter()
			AddService(router, tt.args.config)
			router.ServeHTTP(rr, req)

			assert.Equal(t, http.StatusOK, rr.Code)

		})
	}
}
