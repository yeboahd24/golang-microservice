package api

import (
    "net/http"
    "user-service/model"
    "user-service/service"
    "github.com/gin-gonic/gin"
    "strconv"
)

// Handler struct will hold all the service dependencies needed for the handlers.
type Handler struct {
    userService *service.UserService
}

// NewHandler creates a new instance of the user API handler.
func NewHandler(userService *service.UserService) *Handler {
    return &Handler{
        userService: userService,
    }
}

// CreateUser handles the creation of a new user.
func (h *Handler) CreateUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.userService.CreateUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}

// GetUser handles retrieving a user by ID.
func (h *Handler) GetUser(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := h.userService.GetUserByID(uint(id))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}


func (h *Handler) UpdateUser(c *gin.Context) {
  id, err := strconv.ParseUint(c.Param("id"), 10, 32)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
    return
  }

  var user model.User
  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  user.ID = uint(id)

  if err := h.userService.UpdateUser(&user); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}


func (h *Handler) GetUserName(c *gin.Context) {
  username := c.Param("username")
  user, err := h.userService.GetUserByUsername(username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, user)
}

  
// RegisterRoutes sets up the routes for the handler.
func (h *Handler) RegisterRoutes(router *gin.Engine) {
    router.POST("/users", h.CreateUser)
    router.GET("/users/:id", h.GetUser)
    router.PUT("/users/:id", h.UpdateUser)
    router.GET("/users/:username", h.GetUserName)
}

