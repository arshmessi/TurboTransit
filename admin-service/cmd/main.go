package main

import (
	"TurboTransit/admin-service/internal/controller"

	"github.com/gin-gonic/gin"
)

func main() {
    controller := controller.NewAdminController()

    router := gin.Default()

    router.GET("/admin/dashboard", controller.GetDashboard)
    router.POST("/admin/users/ban", controller.BanUser)
    router.POST("/admin/drivers/suspend", controller.SuspendDriver)
    router.GET("/admin/reports/bookings", controller.GenerateBookingReport)

    router.Run(":8087")
}