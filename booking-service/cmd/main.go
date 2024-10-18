package main

import (
	"TurboTransit/booking-service/internal/controller"
	"TurboTransit/booking-service/internal/repository"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./bookings.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    repo := repository.NewBookingRepository(db)
    controller := controller.NewBookingController(repo)

    router := gin.Default()

    router.POST("/bookings", controller.CreateBooking)
    router.GET("/bookings/:id", controller.GetBooking)
    router.PUT("/bookings/:id/status", controller.UpdateBookingStatus)
    router.GET("/bookings/user/:userId", controller.GetUserBookings)

    router.Run(":8083")
}