package entities

import "gorm.io/gorm"

type ResearchPrerequisite struct {
	gorm.Model
	Type    string `gorm:"uniqueIndex:idx_unique_preq; type:varchar(255);not null"`
	SubType string `gorm:"uniqueIndex:idx_unique_preq; type:varchar(255);not null"`
	Formula string `gorm:"type:varchar(255);null"`

	ResearchID uint     `gorm:"uniqueIndex:idx_unique_preq;not null"`
	Research   Research `gorm:"foreignKey:ResearchID"`

	RequiredStructureID *uint `gorm:"null"`
	RequiredResearchID  *uint `gorm:"null"`
	RequiredResourceID  *uint `gorm:"null"`
}
