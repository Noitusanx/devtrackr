package dto

type CreateLogRequest struct {
	ProjectID       string `json:"project_id" validate:"required,uuid"`
	Text            string `json:"text" validate:"required"`
	DurationMinutes int    `json:"duration_minutes" validate:"required,min=1"`
	LoggedAt        string `json:"logged_at" validate:"required,datetime=2006-01-02T15:04:05Z07:00"` // RFC3339 format
}

type UpdateLogRequest struct {
	ProjectID       *string `json:"project_id" validate:"omitempty,uuid"`
	Text            *string `json:"text" validate:"omitempty"`
	DurationMinutes *int   `json:"duration_minutes" validate:"omitempty,min=1"`
	LoggedAt        *string `json:"logged_at" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"` // RFC3339 format
}


