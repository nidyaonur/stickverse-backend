package entities

import "gorm.io/gorm"

type Characteristic struct {
	gorm.Model
	Name        string            `gorm:"unique;not null"`
	NameLocal   map[string]string `gorm:"type:jsonb"`
	Description map[string]string `gorm:"type:jsonb"`
}

func (c *Characteristic) BeforeCreate(tx *gorm.DB) (err error) {
	if c.NameLocal == nil || len(c.NameLocal) == 0 {
		c.NameLocal = map[string]string{
			"en": c.Name,
		}
	}
	if c.Description == nil || len(c.Description) == 0 {
		c.Description = map[string]string{
			"en": "",
		}
	}
	return
}
