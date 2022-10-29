package entities

import (
	"time"

	"gorm.io/gorm"
)

type GroupMovement struct {
	gorm.Model
	UserID         uint          `gorm:"not null"`
	MovementTypeID uint          `gorm:"not null"`
	LocationFromID uint          `gorm:"not null"`
	LocationToID   uint          `gorm:"not null"`
	ArrivalDate    time.Time     `gorm:"not null"`
	ReturnDate     time.Time     `gorm:"not null"`
	WaitingTime    time.Duration `gorm:"not null"`
	Comment        string        `gorm:"not null"`
	MovementType   MovementType  `gorm:"foreignkey:MovementTypeID"`
	LocationFrom   Location      `gorm:"foreignkey:LocationFromID"`
	LocationTo     Location      `gorm:"foreignkey:LocationToID"`
	Units          []Unit        `gorm:"many2many:group_movement_units;"`
	Resources      []GroupMovementResource
}
