package service

import (
	"context"
	"errors"
	"time"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"
	"devtracker/pkg/util"

	"fmt"

	"github.com/google/uuid"
)

type AIInsightService struct {
	repo repository.AIInsightRepository

	analyticsSvc *AnalyticsService
}


func NewAIInsightService(repo repository.AIInsightRepository, analyticsSvc *AnalyticsService) *AIInsightService{
	return &AIInsightService{
		repo: repo,
		analyticsSvc: analyticsSvc,
	}
}

func (s *AIInsightService) GenerateProgressInsight(ctx context.Context, projectID uuid.UUID) (*model.AIInsight, error) {
    // 1. Dapatkan metrik dari AnalyticsService
    metrics, err := s.analyticsSvc.CalculateProgress(ctx, projectID)
    if err != nil {
        return nil, err
    }

    // 2. Buat Teks Insight berdasarkan Rule (Mocking AI)
    var content string
    var status model.InsightStatus 

    switch metrics.PredictedCompletion {
    case "on-track":
        content = fmt.Sprintf(
            "✅ Kamu on-track! Progress: %.1f%% (%d/%d Milestones). Total jam kerja: %.1f jam. Pertahankan ritme ini untuk selesai tepat waktu.",
            metrics.ProgressPercent,
            metrics.MilestonesCompleted,
            metrics.MilestonesTotal,
            metrics.TotalHours,
        )
        status = model.StatusOnTrack
    case "at-risk":
        content = fmt.Sprintf(
            "⚠️ Perhatian! Progress: %.1f%%, tapi hanya %d hari tersisa. Project berisiko tertunda. Fokus pada milestone paling kritikal dan tingkatkan jam kerja.",
            metrics.ProgressPercent,
            metrics.DaysRemaining,
        )
        status = model.StatusAtRisk
    case "delayed":
        content = fmt.Sprintf(
            "🔴 Terlambat! Progress baru %.1f%%, padahal deadline tinggal %d hari. Segera konsultasi dengan mentor/dosen dan buat rencana pemulihan proyek.",
            metrics.ProgressPercent,
            metrics.DaysRemaining,
        )
        status = model.StatusDelayed
    default:
        content = "Analisis tidak dapat disimpulkan. Perlu lebih banyak data log."
        status = model.InsightStatus("UNKNOWN")
    }

    // 3. Simpan ke database
    // Perhatikan tipe kembalian di NewAIInsightService
    insight := &model.AIInsight{
        ProjectID:   projectID,
        Type:        model.InsightProgress,
        Content:     content,
        Status:      status, // ✅ Simpan status ke model
    }

    if err := s.repo.Create(ctx, insight); err != nil {
        return nil, err
    }

    return insight, nil
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
	orig, err := s.repo.FindByID(ctx, insight.ID)

	if err != nil {
		return nil, err
	}


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

	insight.GeneratedAt = orig.GeneratedAt


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



