package model

import (
    "gorm.io/gorm"
)

// Profile represents a user's profile in the system.
type Profile struct {
    gorm.Model
    UserID      uint      // Foreign key to the User model
    Bio         string
    PictureURL  string
    ContactInfo string
}

