package model

import (
	"time"

	"github.com/google/uuid"
)

type InsightType string

const (
	InsightSummary  InsightType = "summary"
	InsightProgress InsightType = "progress"
)


type AIInsight struct {
	ID         uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProjectID  uuid.UUID   `gorm:"type:uuid;index;not null"`
	Type       InsightType `gorm:"type:text;not null"`
	Content    string      `gorm:"type:text"` 
	GeneratedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}