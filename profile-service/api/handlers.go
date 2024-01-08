package api

import (
    "net/http"
    "profile-service/model"
    "profile-service/service"
    "github.com/gin-gonic/gin"
    "strconv"
)

// Handler holds the profile service to interact with user profiles.
type Handler struct {
    profileService *service.ProfileService
}

// NewHandler creates a new API handler with the given profile service.
func NewHandler(profileService *service.ProfileService) *Handler {
    return &Handler{
        profileService: profileService,
    }
}

// CreateProfile handles the creation of a new profile.
func (h *Handler) CreateProfile(c *gin.Context) {
    var profile model.Profile
    if err := c.ShouldBindJSON(&profile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.profileService.CreateProfile(&profile); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Profile created successfully", "profile": profile})
}

// UpdateProfile handles updating an existing profile.
func (h *Handler) UpdateProfile(c *gin.Context) {
    userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var profile model.Profile
    if err := c.ShouldBindJSON(&profile); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    profile.UserID = uint(userID)
    if err := h.profileService.UpdateProfile(&profile); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "profile": profile})
}

// GetProfile retrieves a profile by user ID.
func (h *Handler) GetProfile(c *gin.Context) {
    userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    profile, err := h.profileService.GetProfileByUserID(uint(userID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, profile)
}

// RegisterRoutes sets up the routes for profile handling.
func (h *Handler) RegisterRoutes(router *gin.Engine) {
    router.POST("/profiles", h.CreateProfile)
    router.PUT("/profiles/:userID", h.UpdateProfile)
    router.GET("/profiles/:userID", h.GetProfile)
}
