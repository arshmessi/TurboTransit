package controller

import (
	"TurboTransit/matching-service/internal/engine"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MatchingController struct {
    engine *engine.MatchingEngine
}

func NewMatchingController(engine *engine.MatchingEngine) *MatchingController {
    return &MatchingController{engine: engine}
}

func (c *MatchingController) FindDriver(ctx *gin.Context) {
    bookingID, err := strconv.Atoi(ctx.Param("bookingId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
        return
    }

    match, err := c.engine.FindDriver(bookingID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, match)
}