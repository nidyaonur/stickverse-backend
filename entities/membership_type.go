package entities

import "gorm.io/gorm"

type MembershipType struct {
	gorm.Model
	Name        string            `gorm:"unique;not null"`
	NameLocal   map[string]string `gorm:"type:jsonb"`
	Description map[string]string `gorm:"type:jsonb"`
}

func (mt *MembershipType) BeforeCreate(tx *gorm.DB) (err error) {
	if mt.NameLocal == nil || len(mt.NameLocal) == 0 {
		mt.NameLocal = map[string]string{
			"en": mt.Name,
		}
	}
	if mt.Description == nil || len(mt.Description) == 0 {
		mt.Description = map[string]string{
			"en": "",
		}
	}
	return
}
