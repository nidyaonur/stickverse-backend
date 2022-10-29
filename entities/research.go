package entities

import "gorm.io/gorm"

type Research struct {
	gorm.Model
	Name               string            `gorm:"not null"`
	NameLocal          map[string]string `gorm:"type:jsonb"`
	Description        map[string]string `gorm:"type:jsonb"`
	UpgradeTimeFormula string            `gorm:"not null"`
	Resources          []Resource        `gorm:"many2many:research_resources"`
}

func (r *Research) BeforeCreate(tx *gorm.DB) (err error) {
	if r.NameLocal == nil || len(r.NameLocal) == 0 {
		r.NameLocal = map[string]string{
			"en": r.Name,
		}
	}
	if r.Description == nil || len(r.Description) == 0 {
		r.Description = map[string]string{
			"en": "",
		}
	}
	return
}
