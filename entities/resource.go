package entities

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	Name        string  `gorm:"unique;not null"`
	NameLocal   JSONMap `gorm:"type:jsonb"`
	Description JSONMap `gorm:"type:jsonb"`
}

func (r *Resource) BeforeCreate(tx *gorm.DB) (err error) {
	if r.NameLocal == nil || len(r.NameLocal) == 0 {
		r.NameLocal = make(JSONMap)
		r.NameLocal["en"] = r.Name
	}
	if r.Description == nil || len(r.Description) == 0 {
		r.Description = make(JSONMap)
		r.Description["en"] = ""
	}
	return
}
