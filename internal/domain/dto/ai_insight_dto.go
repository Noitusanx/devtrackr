package dto

import "time"

type AIInsightRequest struct {
	ProjectID string `json:"project_id" validate:"required,uuid"`
	InsightType string `json:"insight_type" validate:"required,oneof=summary progress"`
	Content string `json:"content" validate:"required"`
}

type AIInsightResponse struct {
	ID string `json:"id"`
	ProjectID string `json:"project_id"`
	InsightType string `json:"insight_type"`
	Content string `json:"content"`
	GeneratedAt time.Time `json:"generated_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AIInsightUpdateRequest struct {
	ID *string `json:"id" validate:"omitempty,uuid"`
	ProjectID *string `json:"project_id" validate:"omitempty,uuid"`
	InsightType *string `json:"insight_type" validate:"omitempty,oneof=summary progress"`
	Content *string `json:"content" validate:"omitempty"`
	GeneratedAt *time.Time `json:"generated_at" validate:"omitempty"`
}
