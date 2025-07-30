package dto

import "devtracker/internal/domain/model"

type CreateMilestoneRequest struct {
	Name      string `json:"name" validate:"required"`
	OrderIdx int    `json:"order_idx" validate:"omitempty"`
	Status model.MilestoneStatus `json:"status" validate:"omitempty,oneof=pending in_progress done"`
	DueDate   string `json:"due_date" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`

}


type MilestoneResponse struct {
	ID         string `json:"id"`
	ProjectID  string `json:"project_id"`
	OrderIdx   int    `json:"order_idx"`
	Name       string `json:"name"`
	Status     model.MilestoneStatus `json:"status"`
	DueDate    string `json:"due_date,omitempty"`
	CompletedAt string `json:"completed_at,omitempty"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

type UpdateMilestoneRequest struct {
	ID         string `json:"id" validate:"required,uuid"`
	ProjectID  string `json:"project_id" validate:"required,uuid"`
	Name       string `json:"name" validate:"omitempty"`
	OrderIdx   *int    `json:"order_idx" validate:"omitempty"`
	Status     model.MilestoneStatus `json:"status" validate:"omitempty,oneof=pending in_progress done"`
	DueDate    *string `json:"due_date" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	CompletedAt string `json:"completed_at" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
}
