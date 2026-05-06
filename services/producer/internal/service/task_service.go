package service

import (
	"github.com/Sahadat20/distributed-video-processing-pipeline.git/producer/internal/domain"
	"github.com/Sahadat20/distributed-video-processing-pipeline.git/producer/internal/publisher"
	"github.com/google/uuid"
)

type TaskService struct {
	publisher *publisher.RabbitMQPublisher
}

func NewTaskService(pub *publisher.RabbitMQPublisher) *TaskService {
	return &TaskService{
		publisher: pub,
	}
}

func (s *TaskService) CreateTask(name string) error {
	task := domain.Task{
		ID:   uuid.New().String(),
		Name: name,
	}
	return s.publisher.Publish(task)

}
