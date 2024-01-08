package model

import (
    "gorm.io/gorm"
    "time"
)

// User represents the user's personal and authentication information in the system.
type User struct {
    gorm.Model
    Username    string     `gorm:"uniqueIndex;not null"`
    Email       string     `gorm:"uniqueIndex;not null"`
    Password    string     `gorm:"not null"` // Password should be hashed
    FirstName   string
    LastName    string
    DateOfBirth time.Time
}

