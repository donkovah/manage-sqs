package persistence

import (
	"be/src/domain/models"
	"be/src/domain/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type TimelineRepositoryImpl struct {
	db *gorm.DB
}

func NewTimelineRepository(db *gorm.DB) repository.TimelineRepository {
	return &TimelineRepositoryImpl{db: db}
}

func (r *TimelineRepositoryImpl) GetTimelines(ctx context.Context) ([]models.Timeline, error) {
	var projects []models.Timeline
	result := r.db.WithContext(ctx).Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (r *TimelineRepositoryImpl) GetTimeline(ctx context.Context, id string) (*models.Timeline, error) {
	var project models.Timeline
	result := r.db.WithContext(ctx).First(&project, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func (r *TimelineRepositoryImpl) CreateTimeline(ctx context.Context, project *models.Timeline) (*models.Timeline, error) {
	result := r.db.WithContext(ctx).Create(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *TimelineRepositoryImpl) UpdateTimeline(ctx context.Context, project *models.Timeline) (*models.Timeline, error) {
	result := r.db.WithContext(ctx).Save(project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (r *TimelineRepositoryImpl) DeleteTimeline(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Delete(&models.Timeline{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
