package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/study_service/internal/configs"
	"github.com/mrForza/SaturnLMS/study_service/internal/dal"
	"github.com/mrForza/SaturnLMS/study_service/internal/routers"
)

func main() {
	fmt.Println("Starting profile_service...")

	fmt.Println(*configs.ApiConfig)
	fmt.Println(*configs.DbConfig)

	dal.ConnectToMongo("mongodb://admin:secret@studyService.mongo:27017")

	server := gin.Default()

	routers.InitRouters(server)

	server.Run(":8083")
}
