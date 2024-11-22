package repository

import (
	"TaskWeave/pkg/models"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Task) (int, error) // Возвращает ID
	GetTask(ctx context.Context, id int) (*models.Task, error)
	GetAllTasks(ctx context.Context) ([]*models.Task, error)
	DeleteTask(ctx context.Context, id int) error
	UpdateTask(ctx context.Context, task *models.Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(ctx context.Context, task *models.Task) (int, error) {
	result := r.db.Create(task)
	return int(task.ID), result.Error // Возвращаем ID
}

func (r *taskRepository) GetTask(ctx context.Context, id int) (*models.Task, error) {
	var task models.Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *taskRepository) GetAllTasks(ctx context.Context) ([]*models.Task, error) {
	var tasks []*models.Task
	err := r.db.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	return r.db.Delete(&models.Task{}, id).Error 
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *models.Task) error {
	return r.db.Save(task).Error
}