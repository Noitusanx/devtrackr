package dto

type CreateProjectRequest struct {
	Name     string `json:"name" validate:"required"`
	Deadline string `json:"deadline" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}

type UpdateProjectRequest struct {
	ID	   string `json:"id" validate:"required,uuid"`
	Name     string `json:"name" validate:"omitempty"`
	Deadline string `json:"deadline" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}


type ProjectResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Deadline string `json:"deadline" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

