package entities

import "gorm.io/gorm"

type StructurePrerequisite struct {
	gorm.Model
	Type        string    `gorm:"type:varchar(255);not null"`
	SubType     string    `gorm:"type:varchar(255);not null"`
	UniqueID    string    `gorm:"unique;not null"`
	StructureID uint      `gorm:"not null"`
	Structure   Structure `gorm:"foreignkey:StructureID"`
	Formula     string    `gorm:"type:varchar(255);null"`
	Amount      int       `gorm:"type:int;default:0"`
}
