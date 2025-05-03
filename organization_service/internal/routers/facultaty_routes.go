package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/organization_service/internal/usecases"
)

func RegisterFacultatyRouters(baseRouter *gin.Engine) {
	facultatyRouters := baseRouter.Group("/facultaty")
	{
		facultatyRouters.GET("/", GetFacultaty)
		facultatyRouters.GET("/:name", GetFacultatyByName)
		facultatyRouters.POST("/", CreateFacultaty)
		facultatyRouters.DELETE("/:name", DeleteFacultaty)
	}
}

func GetFacultaty(ctx *gin.Context) {
	var response = usecases.GetAllFacultaties()
	ctx.JSON(http.StatusOK, response)
}

func GetFacultatyByName(ctx *gin.Context) {
	name := ctx.Param("name")
	fmt.Println(name)
	response, err := usecases.GetFacultatyByName(
		dtos.GetFacultatyByNameRequestDto{Name: name},
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateFacultaty(ctx *gin.Context) {
	var request dtos.CreateFacultatyRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "incorrect structure of request"})
		return
	}

	message, err := usecases.CreateFacultaty(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.CreateFacultatyResponseDto{Message: err.Error()})
	}

	ctx.JSON(http.StatusCreated, message)
}

func DeleteFacultaty(ctx *gin.Context) {
	name := ctx.Param("name")
	err := usecases.DeleteFacultaty(dtos.DeleteFacultatyRequestDto{Name: name})

	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
