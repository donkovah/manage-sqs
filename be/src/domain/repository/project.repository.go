package repository

import (
	"be/src/domain/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetProjects(ctx context.Context) ([]models.Project, error)
	GetProject(ctx context.Context, id string) (*models.Project, error)
	CreateProject(ctx context.Context, project *models.Project) (*models.Project, error)
	UpdateProject(ctx context.Context, project *models.Project) (*models.Project, error)
	DeleteProject(ctx context.Context, id string) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) GetProjects(ctx context.Context) ([]models.Project, error) {
	var projects []models.Project
	result := r.db.WithContext(ctx).Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (r *projectRepository) GetProject(ctx context.Context, id string) (*models.Project, error) {
	var project models.Project
	result := r.db.WithContext(ctx).First(&project, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func (r *projectRepository) CreateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	result := r.db.WithContext(ctx).Create(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *projectRepository) UpdateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	result := r.db.WithContext(ctx).Save(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *projectRepository) DeleteProject(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.Project{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
