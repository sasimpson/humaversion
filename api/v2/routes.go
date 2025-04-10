package v2

import "github.com/danielgtaylor/huma/v2"

// RegisterRoutes gets run by the huma.AutoRegister function
func (h Handler) RegisterRoutes(api huma.API) {
	// GET /neato
	huma.Register(api, huma.Operation{
		OperationID: "GetNeato",
		Method:      "GET",
		Path:        "/neato",
		Tags:        []string{"v2"},
	}, h.GetNeatoHandler)
}
