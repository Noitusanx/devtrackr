package service

import (
	"context"
	"errors"
	"time"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"
	"devtracker/pkg/util"

	"github.com/google/uuid"
)

type AIInsightService struct {
	repo repository.AIInsightRepository
}


func NewAIInsightService(repo repository.AIInsightRepository) *AIInsightService{
	return &AIInsightService{
		repo: repo,
	}
}

func (s *AIInsightService) CreateInsight(ctx context.Context, projectID uuid.UUID, insightType model.InsightType, content string) (*model.AIInsight, error) {
	if projectID == uuid.Nil{
		return nil, util.ErrBadRequest("project ID cannot be empty")
	}

	if insightType == ""{
		return nil, util.ErrBadRequest("insight type cannot be empty")
}

insights:= &model.AIInsight{
	ProjectID: projectID,
	Type: insightType,
	Content: content,
	GeneratedAt: time.Now(),
}

	if err := s.repo.Create(ctx, insights); err != nil {
		return nil, err
	}
	return insights, nil
}

func (s *AIInsightService) GetInsightByID(ctx context.Context, id uuid.UUID) (*model.AIInsight, error) {
	if id == uuid.Nil {
		return nil, util.ErrBadRequest("insight ID cannot be empty")
	}

	insight, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return insight, nil
}

func (s *AIInsightService) GetInsightsByProject(ctx context.Context, projectID uuid.UUID) ([]model.AIInsight, error) {
	if projectID == uuid.Nil {
		return nil, errors.New("project ID cannot be empty")
	}

	insights, err := s.repo.FindByProject(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return insights, nil
}

func (s *AIInsightService) GetInsightByProjectAndTypeAndDate(ctx context.Context, projectID uuid.UUID, insightType model.InsightType, date time.Time) (*model.AIInsight, error) {
	if projectID == uuid.Nil {
		return nil, util.ErrBadRequest("project ID cannot be empty")
	}

	if insightType == "" {
		return nil, util.ErrBadRequest("insight type cannot be empty")
	}

	if date.IsZero() {
		return nil, util.ErrBadRequest("date cannot be zero value")
	}

	insight, err := s.repo.FindByProjectAndTypeAndDate(ctx, projectID, insightType, date)
	if err != nil {
		return nil, err
	}
	return insight, nil
}


func (s *AIInsightService) UpdateInsight(ctx context.Context, insight *model.AIInsight) (*model.AIInsight, error) {
	if insight.ID == uuid.Nil {
		return nil, util.ErrBadRequest("insight ID cannot be empty")
	}

	if insight.ProjectID == uuid.Nil {
		return nil, util.ErrBadRequest("project ID cannot be empty")
	}

	if insight.Type == "" {
		return nil, util.ErrBadRequest("insight type cannot be empty")
	}

	if insight.Content == "" {
		return nil, util.ErrBadRequest("content cannot be empty")
	}


	if err := s.repo.Update(ctx, insight); err != nil {
		return nil, err
	}

	return insight, nil
}


func (s *AIInsightService) DeleteInsight(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return util.ErrBadRequest("insight ID cannot be empty")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}



