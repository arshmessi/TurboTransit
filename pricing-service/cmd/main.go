package main

import (
	"TurboTransit/pricing-service/internal/controller"
	"TurboTransit/pricing-service/internal/manager"

	"github.com/gin-gonic/gin"
)

func main() {
    calculator := manager.NewFareCalculator()
    controller := controller.NewPricingController(calculator)

    router := gin.Default()

    router.POST("/pricing/calculate", controller.CalculateFare)

    router.Run(":8086")
}