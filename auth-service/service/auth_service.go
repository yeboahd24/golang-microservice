package service

import (
    "auth-service/model"
    "auth-service/repository"
    "golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt"
    "time"
)

type AuthenticationService struct {
    userRepo *repository.UserRepository
}

func NewAuthenticationService(userRepo *repository.UserRepository) *AuthenticationService {
    return &AuthenticationService{
        userRepo: userRepo,
    }
}

// CreateUser handles the creation of a new user, including password hashing.
func (s *AuthenticationService) CreateUser(user *model.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return s.userRepo.CreateUser(user)
}

// AuthenticateUser checks if the provided credentials are valid.
func (s *AuthenticationService) AuthenticateUser(username, password string) (bool, error) {
    user, err := s.userRepo.FindUserByUsername(username)
    if err != nil {
        return false, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    return err == nil, nil
}

// GenerateJWT generates a JWT token for authenticated users.
func (s *AuthenticationService) GenerateJWT(user *model.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(),
    })

    // Replace "YourSigningKey" with a secure key from configuration
    tokenString, err := token.SignedString([]byte("YourSigningKey"))
    return tokenString, err
}


