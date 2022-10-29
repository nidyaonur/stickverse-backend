package entities

import "gorm.io/gorm"

type ResearchResource struct {
	gorm.Model

	UpgradeFormula string `gorm:"type:text"`
	ResearchID     uint   `gorm:"not null"`
	ResourceID     uint   `gorm:"not null"`

	Research Research `gorm:"foreignkey:ResearchID"`
	Resource Resource `gorm:"foreignkey:ResourceID"`
}
