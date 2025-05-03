package routers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/organization_service/internal/usecases"
)

func RegisterProgramGroupRouters(baseRouter *gin.Engine) {
	programGroupRouters := baseRouter.Group("/group")
	{
		programGroupRouters.GET("/", GetProgramGroup)
		programGroupRouters.GET("/:number", GetProgramGroupByNumber)
		programGroupRouters.POST("/", CreateProgramGroup)
		programGroupRouters.DELETE("/:number", DeleteProgramGroup)
	}
}

func GetProgramGroup(ctx *gin.Context) {
	var response = usecases.GetAllProgramGroups()
	ctx.JSON(http.StatusOK, response)
}

func GetProgramGroupByNumber(ctx *gin.Context) {
	number := ctx.Param("number")
	value, err := strconv.ParseUint(number, 10, 16)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "number should be a number"})
		return
	}

	response, err := usecases.GetProgramGroupByNumber(
		dtos.GetProgramGroupByNumberRequestDto{Number: uint16(value)},
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateProgramGroup(ctx *gin.Context) {
	var request dtos.CreateProgramGroupRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "incorrect structure of request"})
		return
	}

	message, err := usecases.CreateProgramGroup(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.CreateProgramGroupResponseDto{Message: err.Error()})
	}

	ctx.JSON(http.StatusCreated, message)
}

func DeleteProgramGroup(ctx *gin.Context) {
	number := ctx.Param("number")
	value, err := strconv.ParseUint(number, 10, 16)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "number should be a number"})
		return
	}

	err = usecases.DeleteProgramGroup(dtos.DeleteProgramGroupRequestDto{Number: uint16(value)})

	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
