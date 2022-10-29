package entities

import "gorm.io/gorm"

type GroupMovementUnit struct {
	gorm.Model
	GroupMovementID uint `gorm:"not_null"`
	UnitID          uint `gorm:"not_null"`
	Amount          uint `gorm:"not_null"`

	GroupMovement *GroupMovement `gorm:"foreignkey:GroupMovementID"`
	Unit          *Unit          `gorm:"foreignkey:UnitID"`
}
