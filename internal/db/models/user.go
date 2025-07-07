package models

import(
	"time"
	"gorm.io/gorm"
)


type UserRole string

const(
	RolePatient UserRole = "PATIENT"
	RolerAdmin UserRole = "ADMIN"
)


type User struct {
    ID        uint           `gorm:"primaryKey"`
    Name      string         `gorm:"not null"`
    Email     string         `gorm:"uniqueIndex;not null"`
    Password  string         `gorm:"not null"` // stored as hashed
    Role      UserRole       `gorm:"type:VARCHAR(20);default:'PATIENT'"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}