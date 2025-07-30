package service

import (
	"context"
	"time"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"
	"devtracker/pkg/util"

	"github.com/google/uuid"
)

type ReportService struct {
	repo repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) *ReportService{
	return &ReportService{
		repo: repo,
	}
}


func (s *ReportService) CreateReport(ctx context.Context, projectID, userID uuid.UUID, urlPDF string) (*model.Report, error) {
	
	
	if urlPDF == "" {
		return nil, util.ErrBadRequest("urlPDF is required")
	}

	report := &model.Report{
		ID:        uuid.New(),
		ProjectID: projectID,
		URLPDF:    urlPDF,
		GeneratedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, report); err != nil {
		return nil, err
	}
	return report, nil
}

func (s *ReportService) GetReportByID(ctx context.Context, id uuid.UUID) (*model.Report, error) {
	if id == uuid.Nil {
		return nil, util.ErrBadRequest("report ID is required")
	}

	report, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (s *ReportService) GetReportsByProject(ctx context.Context, projectID uuid.UUID) ([]model.Report, error) {
	if projectID == uuid.Nil {
		return nil, util.ErrBadRequest("project ID is required")
	}

	reports, err := s.repo.FindByProject(ctx, projectID)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (s *ReportService) UpdateReport(ctx context.Context, report *model.Report) (*model.Report, error) {
	if report.ID == uuid.Nil {
		return nil, util.ErrBadRequest("report ID is required")
	}
	if report.URLPDF == "" {
		return nil, util.ErrBadRequest("urlPDF is required")
	}
	if err := s.repo.Update(ctx, report); err != nil {
		return nil, err
	}
	return report, nil
}


func (s *ReportService) DeleteReport(ctx context.Context, id uuid.UUID) error {
	if id == uuid.Nil {
		return util.ErrBadRequest("report ID is required")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

