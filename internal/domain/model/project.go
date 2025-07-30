package model

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null"`
	Name      string    `gorm:"size:120;not null"`
	Deadline  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
