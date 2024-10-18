package controller

import (
	"TurboTransit/booking-service/internal/model"
	"TurboTransit/booking-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
    repo *repository.BookingRepository
}

func NewBookingController(repo *repository.BookingRepository) *BookingController {
    return &BookingController{repo: repo}
}

func (c *BookingController) CreateBooking(ctx *gin.Context) {
    var booking model.Booking
    if err := ctx.ShouldBindJSON(&booking); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.repo.CreateBooking(&booking); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, booking)
}

func (c *BookingController) GetBooking(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
        return
    }

    booking, err := c.repo.GetBookingByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
        return
    }

    ctx.JSON(http.StatusOK, booking)
}

func (c *BookingController) UpdateBookingStatus(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
        return
    }

    var req struct {
        Status string `json:"status"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.repo.UpdateBookingStatus(id, req.Status); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Booking status updated"})
}

func (c *BookingController) GetUserBookings(ctx *gin.Context) {
    userID, err := strconv.Atoi(ctx.Param("userId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    bookings, err := c.repo.GetUserBookings(userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, bookings)
}