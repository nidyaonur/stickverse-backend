package entities

import "gorm.io/gorm"

type Prerequisite struct {
	gorm.Model
	Type                string `gorm:"uniqueIndex:idx_unique_preq; type:varchar(255);not null"`
	PrerequisiteFormula string `gorm:"uniqueIndex:idx_unique_preq; type:varchar(255);null"`

	// Foreign keys
	StructureID *uint `gorm:"uniqueIndex:idx_unique_preq; null"`
	ResearchID  *uint `gorm:"null"`
	UnitID      *uint `gorm:"null"`

	// Required for prerequisites
	RequiredStructureID *uint `gorm:"null"`
	RequiredResearchID  *uint `gorm:"null"`
	RequiredLevel       *uint `gorm:"null"`

	Structure         *Structure `gorm:"foreignkey:StructureID"`
	Research          *Research  `gorm:"foreignkey:ResearchID"`
	Unit              *Unit      `gorm:"foreignkey:UnitID"`
	RequiredStructure *Structure `gorm:"foreignkey:RequiredStructureID"`
	RequiredResearch  *Research  `gorm:"foreignkey:RequiredResearchID"`
}
