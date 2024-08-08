package persistence

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) repository.ProjectRepository {
	return &ProjectRepositoryImpl{db: db}
}

func (r *ProjectRepositoryImpl) GetProjects(ctx context.Context) ([]models.Project, error) {
	var projects []models.Project
	result := r.db.WithContext(ctx).Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (r *ProjectRepositoryImpl) GetProject(ctx context.Context, id string) (*models.Project, error) {
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

func (r *ProjectRepositoryImpl) CreateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	result := r.db.WithContext(ctx).Create(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *ProjectRepositoryImpl) UpdateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	result := r.db.WithContext(ctx).Save(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *ProjectRepositoryImpl) DeleteProject(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.Project{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
