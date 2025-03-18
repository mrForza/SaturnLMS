package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/profile_service/internal/config"
	"github.com/mrForza/SaturnLMS/profile_service/internal/routes"
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

	routes.RegisterAuthRouters(server)

	server.Run(":8080")
}
