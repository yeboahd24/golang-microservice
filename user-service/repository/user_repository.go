package repository

import (
    "user-service/model"
    "gorm.io/gorm"
)

// UserRepository defines the interface for user data operations.
type UserRepository interface {
    CreateUser(user *model.User) error
    UpdateUser(user *model.User) error
    GetUserByID(id uint) (*model.User, error)
    GetUserByUsername(username string) (*model.User, error)
}

type userRepository struct {
    db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

// CreateUser adds a new user to the database.
func (repo *userRepository) CreateUser(user *model.User) error {
    return repo.db.Create(user).Error
}

// UpdateUser updates an existing user in the database.
func (repo *userRepository) UpdateUser(user *model.User) error {
    return repo.db.Save(user).Error
}

// GetUserByID retrieves a user by their ID.
func (repo *userRepository) GetUserByID(id uint) (*model.User, error) {
    var user model.User
    err := repo.db.First(&user, id).Error
    return &user, err
}

// GetUserByUsername retrieves a user by their username.
func (repo *userRepository) GetUserByUsername(username string) (*model.User, error) {
    var user model.User
    err := repo.db.Where("username = ?", username).First(&user).Error
    return &user, err
}


