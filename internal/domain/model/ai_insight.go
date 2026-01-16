package model

import (
	"time"

	"github.com/google/uuid"
)

type InsightType string
type InsightStatus string // NEW: Untuk status prediksi

const (
    InsightSummary  InsightType = "summary"
    InsightProgress InsightType = "progress"
    
    // Status Taktis (Untuk UI)
    StatusOnTrack   InsightStatus = "ON_TRACK"
    StatusAtRisk    InsightStatus = "AT_RISK"
    StatusDelayed   InsightStatus = "DELAYED"
)

type AIInsight struct {
    ID          uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    ProjectID   uuid.UUID     `gorm:"type:uuid;index;not null"`
    Type        InsightType   `gorm:"type:text;not null"`
    Content     string        `gorm:"type:text"` 
    // ✅ NEW FIELD: Menyimpan Status Prediksi untuk UI
    Status      InsightStatus `gorm:"type:varchar(20);default:'UNKNOWN';not null"` 
    GeneratedAt time.Time     `gorm:"autoCreateTime"` // Waktu AI membuat analisis
}