package main

import (
	"TurboTransit/common/events"
	"TurboTransit/common/nats"
	"TurboTransit/common/redis"
	"TurboTransit/user-service/internal/controller"
	"TurboTransit/user-service/internal/repository"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
    nats.InitNATS()
    redis.InitRedis()

    db, err := sql.Open("sqlite3", "./users.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    repo := repository.NewUserRepository(db)
    controller := controller.NewUserController(repo)

    // Subscribe to events
    events.SubscribeToEvent("BookingCreated", func(data []byte) {
        // Handle BookingCreated event
    })

    router := gin.Default()

    router.POST("/users", controller.RegisterUser)
    router.GET("/users/:id", controller.GetUser)
    router.PUT("/users/:id", controller.UpdateUser)

    router.Run(":8080")
}