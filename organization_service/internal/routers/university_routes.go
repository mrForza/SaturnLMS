package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/organization_service/internal/usecases"
)

func RegisterUniversityProfileRouters(baseRouter *gin.Engine) {
	universityRouters := baseRouter.Group("/university")
	{
		universityRouters.GET("/", GetUniversity)
		universityRouters.GET("/:name", GetUniversityByName)
		universityRouters.POST("/", CreateUniversity)
		universityRouters.DELETE("/:name", DeleteUniversity)
	}
}

func GetUniversity(ctx *gin.Context) {
	var response = usecases.GetAllUniversities()
	ctx.JSON(http.StatusOK, response)
}

func GetUniversityByName(ctx *gin.Context) {
	name := ctx.Param("name")
	fmt.Println(name)
	response, err := usecases.GetUniversityByName(
		dtos.GetUniversityByNameRequestDto{Name: name},
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateUniversity(ctx *gin.Context) {
	var request dtos.CreateUniversityRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "incorrect structure of request"})
		return
	}

	message, err := usecases.CreateUniversityProfile(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.CreateUniversityResponseDto{Message: err.Error()})
	}

	ctx.JSON(http.StatusCreated, message)
}

func DeleteUniversity(ctx *gin.Context) {
	name := ctx.Param("name")
	err := usecases.DeleteUniversityProfile(dtos.DeleteUniversityRequestDto{Name: name})

	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
