package entities

import "gorm.io/gorm"

type StructureResource struct {
	gorm.Model
	StructureID uint `gorm:"index"`
	ResourceID  uint `gorm:"index"`

	UpgradeFormula    string `gorm:"type:text"`
	ProductionFormula string `gorm:"type:text"`

	Structure Structure `gorm:"foreignkey:StructureID"`
	Resource  Resource  `gorm:"foreignkey:ResourceID"`
}
