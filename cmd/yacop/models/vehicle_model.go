package models

import (
	"github.com/fitchlol/yacop/cmd/yacop/enums"
	"gorm.io/gorm"
	"time"
)

// Vehicle Model
type Vehicle struct {
	gorm.Model
	ID                string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FullName          string         `gorm:"column:full_name;uniqueIndex;type:string;size:255" json:"full_name"`
	FuelType          enums.FuelType `gorm:"column:fuel_type;type:fuel_type_enum;size:255" json:"fuel_type"`
	ConstructionStart time.Time
	ConstructionEnd   time.Time
	ManufacturerID    string `gorm:"column:manufacturer_id" json:"-"`
	Manufacturer      Manufacturer
}
