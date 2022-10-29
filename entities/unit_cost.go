package entities

import "gorm.io/gorm"

type UnitCost struct {
	gorm.Model
	UnitID uint    `gorm:"not null"`
	Cost   float64 `gorm:"not null"`
	Unit   Unit    `gorm:"foreignkey:UnitID"`
}
