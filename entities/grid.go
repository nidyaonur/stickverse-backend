package entities

import "gorm.io/gorm"

type Grid struct {
	gorm.Model
	X          int      `gorm:"index"`
	Y          int      `gorm:"index"`
	GridTypeID uint     `gorm:"index"`
	GridType   GridType `gorm:"foreignkey:GridTypeID"`
	Capacity   uint     `gorm:"default:0"`
	Occupied   uint     `gorm:"default:0"`
}
