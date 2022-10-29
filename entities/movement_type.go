package entities

import "gorm.io/gorm"

type MovementType struct {
	gorm.Model
	Name       string `gorm:"not null"`
	AllowsWait bool   `gorm:"not null"`
}
