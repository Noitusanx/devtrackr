package postgres

import (
	"context"
	"time"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type AIInsightPG struct {
	db *gorm.DB
}

func NewAIInsightPG(db *gorm.DB) repository.AIInsightRepository {
	return &AIInsightPG{db}
}

func (r *AIInsightPG) Create(ctx context.Context, insight *model.AIInsight) error {
	return r.db.WithContext(ctx).Create(insight).Error
}

func (r *AIInsightPG) FindByID(ctx context.Context, id uuid.UUID) (*model.AIInsight, error) {
	var insight model.AIInsight
	err := r.db.WithContext(ctx).First(&insight, "id = ?", id).Error
	return &insight, err
}

func (r *AIInsightPG) FindByProject(ctx context.Context, projectID uuid.UUID) ([]model.AIInsight, error) {
	var insights []model.AIInsight
	err := r.db.WithContext(ctx).
		Where("project_id = ?", projectID).Order("created_at desc").Find(&insights).Error
	return insights, err
}


func (r *AIInsightPG) FindByProjectAndTypeAndDate(ctx context.Context, projectID uuid.UUID, insightType model.InsightType, date time.Time) (*model.AIInsight, error) {
	var insight model.AIInsight
	err := r.db.WithContext(ctx).
		Where("project_id = ? AND type = ? AND DATE(generated_at) = ?", projectID, insightType, date).
		First(&insight).Error

		if err != nil {
			return nil, err
		}
	return &insight, err
}



func (r *AIInsightPG) Update(ctx context.Context, insight *model.AIInsight) error {
	return r.db.WithContext(ctx).Save(insight).Error
}

func (r *AIInsightPG) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&model.AIInsight{}, "id = ?", id).Error
}