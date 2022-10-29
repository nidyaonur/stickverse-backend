package entities

import (
	"time"

	"gorm.io/gorm"
)

type LoginHistory struct {
	gorm.Model
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignkey:UserID"`
	login_time  time.Time `gorm:"not null"`
	logout_time time.Time `gorm:"null"`
	IP          string    `gorm:"not null"`
}
