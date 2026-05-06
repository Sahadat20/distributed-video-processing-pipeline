package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Sahadat20/distributed-video-processing-pipeline.git/producer/internal/handler"
	"github.com/Sahadat20/distributed-video-processing-pipeline.git/producer/internal/publisher"
	"github.com/Sahadat20/distributed-video-processing-pipeline.git/producer/internal/service"
)

func Serve() {

	pub := publisher.NewRabbitMQPublisher()
	service := service.NewTaskService(pub)
	handler := handler.NewHandler(service)
	mux := http.NewServeMux()
	mux.Handle("GET /health", http.HandlerFunc(handler.Hello))
	mux.Handle("POST /task", http.HandlerFunc(handler.CreateTask))

	err := http.ListenAndServe(":3001", mux)
	if err != nil {
		fmt.Println("Error sta rti ng the server ", err)
		os.Exit(1)

	}
}
