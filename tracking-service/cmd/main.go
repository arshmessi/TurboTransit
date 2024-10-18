package main

import (
	"TurboTransit/tracking-service/internal/controller"
	"TurboTransit/tracking-service/internal/repository"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./locations.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    repo := repository.NewTrackingRepository(db)
    controller := controller.NewTrackingController(repo)

    router := gin.Default()

    router.POST("/locations", controller.UpdateLocation)
    router.GET("/locations/driver/:driverId", controller.GetDriverLocation)
    router.GET("/locations/booking/:bookingId", controller.GetBookingLocations)

    router.Run(":8085")
}