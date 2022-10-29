package entities

import "gorm.io/gorm"

type MembershipAction struct {
	gorm.Model
	Name      string            `gorm:"unique;not null"`
	NameLocal map[string]string `gorm:"type:jsonb"`
}
