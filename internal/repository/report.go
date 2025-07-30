package repository

import (
	"context"

	"devtracker/internal/domain/model"

	"github.com/google/uuid"
)


type ReportRepository interface {
	Create(ctx context.Context, report *model.Report) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Report, error)
	FindByProject(ctx context.Context, projectID uuid.UUID) ([]model.Report, error)
	Update(ctx context.Context, report *model.Report) error
	Delete(ctx context.Context, id uuid.UUID) error
}