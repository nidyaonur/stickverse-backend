package entities

import (
	"time"

	"gorm.io/gorm"
)

type MembershipHistory struct {
	gorm.Model
	AllianceMemberID uint       `gorm:"not null"`
	MembershipTypeID uint       `gorm:"not null"`
	DateFrom         time.Time  `gorm:"not null"`
	EndedAt          *time.Time `gorm:"null"`

	AllianceMember *AllianceMember `gorm:"foreignkey:AllianceMemberID"`
	MembershipType *MembershipType `gorm:"foreignkey:MembershipTypeID"`
}
