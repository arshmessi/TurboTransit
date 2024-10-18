package main

import (
	"TurboTransit/auth-service/internal/controller"
	"TurboTransit/auth-service/internal/manager"

	"github.com/gin-gonic/gin"
)

func main() {
    tokenManager := manager.NewTokenManager("secret_key")
    controller := controller.NewAuthController(tokenManager)

    router := gin.Default()

    router.POST("/auth/login", controller.Login)
    router.POST("/auth/logout", controller.Logout)
    router.POST("/auth/refresh", controller.RefreshToken)

    router.Run(":8081")
}