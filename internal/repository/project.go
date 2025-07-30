package repository

import (
	"context"

	"devtracker/internal/domain/model"

	"github.com/google/uuid"
)


type ProjectRepository interface {
	Create(ctx context.Context, p *model.Project) error
	FindByUser(ctx context.Context, userID uuid.UUID) ([]model.Project, error)
	FindByID(ctx context.Context, id uuid.UUID) (*model.Project, error)
	Update(ctx context.Context, p *model.Project) error
	Delete(ctx context.Context, id uuid.UUID) error
}