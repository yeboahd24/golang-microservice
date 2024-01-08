package main

import (
    "log"
    "profile-service/api"
    "profile-service/model"
    "profile-service/repository"
    "profile-service/service"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "profile-service/pkg/config"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize database connection
    db, err := gorm.Open(postgres.Open(cfg.GetDBConnectionString()), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Auto migrate our model
    if err := db.AutoMigrate(&model.Profile{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Initialize repository
    profileRepo := repository.NewProfileRepository(db)

    // Initialize service
    profileService := service.NewProfileService(profileRepo)

    // Initialize Gin router and setup routes
    router := gin.Default()
    handler := api.NewHandler(profileService)
    handler.RegisterRoutes(router)

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

