package entities

import "gorm.io/gorm"

type UnitCharacteristic struct {
	gorm.Model
	UnitID             uint           `gorm:"not null"`
	CharacteristicID   uint           `gorm:"not null"`
	CharacteristicType string         `gorm:"not null"`
	StringValue        string         `gorm:"null"`
	FloatValue         float64        `gorm:"null"`
	BooleanValue       bool           `gorm:"null"`
	Unit               Unit           `gorm:"foreignkey:UnitID"`
	Characteristic     Characteristic `gorm:"foreignkey:CharacteristicID"`
}
