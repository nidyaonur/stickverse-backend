package entities

import "gorm.io/gorm"

type MovementType struct {
	gorm.Model
	Name       string `gorm:"unique;not null"`
	AllowsWait bool   `gorm:"not null"`
}
