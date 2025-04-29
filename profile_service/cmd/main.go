package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/profile_service/internal/configs"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dal"
	routers "github.com/mrForza/SaturnLMS/profile_service/internal/routers"
)

func main() {
	fmt.Println("Starting profile_service...")

	fmt.Println(*configs.ApiConfig)
	fmt.Println(*configs.DbConfig)

	dal.InitDB()

	server := gin.Default()

	routers.InitRouters(server)

	server.Run(":8080")
}
