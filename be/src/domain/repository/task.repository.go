package repository

import (
	"be/src/domain/models"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context) ([]models.Task, error)
	GetTask(ctx context.Context, id string) (*models.Task, error)
	CreateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

// GetTask implements TaskRepository.
func (r *taskRepository) GetTask(ctx context.Context, id string) (*models.Task, error) {
	var task models.Task
	result := r.db.WithContext(ctx).Find(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

// CreateTask implements TaskRepository.
func (r *taskRepository) GetTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	result := r.db.WithContext(ctx).Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *taskRepository) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	result := r.db.WithContext(ctx).Create((&task))
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil

}

// DeleteTask implements TaskRepository.
func (r *taskRepository) DeleteTask(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.Task{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateTask implements TaskRepository.
func (r *taskRepository) UpdateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	result := r.db.WithContext(ctx).Save(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}
