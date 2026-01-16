package dto

type CreateLogRequest struct {
	ProjectID       string `json:"project_id" validate:"required,uuid"`
	MilestoneID 	*string `json:"milestone_id" validate:"omitempty,uuid"`
	Description            string `json:"description" validate:"required"`
	DurationMinutes int    `json:"duration_minutes" validate:"required,min=1"`
	LoggedAt        string `json:"logged_at" validate:"required,datetime=2006-01-02T15:04:05Z07:00"` // RFC3339 format
}

type UpdateLogRequest struct {
	ProjectID       *string `json:"project_id" validate:"omitempty,uuid"`
	MilestoneID *string `json:"milestone_id" validate:"omitempty,uuid"`
	Description            *string `json:"description" validate:"omitempty"`
	DurationMinutes *int   `json:"duration_minutes" validate:"omitempty,min=1"`
	LoggedAt        *string `json:"logged_at" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"` // RFC3339 format
	
}


type LogResponse struct {
	ID              string  `json:"id"`
	ProjectID       string  `json:"project_id"`
	MilestoneID     *string `json:"milestone_id,omitempty"`
	UserID          string  `json:"user_id"`
	Description     string  `json:"description"`
	DurationMinutes int     `json:"duration_minutes"`
	LoggedAt        string  `json:"logged_at"`          // RFC3339 string
	CreatedAt       string  `json:"created_at"`         // RFC3339 string
	UpdatedAt       string  `json:"updated_at,omitempty"` // RFC3339 string, boleh kosong
}