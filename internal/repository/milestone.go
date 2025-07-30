package repository

import (
	"context"

	"devtracker/internal/domain/model"

	"github.com/google/uuid"
)

type MilestoneRepository interface {
	Create(ctx context.Context, m *model.Milestone) error
	FindByProject(ctx context.Context, projectID uuid.UUID) ([]model.Milestone, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Milestone, error)
	Update(ctx context.Context, m *model.Milestone) error
	Delete(ctx context.Context, id uuid.UUID) error
}