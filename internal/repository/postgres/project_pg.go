package postgres

import (
	"context"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectPG struct {
	db *gorm.DB
}

// Constructor
func NewProjectPG(db *gorm.DB) repository.ProjectRepository {
	return &ProjectPG{db}
}

func (r *ProjectPG) Create(ctx context.Context, p *model.Project) error {
	return r.db.WithContext(ctx).Create(p).Error
}

func (r *ProjectPG) FindByUser(ctx context.Context, userID uuid.UUID) ([]model.Project, error) {
	var res []model.Project
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).Order("created_at desc").Find(&res).Error
	return res, err
}

func (r *ProjectPG) FindByID(ctx context.Context, id uuid.UUID) (*model.Project, error) {
	var proj model.Project
	err := r.db.WithContext(ctx).First(&proj, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil
		}
		return nil, err
	}

	return &proj, err
}

func (r *ProjectPG) Update(ctx context.Context, p *model.Project) error {
	return r.db.WithContext(ctx).Save(p).Error
}

func (r *ProjectPG) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.Project{}, "id = ?", id).Error
}
