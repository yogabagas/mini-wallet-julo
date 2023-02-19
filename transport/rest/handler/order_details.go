package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/yogabagas/jatis-BE/domain/service"
)

func (h *HandlerImpl) CreateOrderDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req []service.CreateOrderDetailsReq

	resp, err := h.Controller.OrderDetailsController.CreateOrderDetails(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *HandlerImpl) GetOrderDetails(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "wrong method", http.StatusMethodNotAllowed)
		return
	}

	resp, err := h.Controller.OrderDetailsController.GetOrderDetail(r.Context())
	if err != nil && resp != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if resp == nil {
		http.Error(w, "data not found", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
