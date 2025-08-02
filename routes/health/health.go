package health

import (
	"encoding/json"
	"net/http"
	
	"github.com/gorilla/mux"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "healthy",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RegisterHealthRoutes(r *mux.Router) {
	r.HandleFunc("/health", HealthCheckHandler).Methods("GET")
}