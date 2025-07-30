package model

import (
	"time"

	"github.com/google/uuid"
)

type Report struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProjectID uuid.UUID `gorm:"type:uuid;index;not null"`
	URLPDF    string    `gorm:"size:255;not null"`
	GeneratedAt time.Time
}