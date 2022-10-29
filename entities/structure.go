package entities

import "gorm.io/gorm"

type Structure struct {
	gorm.Model
	Name               string         `gorm:"unique;not null"`
	NameLocal          JSONMap        `gorm:"type:jsonb"`
	Description        JSONMap        `gorm:"type:jsonb"`
	UpgradeTimeFormula string         `gorm:"not null"`
	Prerequisites      []Prerequisite `gorm:"foreignKey:StructureID"`
	Resources          []Resource     `gorm:"many2many:structure_resources"`
}

func (s *Structure) BeforeCreate(tx *gorm.DB) (err error) {
	if s.NameLocal == nil || len(s.NameLocal) == 0 {
		s.NameLocal = make(JSONMap)
		s.NameLocal["en"] = s.Name
	}
	if s.Description == nil || len(s.Description) == 0 {
		s.Description = make(JSONMap)
		s.Description["en"] = ""
	}
	return
}
