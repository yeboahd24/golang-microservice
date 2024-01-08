package main

import (
    "auth-service/api"
    "auth-service/model"
    "auth-service/repository"
    "auth-service/service"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "auth-service/pkg/config"
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

    // Auto migrate our models
    if err := db.AutoMigrate(&model.User{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }

    // Initialize repository
    userRepo := repository.NewUserRepository(db)

    // Initialize service
    authService := service.NewAuthenticationService(userRepo)

    // Initialize Gin router and setup routes
    router := gin.Default()
    handler := api.NewHandler(authService)

    // Setup routes
    router.POST("/register", handler.RegisterUser)
    router.POST("/login", handler.LoginUser)

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

