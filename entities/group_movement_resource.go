package entities

import "gorm.io/gorm"

type GroupMovementResource struct {
	gorm.Model
	GroupMovementID uint `gorm:"column:group_movement_id"`
	ResourceID      uint `gorm:"column:resource_id"`
	Amount          uint `gorm:"column:amount"`

	GroupMovement *GroupMovement `gorm:"foreignkey:GroupMovementID"`
	Resource      *Resource      `gorm:"foreignkey:ResourceID"`
}
