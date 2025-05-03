package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/organization_service/internal/usecases"
)

func RegisterProgramRouters(baseRouter *gin.Engine) {
	programRouters := baseRouter.Group("/program")
	{
		programRouters.GET("/", GetProgram)
		programRouters.GET("/:name", GetProgramByName)
		programRouters.POST("/", CreateProgram)
		programRouters.DELETE("/:name", DeleteProgram)
	}
}

func GetProgram(ctx *gin.Context) {
	var response = usecases.GetAllPrograms()
	ctx.JSON(http.StatusOK, response)
}

func GetProgramByName(ctx *gin.Context) {
	name := ctx.Param("name")
	fmt.Println(name)
	response, err := usecases.GetProgramByName(
		dtos.GetProgramByNameRequestDto{Name: name},
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateProgram(ctx *gin.Context) {
	var request dtos.CreateProgramRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "incorrect structure of request"})
		return
	}

	message, err := usecases.CreateProgram(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.CreateProgramResponseDto{Message: err.Error()})
	}

	ctx.JSON(http.StatusCreated, message)
}

func DeleteProgram(ctx *gin.Context) {
	name := ctx.Param("name")
	err := usecases.DeleteProgram(dtos.DeleteProgramRequestDto{Name: name})

	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
