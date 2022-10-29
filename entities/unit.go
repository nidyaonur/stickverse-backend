package entities

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	Name        string  `gorm:"unique;not null"`
	NameLocal   JSONMap `gorm:"type:jsonb"`
	Description JSONMap `gorm:"type:jsonb"`
}

func (u *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	if u.NameLocal == nil || len(u.NameLocal) == 0 {
		u.NameLocal = make(JSONMap)
		u.NameLocal["en"] = u.Name
	}
	if u.Description == nil || len(u.Description) == 0 {
		u.Description = make(JSONMap)
		u.Description["en"] = ""
	}
	return
}
