package service

import (
	"context"
	"time"

	"devtracker/internal/domain/model"
	"devtracker/internal/repository"
	"devtracker/pkg/util"

	"github.com/google/uuid"
)

type ProjectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) CreateProject(ctx context.Context, userID uuid.UUID, name string, deadline *time.Time) (*model.Project, error) {


	if name == "" {
		return nil, util.ErrBadRequest("name required")
	}

	if deadline != nil{
		if deadline.Before(time.Now()){
			return nil, util.ErrBadRequest("deadline cannot be in the past")
		}
	}

	project := &model.Project{
		ID:       uuid.New(),
		UserID:    userID,
		Name:     name,
		Deadline: deadline,
	}

	if err := s.repo.Create(ctx, project); err != nil {
	if util.IsUniqueViolation(err) {
		return nil, util.ErrConflict("project with this name already exists")
	}
	return nil, err
}
	return project, nil
}

func (s *ProjectService) GetProjectByID(ctx context.Context, id uuid.UUID) (*model.Project, error) {
	proj, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return proj, nil

}

func (s *ProjectService) ListProjectByUser(ctx context.Context, userID uuid.UUID) ([]model.Project, error){
	projects, err := s.repo.FindByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return projects, nil
}



func (s *ProjectService) UpdateProject(ctx context.Context, userID uuid.UUID, p *model.Project) error {
	if p.ID == uuid.Nil {
		return util.ErrBadRequest("project ID is required")
	}

	if p.Name == "" {
		return util.ErrBadRequest("name is required")
	}

	if p.Deadline != nil && p.Deadline.Before(time.Now()) {
		return util.ErrBadRequest("deadline cannot be in the past")
	}

	existingProject, err := s.repo.FindByID(ctx, p.ID)
	if err != nil {
		return util.ErrNotFound("project not found")
	}

	if existingProject.UserID != userID {
		return util.ErrUnauthorized("you do not have permission to update this project")
	}

	p.UserID = userID
	p.UpdatedAt = time.Now()

	return s.repo.Update(ctx, p)
}

func (s *ProjectService) DeleteProject(ctx context.Context, userID, id uuid.UUID) error {
	proj, err := s.repo.FindByID(ctx, id); if err != nil {
		return util.ErrNotFound("project not found")
	}

	if proj.UserID != userID {
		return util.ErrUnauthorized("you do not have permission to delete this project")
	}


	return s.repo.Delete(ctx, id)
}