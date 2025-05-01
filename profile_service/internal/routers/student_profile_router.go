package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/usecases"
)

func RegisterStudentProfileRouters(baseRouter *gin.Engine) {
	studentProfileRouters := baseRouter.Group("/student-profile")
	{
		studentProfileRouters.GET("/", GetAllStudentProfilesRoute)
		studentProfileRouters.GET("/:id", GetStudentProfileById)
		studentProfileRouters.POST("/", CreateStudentProfileRoute)
		studentProfileRouters.PATCH("/:id", UpdateStudentProfileById)
		studentProfileRouters.DELETE("/:id", DeleteStudentProfileRoute)
	}
}

func GetAllStudentProfilesRoute(ctx *gin.Context) {
	var response = usecases.GetAllStudentProfiles()
	ctx.JSON(http.StatusOK, response)
}

func GetStudentProfileById(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := usecases.GetStudentProfileById(
		dtos.GetStudentProfileByIdRequest{Id: id},
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateStudentProfileRoute(ctx *gin.Context) {
	var request dtos.CreateStudentProfileRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "incorrect structure of request"})
		return
	}

	message, err := usecases.CreateStudentProfile(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.CreateStudentProfileResponseDto{Message: err.Error()})
	}
	ctx.JSON(http.StatusCreated, dtos.CreateStudentProfileResponseDto{Message: *message})
}

func UpdateStudentProfileById(ctx *gin.Context) {

}

func DeleteStudentProfileRoute(ctx *gin.Context) {
	id := ctx.Param("id")
	err := usecases.DeleteStudentProfile(dtos.DeleteStudentProfileByIdRequest{Id: id})

	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
