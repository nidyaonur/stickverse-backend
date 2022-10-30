package entities

import "gorm.io/gorm"

type Characteristic struct {
	gorm.Model
	Name        string  `gorm:"unique;not null"`
	NameLocal   JSONMap `gorm:"type:jsonb"`
	Description JSONMap `gorm:"type:jsonb"`
}

func (c *Characteristic) BeforeCreate(tx *gorm.DB) (err error) {
	if c.NameLocal == nil || len(c.NameLocal) == 0 {
		c.NameLocal = make(JSONMap)
		c.NameLocal["en"] = c.Name
	}
	if c.Description == nil || len(c.Description) == 0 {
		c.Description = make(JSONMap)
		c.Description["en"] = ""
	}
	return
}
