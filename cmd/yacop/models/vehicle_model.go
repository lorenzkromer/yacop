package models

import (
	"github.com/fitchlol/yacop/cmd/yacop/enums"
	"gorm.io/gorm"
)

// Vehicle Model
type Vehicle struct {
	gorm.Model
	ID                       string         `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FullName                 string         `gorm:"column:full_name;uniqueIndex;type:string;size:255" json:"full_name"`
	FuelType                 enums.FuelType `gorm:"column:fuel_type;type:fuel_type_enum;size:255" json:"fuel_type"`
	MaximumKilometersPerHour int            `gorm:"column:maximum_kilometers_per_hour;type:int" json:"maximum_fuel_type"`
	MaximumKilowatts         int            `gorm:"column:maximum_kilowatts;type:int" json:"maximum_kilowatts"`
	WeightInKilograms        int            `gorm:"column:weight_in_kilograms;type:int" json:"weight_in_kilograms"`
	ManufacturerID           string         `gorm:"column:manufacturer_id" json:"-"`
	GarageID                 string         `gorm:"column:garage_id" json:"-"`
	Manufacturer             Manufacturer
	Garage                   Garage
}
