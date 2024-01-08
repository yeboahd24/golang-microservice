package repository

import (
    "profile-service/model"
    "gorm.io/gorm"
)

// ProfileRepository is the interface for the database operations for profiles.
type ProfileRepository interface {
    CreateProfile(profile *model.Profile) error
    UpdateProfile(profile *model.Profile) error
    GetProfileByUserID(userID uint) (*model.Profile, error)
}

type profileRepository struct {
    db *gorm.DB
}

// NewProfileRepository creates a new instance of ProfileRepository.
func NewProfileRepository(db *gorm.DB) ProfileRepository {
    return &profileRepository{
        db: db,
    }
}

// CreateProfile adds a new profile to the database.
func (repo *profileRepository) CreateProfile(profile *model.Profile) error {
    return repo.db.Create(profile).Error
}

// UpdateProfile updates an existing profile in the database.
func (repo *profileRepository) UpdateProfile(profile *model.Profile) error {
    return repo.db.Save(profile).Error
}

// GetProfileByUserID retrieves a profile by the user's ID.
func (repo *profileRepository) GetProfileByUserID(userID uint) (*model.Profile, error) {
    var profile model.Profile
    err := repo.db.Where("user_id = ?", userID).First(&profile).Error
    return &profile, err
}

