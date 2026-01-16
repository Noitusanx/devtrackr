package postgres

import (
	"context"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogPG struct {
	db *gorm.DB

}


func NewLogPG(db *gorm.DB) repository.LogRepository {
	return &LogPG{db: db}
}

func (r *LogPG) Create(ctx context.Context, log *model.Log) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *LogPG) FindByID(ctx context.Context, id uuid.UUID) (*model.Log, error) {
	var log model.Log
	err := r.db.WithContext(ctx).First(&log, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil;
		}
		return nil, err;
	}

	return &log, err
}

func (r *LogPG) FindByUser(ctx context.Context, userID uuid.UUID) ([]model.Log, error) {
	var logs []model.Log
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).Order("created_at desc").Find(&logs).Error

	if err != nil {
		return nil, err
	}
	
	return logs, err
}	

func (r *LogPG) FindByProject(ctx context.Context, projectID uuid.UUID) ([]model.Log, error) {
	var logs []model.Log
	err := r.db.WithContext(ctx).
		Where("project_id = ?", projectID).Order("created_at desc").Find(&logs).Error

	if err != nil {
		return nil, err
	}

	return logs, err
}

func (r *LogPG) Update(ctx context.Context, log *model.Log) error {
	return r.db.WithContext(ctx).Save(log).Error
}

func (r *LogPG) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Log{}, "id = ?", id).Error
}

