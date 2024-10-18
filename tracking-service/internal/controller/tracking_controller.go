package controller

import (
	"TurboTransit/tracking-service/internal/model"
	"TurboTransit/tracking-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrackingController struct {
    repo *repository.TrackingRepository
}

func NewTrackingController(repo *repository.TrackingRepository) *TrackingController {
    return &TrackingController{repo: repo}
}

func (c *TrackingController) UpdateLocation(ctx *gin.Context) {
    var location model.Location
    if err := ctx.ShouldBindJSON(&location); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.repo.UpdateLocation(&location); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, location)
}

func (c *TrackingController) GetDriverLocation(ctx *gin.Context) {
    driverID, err := strconv.Atoi(ctx.Param("driverId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
        return
    }

    location, err := c.repo.GetDriverLocation(driverID)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
        return
    }

    ctx.JSON(http.StatusOK, location)
}

func (c *TrackingController) GetBookingLocations(ctx *gin.Context) {
    bookingID, err := strconv.Atoi(ctx.Param("bookingId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
        return
    }

    locations, err := c.repo.GetBookingLocations(bookingID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, locations)
}