package service

import (
    "profile-service/model"
    "profile-service/repository"
)

// ProfileService provides methods to manage user profiles.
type ProfileService struct {
    profileRepo repository.ProfileRepository
}

// NewProfileService creates a new instance of ProfileService.
func NewProfileService(profileRepo repository.ProfileRepository) *ProfileService {
    return &ProfileService{
        profileRepo: profileRepo,
    }
}

// CreateProfile handles the creation of a new profile.
func (s *ProfileService) CreateProfile(profile *model.Profile) error {
    return s.profileRepo.CreateProfile(profile)
}

// UpdateProfile handles updating an existing profile.
func (s *ProfileService) UpdateProfile(profile *model.Profile) error {
    return s.profileRepo.UpdateProfile(profile)
}

// GetProfileByUserID retrieves a profile based on the user's ID.
func (s *ProfileService) GetProfileByUserID(userID uint) (*model.Profile, error) {
    return s.profileRepo.GetProfileByUserID(userID)
}


