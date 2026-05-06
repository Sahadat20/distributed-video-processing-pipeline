package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Sahadat20/distributed-video-processing-pipeline.git/producer/internal/service"
)

type Handler struct {
	service *service.TaskService
}

func NewHandler(s *service.TaskService) *Handler {
	return &Handler{
		service: s,
	}
}
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "Hello World from Producer Service 🚀",
	}

	json.NewEncoder(w).Encode(response)
}
