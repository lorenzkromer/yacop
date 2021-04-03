package models

import (
	"gorm.io/gorm"
	"time"
)

// Vehicle Model
type Vehicle struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FullName  string         `gorm:"column:full_name;uniqueIndex;type:string;size:255" json:"full_name"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}
