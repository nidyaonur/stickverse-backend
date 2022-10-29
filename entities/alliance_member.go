package entities

import (
	"time"

	"gorm.io/gorm"
)

type AllianceMember struct {
	gorm.Model
	AllianceID       uint           `gorm:"not null"`
	UserID           uint           `gorm:"not null"`
	MembershipTypeID uint           `gorm:"not null"`
	Alliance         Alliance       `gorm:"foreignkey:AllianceID"`
	User             User           `gorm:"foreignkey:UserID"`
	MembershipType   MembershipType `gorm:"foreignkey:MembershipTypeID"`
	EndedAt          *time.Time
}
