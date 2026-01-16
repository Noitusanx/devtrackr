package postgres

import (
	"context"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReportPG struct {
	db *gorm.DB
}

func NewReportPG(db *gorm.DB) repository.ReportRepository {
	return &ReportPG{db}
}

func (r *ReportPG) Create(ctx context.Context, report *model.Report) error {
	return r.db.WithContext(ctx).Create(report).Error
}

func (r *ReportPG) FindByID(ctx context.Context, id uuid.UUID) (*model.Report, error) {
	var report model.Report
	err := r.db.WithContext(ctx).First(&report, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil
		}

		return nil, err
	}

	return &report, err
}

func (r *ReportPG) FindByProject(ctx context.Context, projectID uuid.UUID) ([]model.Report, error) {
	var reports []model.Report
	err := r.db.WithContext(ctx).
		Where("project_id = ?", projectID).Order("generated_at desc").Find(&reports).Error

	if err != nil {
		return nil, err
	}
	
	return reports, err
}

func (r *ReportPG) Update(ctx context.Context, report *model.Report) error {
	return r.db.WithContext(ctx).Save(report).Error
}

func (r *ReportPG) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Report{}, "id = ?", id).Error
}

