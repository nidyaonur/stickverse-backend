package entities

import "gorm.io/gorm"

type GridType struct {
	gorm.Model
	Name        string  `gorm:"unique;not null"`
	Description JSONMap `gorm:"type:jsonb"`
	NameLocal   JSONMap `gorm:"type:jsonb"`
}

func (gt *GridType) BeforeCreate(tx *gorm.DB) (err error) {
	if gt.NameLocal == nil || len(gt.NameLocal) == 0 {
		gt.NameLocal = make(JSONMap)
		gt.NameLocal["en"] = gt.Name
	}
	if gt.Description == nil || len(gt.Description) == 0 {
		gt.Description = make(JSONMap)
		gt.Description["en"] = ""
	}
	return
}
