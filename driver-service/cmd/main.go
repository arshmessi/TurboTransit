package main

import (
	"TurboTransit/driver-service/internal/controller"
	"TurboTransit/driver-service/internal/repository"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "./drivers.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    repo := repository.NewDriverRepository(db)
    controller := controller.NewDriverController(repo)

    router := gin.Default()

    router.POST("/drivers", controller.RegisterDriver)
    router.GET("/drivers/:id", controller.GetDriver)
    router.PUT("/drivers/:id/status", controller.UpdateDriverStatus)
    router.POST("/drivers/:id/vehicles", controller.AssignVehicle)

    router.Run(":8082")
}