package openapi

import (
	_ "embed"
	"net/http"
)

// Spec contains the OpenAPI document served by the application.
//
//go:embed openapi.yaml
var Spec []byte

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/yaml")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(Spec)
}
