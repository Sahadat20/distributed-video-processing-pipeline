package handler

import (
	"encoding/json"
	"net/http"
)

type CreateTaskRequest struct {
	Name string `json:"name"`
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var req CreateTaskRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := h.service.CreateTask(req.Name)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Failed to publish task", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task published successfully"))
}
