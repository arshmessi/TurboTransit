package main

import (
	"TurboTransit/matching-service/internal/controller"
	"TurboTransit/matching-service/internal/engine"

	"github.com/gin-gonic/gin"
)

func main() {
    engine := engine.NewMatchingEngine()
    controller := controller.NewMatchingController(engine)

    router := gin.Default()

    router.POST("/matching/find-driver/:bookingId", controller.FindDriver)

    router.Run(":8084")
}