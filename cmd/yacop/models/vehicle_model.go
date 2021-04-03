package models

import (
	"gorm.io/gorm"
)

// Vehicle Model
type Vehicle struct {
	gorm.Model
	ID             string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FullName       string         `gorm:"column:full_name;uniqueIndex;type:string;size:255" json:"full_name"`
	ManufacturerID string         `gorm:"column:manufacturer_id" json:"-"`
	Manufacturer   Manufacturer
}
