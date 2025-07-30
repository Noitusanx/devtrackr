package repository

import (
	"context"

	"devtracker/internal/domain/model"

	"github.com/google/uuid"
)


type LogRepository interface {
	Create(ctx context.Context, log *model.Log) error
	FindByID(ctx context.Context, id uuid.UUID) (*model.Log, error)
	FindByUser(ctx context.Context, userID uuid.UUID) ([]model.Log, error)
	FindByProject(ctx context.Context, projectID uuid.UUID) ([]model.Log, error)
	Update(ctx context.Context, log *model.Log) error
	Delete(ctx context.Context, id uuid.UUID) error
}
