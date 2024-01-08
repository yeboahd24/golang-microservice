package repository

import (
    "auth-service/model"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{
        db: db,
    }
}

// CreateUser saves a new user in the database.
func (repo *UserRepository) CreateUser(user *model.User) error {
    return repo.db.Create(user).Error
}

// FindUserByUsername finds a user by username.
func (repo *UserRepository) FindUserByUsername(username string) (*model.User, error) {
    var user model.User
    err := repo.db.Where("username = ?", username).First(&user).Error
    return &user, err
}

// FindUserByEmail finds a user by email.
func (repo *UserRepository) FindUserByEmail(email string) (*model.User, error) {
    var user model.User
    err := repo.db.Where("email = ?", email).First(&user).Error
    return &user, err
}


