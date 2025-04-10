package v1

import (
	"github.com/danielgtaylor/huma/v2"
	"net/http"
)

// RegisterRoutes gets run by the huma.AutoRegister function
func (h Handler) RegisterRoutes(api huma.API) {
	// GET /v1/neato
	huma.Register(api, huma.Operation{
		Deprecated:  true,
		OperationID: "GetNeato",
		Method:      http.MethodGet,
		Path:        "/neato",
		Tags:        []string{"v1"},
	}, h.GetNeatoHandler)
}
