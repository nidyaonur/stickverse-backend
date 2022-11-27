package entities

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name           string             `gorm:"not null"`
	NameLocal      JSONMap            `gorm:"type:jsonb"`
	LocationTypeID uint               `gorm:"not null"`
	LocationType   LocationType       `gorm:"foreignkey:LocationTypeID"`
	Level          int                `gorm:"not null"`
	Workers        uint64             `gorm:"null"`
	UserID         *uint              `gorm:"null"`
	User           *User              `gorm:"foreignkey:UserID"`
	GridID         uint               `gorm:"uniqueIndex:grid_id_index;null"`
	Grid           Grid               `gorm:"foreignkey:GridID"`
	GridIndex      int                `gorm:"uniqueIndex:grid_id_index;null"`
	Units          []LocationUnit     `gorm:"foreignkey:LocationID"`
	Resources      []LocationResource `gorm:"foreignkey:LocationID"`
	Structures     []StructureBuilt   `gorm:"foreignkey:LocationID"`
}

func (l *Location) BeforeCreate(tx *gorm.DB) (err error) {
	if l.NameLocal == nil || len(l.NameLocal) == 0 {
		l.NameLocal = make(JSONMap)
		l.NameLocal["en"] = l.Name
	}
	return
}
