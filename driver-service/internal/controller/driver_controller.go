package controller

import (
	"TurboTransit/driver-service/internal/model"
	"TurboTransit/driver-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DriverController struct {
    repo *repository.DriverRepository
}

func NewDriverController(repo *repository.DriverRepository) *DriverController {
    return &DriverController{repo: repo}
}

func (c *DriverController) RegisterDriver(ctx *gin.Context) {
    var driver model.Driver
    if err := ctx.ShouldBindJSON(&driver); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.repo.CreateDriver(&driver); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, driver)
}

func (c *DriverController) GetDriver(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
        return
    }

    driver, err := c.repo.GetDriverByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
        return
    }

    ctx.JSON(http.StatusOK, driver)
}

func (c *DriverController) UpdateDriverStatus(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
        return
    }

    var req struct {
        Status string `json:"status"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.repo.UpdateDriverStatus(id, req.Status); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Driver status updated"})
}

func (c *DriverController) AssignVehicle(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
        return
    }

    var vehicle model.Vehicle
    if err := ctx.ShouldBindJSON(&vehicle); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.repo.AssignVehicle(id, &vehicle); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, vehicle)
}