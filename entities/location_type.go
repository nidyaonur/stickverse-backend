package entities

import (
	"time"

	"gorm.io/gorm"
)

type LocationType struct {
	gorm.Model
	Name        string  `gorm:"index:,unique"`
	NameLocal   JSONMap `gorm:"type:jsonb"`
	Description JSONMap `gorm:"type:jsonb"`
}

func (lt *LocationType) BeforeCreate(tx *gorm.DB) (err error) {
	lt.CreatedAt = time.Now()
	lt.UpdatedAt = lt.CreatedAt
	if lt.NameLocal == nil || len(lt.NameLocal) == 0 {
		lt.NameLocal = make(JSONMap)
		lt.NameLocal["en"] = lt.Name
	}

	if lt.Description == nil || len(lt.Description) == 0 {
		lt.Description = make(JSONMap)
		lt.Description["en"] = ""
	}
	return
}
