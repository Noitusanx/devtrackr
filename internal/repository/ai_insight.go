package repository

import (
	"context"
	"time"

	"devtracker/internal/domain/model"

	"github.com/google/uuid"
)

type AIInsightRepository interface {
	Create(ctx context.Context, insight *model.AIInsight) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.AIInsight, error)
	FindByProject(ctx context.Context, projectID uuid.UUID) ([]model.AIInsight, error)
	FindByProjectAndTypeAndDate(ctx context.Context, projectID uuid.UUID, insightType model.InsightType, date time.Time) (*model.AIInsight, error)
	Update(ctx context.Context, insight *model.AIInsight) error
	Delete(ctx context.Context, id uuid.UUID) error
}