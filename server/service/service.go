package service

import (
	"TaskWeave/pkg/models"
	"TaskWeave/server/repository"
	"context"
)

type TaskService interface {
	CreateTask(ctx context.Context, task *models.Task) (int, error) // Возвращает ID
	GetTask(ctx context.Context, id int) (*models.Task, error)
	GetAllTasks(ctx context.Context) ([]*models.Task, error)
	DeleteTask(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, task *models.Task) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, task *models.Task) (int, error) {
	return s.repo.CreateTask(ctx, task)
}

func (s *taskService) GetTask(ctx context.Context, id int) (*models.Task, error) {
	return s.repo.GetTask(ctx, id)
}

func (s *taskService) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	return s.repo.GetAllTasks(ctx)
}

func (s *taskService) DeleteTask(ctx context.Context, id int) error {
	return s.repo.DeleteTask(ctx, id)
}

func (s *taskService) UpdateTask(ctx context.Context, task *models.Task) error {
	return s.repo.UpdateTask(ctx, task)
}