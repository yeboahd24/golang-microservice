package service

import (
    "user-service/model"
    "user-service/repository"
    "golang.org/x/crypto/bcrypt"
)

// UserService provides methods to manage users.
type UserService struct {
    userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService(userRepo repository.UserRepository) *UserService {
    return &UserService{
        userRepo: userRepo,
    }
}

// CreateUser handles the creation of a new user, including password hashing.
func (s *UserService) CreateUser(user *model.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    return s.userRepo.CreateUser(user)
}

// UpdateUser handles updating an existing user.
func (s *UserService) UpdateUser(user *model.User) error {
    return s.userRepo.UpdateUser(user)
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
    return s.userRepo.GetUserByID(id)
}

// GetUserByUsername retrieves a user by their username.
func (s *UserService) GetUserByUsername(username string) (*model.User, error) {
    return s.userRepo.GetUserByUsername(username)
}
