package controller

import (
	"TurboTransit/pricing-service/internal/manager"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PricingController struct {
    calculator *manager.FareCalculator
}

func NewPricingController(calculator *manager.FareCalculator) *PricingController {
    return &PricingController{calculator: calculator}
}

func (c *PricingController) CalculateFare(ctx *gin.Context) {
    var req struct {
        Distance    float64       `json:"distance"`
        Duration    time.Duration `json:"duration"`
        VehicleType string        `json:"vehicle_type"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    fare, err := c.calculator.CalculateFare(req.Distance, req.Duration, req.VehicleType)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"fare": fare})
}