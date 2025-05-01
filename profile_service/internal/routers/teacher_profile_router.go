package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/usecases"
)

func RegisterTeacherProfileRouters(baseRouter *gin.Engine) {
	teacherProfileRouters := baseRouter.Group("/teacher-profile")
	{
		teacherProfileRouters.GET("/", GetAllTeacherProfilesRoute)
		teacherProfileRouters.GET("/:id", GetTeacherProfileById)
		teacherProfileRouters.POST("/", CreateTeacherProfileRoute)
		teacherProfileRouters.PATCH("/:id", UpdateTeacherProfileById)
		teacherProfileRouters.DELETE("/:id", DeleteTeacherProfileRoute)
	}
}

func GetAllTeacherProfilesRoute(ctx *gin.Context) {
	var response = usecases.GetAllTeacherProfiles()
	ctx.JSON(http.StatusOK, response)
}

func GetTeacherProfileById(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := usecases.GetTeacherProfileById(
		dtos.GetTeacherProfileByIdRequest{Id: id},
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateTeacherProfileRoute(ctx *gin.Context) {
	var request dtos.CreateTeacherProfileRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "incorrect structure of request"})
		return
	}

	message, err := usecases.CreateTeacherProfile(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.CreateTeacherProfileResponseDto{Message: err.Error()})
	}
	ctx.JSON(http.StatusCreated, dtos.CreateTeacherProfileResponseDto{Message: *message})
}

func UpdateTeacherProfileById(ctx *gin.Context) {

}

func DeleteTeacherProfileRoute(ctx *gin.Context) {
	id := ctx.Param("id")
	err := usecases.DeleteTeacherProfile(dtos.DeleteTeacherProfileByIdRequest{Id: id})

	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
