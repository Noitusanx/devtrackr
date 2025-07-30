package model

import (
	"time"

	"github.com/google/uuid"
)

type Log struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProjectID       uuid.UUID `gorm:"type:uuid;index;not null"`
	UserID          uuid.UUID `gorm:"type:uuid;index;not null"`
	Text            string    `gorm:"type:text"`
	DurationMinutes int       `gorm:"not null"`         
	LoggedAt        time.Time `gorm:"index"`             
	CreatedAt       time.Time
}
