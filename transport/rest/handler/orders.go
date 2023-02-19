package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/yogabagas/jatis-BE/domain/service"
)

func (h *HandlerImpl) CreateOrders(w http.ResponseWriter, r *http.Request) {
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

	var req []service.CreateOrdersReq

	resp, err := h.Controller.OrdersController.CreateOrders(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
