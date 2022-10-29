package entities

import (
	"time"

	"gorm.io/gorm"
)

type Alliance struct {
	gorm.Model
	Name            string     `gorm:"unique;not null"`
	NameLocal       JSONMap    `gorm:"type:jsonb"`
	Description     JSONMap    `gorm:"type:jsonb"`
	DateDisbanded   *time.Time `gorm:"index"`
	AllianceMembers []AllianceMember
}

func (a *Alliance) BeforeCreate(tx *gorm.DB) (err error) {
	if a.NameLocal == nil || len(a.NameLocal) == 0 {
		a.NameLocal = make(JSONMap)
		a.NameLocal["en"] = a.Name
	}
	if a.Description == nil || len(a.Description) == 0 {
		a.Description = make(JSONMap)
		a.Description["en"] = ""
	}
	return
}
