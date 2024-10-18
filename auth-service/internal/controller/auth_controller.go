package controller

import (
	"TurboTransit/auth-service/internal/manager"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
    tokenManager *manager.TokenManager
}

func NewAuthController(tokenManager *manager.TokenManager) *AuthController {
    return &AuthController{tokenManager: tokenManager}
}

func (c *AuthController) Login(ctx *gin.Context) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validate username and password (replace with actual validation logic)
    userID := 1 // Replace with actual user ID retrieval

    token, err := c.tokenManager.GenerateToken(userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, token)
}

func (c *AuthController) Logout(ctx *gin.Context) {
    // Implement logout logic here
    ctx.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
    var req struct {
        Token string `json:"token"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID, err := c.tokenManager.ValidateToken(req.Token)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    newToken, err := c.tokenManager.GenerateToken(userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, newToken)
}