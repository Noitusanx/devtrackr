package model

import (
	"time"

	"github.com/google/uuid"
)

type MilestoneStatus string

const (
	StatusPending     MilestoneStatus = "pending"
	StatusInProgress MilestoneStatus = "in_progress"
	StatusDone       MilestoneStatus = "done"
)


type Milestone struct {
	ID         uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProjectID  uuid.UUID       `gorm:"type:uuid;index;not null"`
	Name       string          `gorm:"size:150;not null"`
	OrderIdx   int             `gorm:"default:0; "` 
	Status     MilestoneStatus `gorm:"type:text;default:'pending'"`
	DueDate    *time.Time
	CompletedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}