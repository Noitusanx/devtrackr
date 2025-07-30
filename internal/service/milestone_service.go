package service

import (
	"context"
	"time"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"
	"devtracker/pkg/util"

	"github.com/google/uuid"
)

type MilestoneService struct {
	repo repository.MilestoneRepository
}


func NewMilestoneService(repo repository.MilestoneRepository) *MilestoneService{
	return &MilestoneService{
		repo: repo,
	}
}


func (s *MilestoneService) CreateMilestone(ctx context.Context, projectID uuid.UUID, name string, orderIdx int, status model.MilestoneStatus, dueDate *time.Time) (*model.Milestone, error) {

	if name == "" {
		return nil, util.ErrBadRequest("name required")
	}

	if dueDate != nil && dueDate.Before(time.Now()) {
		return nil, util.ErrBadRequest("due date cannot be in the past")
	}

	

	var completedAt *time.Time
	if status == model.StatusDone {
		now := time.Now()
		completedAt = &now
	}

	milestone := &model.Milestone{
		ID:        uuid.New(),
		ProjectID: projectID,
		Name:      name,
		OrderIdx:  orderIdx,
		Status:    status,
		DueDate:   dueDate,
		CompletedAt: completedAt,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, milestone); err != nil {
		if util.IsUniqueViolation(err) {
			return nil, util.ErrConflict("milestone with this name already exists")
		}
		return nil, err
	}
	return milestone, nil
	
} 


func (s *MilestoneService) GetMilestonesByProject(ctx context.Context, userId, projectID uuid.UUID) ([]model.Milestone, error){
	milestones, err := s.repo.FindByProject(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return milestones, nil
}


func (s *MilestoneService) GetMilestoneByID(ctx context.Context, id uuid.UUID) (*model.Milestone, error) {
	milestone, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return milestone, nil
}

func (s *MilestoneService) UpdateMilestone(ctx context.Context, m *model.Milestone) error {
	if m.Name == "" {
		return util.ErrBadRequest("name required")
	}
	if m.DueDate != nil && m.DueDate.Before(time.Now()) {
		return util.ErrBadRequest("due date cannot be in the past")
	}
	if m.Status == model.StatusDone && m.CompletedAt == nil {
		now := time.Now()
		m.CompletedAt = &now
	}
	return s.repo.Update(ctx, m)
}


func (s *MilestoneService) DeleteMilestone(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return util.ErrNotFound("milestone not found")
	}
	return nil
}