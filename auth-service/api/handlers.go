package api

import (
    "auth-service/model"
    "auth-service/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

type Handler struct {
    authService *service.AuthenticationService
}

func NewHandler(authService *service.AuthenticationService) *Handler {
    return &Handler{
        authService: authService,
    }
}

// RegisterUser handles user registration.
func (h *Handler) RegisterUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if err := h.authService.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// LoginUser handles user login.
func (h *Handler) LoginUser(c *gin.Context) {
    var loginDetails struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&loginDetails); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    authenticated, err := h.authService.AuthenticateUser(loginDetails.Username, loginDetails.Password)
    if err != nil || !authenticated {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
        return
    }

    token, err := h.authService.GenerateJWT(&model.User{Username: loginDetails.Username})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}


