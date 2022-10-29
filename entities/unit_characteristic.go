package entities

import "gorm.io/gorm"

type UnitCharacteristic struct {
	gorm.Model
	UnitID           uint    `gorm:"not null"`
	CharacteristicID uint    `gorm:"not null"`
	Amount           float64 `gorm:"not null"`

	Unit           *Unit           `gorm:"foreignkey:UnitID"`
	Characteristic *Characteristic `gorm:"foreignkey:CharacteristicID"`
}
