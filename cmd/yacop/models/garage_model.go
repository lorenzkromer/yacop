package models

import (
	"gorm.io/gorm"
)

// Garage Model
type Garage struct {
	gorm.Model
	ID       string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID   string    `gorm:"type:uuid;uniqueIndex" json:"user_id"`
	Vehicles []Vehicle `gorm:"foreignKey:GarageID" json:"vehicles"`
}
