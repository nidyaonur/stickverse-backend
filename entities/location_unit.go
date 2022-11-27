package entities

import "gorm.io/gorm"

type LocationUnit struct {
	gorm.Model
	LocationID *uint `gorm:"null"`
	UnitID     uint  `gorm:"not null"`
	Amount     uint  `gorm:"not null"`

	Location *Location `gorm:"foreignkey:LocationID"`
	Unit     *Unit     `gorm:"foreignkey:UnitID"`
}
