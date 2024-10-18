package main

import (
	"TurboTransit/api-gateway/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Apply middleware
    router.Use(middleware.AuthMiddleware())

    // User Service Routes
    router.POST("/users", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "User registration"})
    })
    router.GET("/users/:id", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "User profile retrieval"})
    })
    router.PUT("/users/:id", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "User profile update"})
    })

    // Auth Service Routes
    router.POST("/auth/login", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "User login"})
    })
    router.POST("/auth/logout", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "User logout"})
    })
    router.POST("/auth/refresh", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Token refresh"})
    })

    // Driver Service Routes
    router.POST("/drivers", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Driver registration"})
    })
    router.GET("/drivers/:id", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Driver profile retrieval"})
    })
    router.PUT("/drivers/:id/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Driver status update"})
    })
    router.POST("/drivers/:id/vehicles", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Vehicle assignment"})
    })

    router.Run(":8080")
}