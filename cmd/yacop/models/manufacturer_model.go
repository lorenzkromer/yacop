package models

import (
	"gorm.io/gorm"
)

// Manufacturer Model
type Manufacturer struct {
	gorm.Model
	ID        string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name      string         `gorm:"column:name;uniqueIndex;type:string;size:255" json:"name"`
	Vehicles  []Vehicle      `gorm:"foreignKey:ID" json:"vehicles"`
}
