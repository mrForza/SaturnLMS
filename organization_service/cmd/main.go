package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/organization_service/internal/configs"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dal"
	"github.com/mrForza/SaturnLMS/organization_service/internal/routers"
)

func main() {
	fmt.Println("Starting organization_service...")

	fmt.Println(*configs.ApiConfig)
	fmt.Println(*configs.DbConfig)

	dal.InitDB()

	server := gin.Default()

	routers.InitRouters(server)

	server.Run(":8081")
}
