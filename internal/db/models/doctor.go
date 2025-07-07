package models

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Specialty string `gorm:"not null"`
	Bio       string `gorm:"type:TEXT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
