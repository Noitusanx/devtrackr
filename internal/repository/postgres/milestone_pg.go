package postgres

import (
	"context"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type MilestonePG struct {
	db *gorm.DB
}

// constructor
func NewMilestonePG(db *gorm.DB) repository.MilestoneRepository {
	return &MilestonePG{db}
}

func (r *MilestonePG) Create(ctx context.Context, m *model.Milestone) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *MilestonePG) FindByProject(ctx context.Context, projectID uuid.UUID) ([]model.Milestone, error) {
	var res []model.Milestone
	err := r.db.WithContext(ctx).
		Where("project_id = ?", projectID).Order("created_at desc").Order("status != 'done', order_idx asc").Find(&res).Error
	return res, err
}

func (r *MilestonePG) FindByID(ctx context.Context, id uuid.UUID) (*model.Milestone, error) {
	var ms model.Milestone
	err := r.db.WithContext(ctx).First(&ms, "id = ?", id).Error
	return &ms, err
}

func (r *MilestonePG) Update(ctx context.Context, m *model.Milestone) error {
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *MilestonePG) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Milestone{}, "id = ?", id).Error
}
