package entities

import "gorm.io/gorm"

type Grid struct {
	gorm.Model
	X          int       `gorm:"uniqueIndex:idx_x_y"`
	Y          int       `gorm:"uniqueIndex:idx_x_y"`
	GridTypeID uint      `gorm:"index"`
	GridType   GridType  `gorm:"foreignkey:GridTypeID"`
	ResourceID *uint     `gorm:"index"`
	Resource   *Resource `gorm:"foreignkey:ResourceID"`
	Capacity   uint      `gorm:"default:0"`
	Occupied   uint      `gorm:"default:0"`
}
