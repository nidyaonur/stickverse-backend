package entities

import "gorm.io/gorm"

type LocationResource struct {
	gorm.Model
	ResourceID       uint    `gorm:"not null"`
	LocationID       uint    `gorm:"not null"`
	Quantity         float64 `gorm:"not null"`
	AllocatedWorkers uint64  `gorm:"not null;default:0"`
	Multiplier       float64 `gorm:"not null;default:0.1"`

	Resource Resource `gorm:"foreignkey:ResourceID"`
	Location Location `gorm:"foreignkey:LocationID"`
}
