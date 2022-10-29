package entities

import (
	"time"

	"gorm.io/gorm"
)

type ResearchLevel struct {
	gorm.Model
	UserID           uint       `gorm:"not null"`
	ResearchID       uint       `gorm:"not null"`
	Level            uint       `gorm:"not null"`
	Research         Research   `gorm:"foreignkey:ResearchID"`
	User             User       `gorm:"foreignkey:UserID"`
	UpgradeOngoing   bool       `gorm:"not null"`
	UpgradeStartedAt *time.Time `gorm:"not null"`
	UpgradeEndedAt   *time.Time `gorm:"not null"`
}
