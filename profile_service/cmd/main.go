package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/profile_service/internal/config"
	routes "github.com/mrForza/SaturnLMS/profile_service/internal/routers"
)

func main() {
	fmt.Println("Starting profile_service...")

	config := config.Load()
	fmt.Println(config)
	server := gin.Default()

	server.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	routes.RegisterHealthcheckRouter(server)
	routes.RegisterAuthRouters(server)
	routes.RegisterUserProfileRouters(server)
	routes.RegisterStudentProfileRouters(server)
	routes.RegisterTeacherProfileRouters(server)
	routes.RegisterAdminProfileRouters(server)

	server.Run(":8080")
}
