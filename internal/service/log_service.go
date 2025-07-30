package service

import (
	"context"
	"time"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"
	"devtracker/pkg/util"

	"github.com/google/uuid"
)


type LogService struct {
	repo repository.LogRepository
}

func NewLogService(repo repository.LogRepository) *LogService {
	return &LogService{
		repo: repo,
	}
}


func (s *LogService) CreateLog(ctx context.Context, projectId, userID uuid.UUID,  text string, durationMinutes int, loggedAt time.Time) error {
	if projectId == uuid.Nil {
		return util.ErrBadRequest("project ID is required")
	}

	if userID == uuid.Nil {
		return util.ErrBadRequest("user ID is required")
	}

	if text == "" {
		return util.ErrBadRequest("log text is required")
	}
	if durationMinutes <= 0 {
		return util.ErrBadRequest("duration must be greater than zero")
	}

	log := &model.Log{
		ID:              uuid.New(),
		ProjectID:       projectId,
		UserID:          userID,
		Text:            text,
		DurationMinutes: durationMinutes,
		LoggedAt:        loggedAt,
		CreatedAt:       time.Now(),
	}

	return s.repo.Create(ctx, log)
}


func (s *LogService) GetLogByID(ctx context.Context, id uuid.UUID) (*model.Log, error) {
	if id == uuid.Nil {
		return nil, util.ErrBadRequest("log ID is required")
	}

	log, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return log, nil
}

func (s *LogService) GetLogsByUser(ctx context.Context, userID uuid.UUID) ([]model.Log, error) {
	if userID == uuid.Nil {
		return nil, util.ErrBadRequest("user ID is required")
	}

	logs, err := s.repo.FindByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (s *LogService) GetLogsByProject(ctx context.Context, projectID uuid.UUID) ([]model.Log, error) {

	if projectID == uuid.Nil {
		return nil, util.ErrBadRequest("project ID is required")
	}

	logs, err := s.repo.FindByProject(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (s *LogService) UpdateLog(ctx context.Context, userId uuid.UUID, log *model.Log) error {
	orig, err := s.repo.FindByID(ctx, log.ID)
	if err != nil {
		return err
	}

	if orig.UserID != userId{
		return util.ErrUnauthorized("you do not have permission to update this log")
	}

	if log.ID == uuid.Nil {
		return util.ErrBadRequest("log ID is required")
	}

	if log.Text == "" {
		return util.ErrBadRequest("log text is required")
	}

	if log.DurationMinutes <= 0 {
		return util.ErrBadRequest("duration must be greater than zero")
	}

	return s.repo.Update(ctx, log)
}

func (s *LogService) DeleteLog(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return util.ErrBadRequest("log ID is required")
	}

	return s.repo.Delete(ctx, id)
}


