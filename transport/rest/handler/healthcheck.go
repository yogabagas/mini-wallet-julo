package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yogabagas/jatis-BE/domain/service"
)

func (h *HandlerImpl) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := service.HealthCheckResponse{
		Status: "OK",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
