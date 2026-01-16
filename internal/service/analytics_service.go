package service

import (
	"context"
	"devtracker/internal/repository"
	"devtracker/pkg/util"
	"time"

	"github.com/google/uuid"
)

type AnalyticsService struct {
	logRepo repository.LogRepository
	projectRepo repository.ProjectRepository
	milestoneRepo repository.MilestoneRepository
}

type ProgressMetrics struct {
    TotalHours        float64
    WeeklyAverage     float64
    DaysActive        int
    DaysInactive      int
    MilestonesTotal   int
    MilestonesCompleted int
    ProgressPercent   float64
    DaysRemaining     int
    PredictedCompletion string  // "on-track", "at-risk", "delayed"
}

func (s *AnalyticsService) CalculateProgress(ctx context.Context, projectID uuid.UUID)(*ProgressMetrics, error){
	project, _ := s.projectRepo.FindByID(ctx, projectID);


	if project == nil{
		return nil, util.ErrNotFound("Project not found")
	}

	logs, _ := s.logRepo.FindByProject(ctx, projectID);


	milestones, _ := s.milestoneRepo.FindByProject(ctx, projectID);

	totalHours := 0.0

	for _, log := range logs {
		totalHours += float64(log.DurationMinutes) / 60.0
	}

	completedMilestones := 0
	for _, milestone := range milestones {
		if milestone.Status == "done" {
			completedMilestones++
		}
	}

	milestoneTotal := len(milestones)

	if milestoneTotal == 0 {
		milestoneTotal = 1
	}

	progressPercent := float64(completedMilestones)/float64(milestoneTotal) * 100;

	daysRemaining := 0
    if project.Deadline != nil {
        duration := time.Until(*project.Deadline) 
   
        if duration > 0 {
            daysRemaining = int(duration.Hours() / 24)
        }
    }


	prediction := "on-track"

	if progressPercent < 50 && daysRemaining < 30 {
		prediction = "at-risk"
	}

	return &ProgressMetrics{
		TotalHours: totalHours,
		MilestonesTotal: len(milestones),
		MilestonesCompleted: completedMilestones,
		ProgressPercent: progressPercent,
		DaysRemaining: daysRemaining,
		PredictedCompletion: prediction,
	}, nil
}

