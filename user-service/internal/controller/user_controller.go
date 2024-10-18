package controller

import (
	"TurboTransit/user-service/internal/model"
	"TurboTransit/user-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    repo *repository.UserRepository
}

func NewUserController(repo *repository.UserRepository) *UserController {
    return &UserController{repo: repo}
}

func (c *UserController) RegisterUser(ctx *gin.Context) {
    var user model.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.repo.CreateUser(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetUser(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := c.repo.GetUserByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var user model.User
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.ID = id
    if err := c.repo.UpdateUser(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, user)
}