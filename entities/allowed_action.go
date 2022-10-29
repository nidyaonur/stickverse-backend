package entities

import "gorm.io/gorm"

type AllowedAction struct {
	gorm.Model
	MembershipActionID uint `gorm:"index"`
	MembershipTypeID   uint `gorm:"index"`

	MembershipAction MembershipAction `gorm:"foreignkey:MembershipActionID"`
	MembershipType   MembershipType   `gorm:"foreignkey:MembershipTypeID"`
}
