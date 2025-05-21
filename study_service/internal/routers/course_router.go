package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/usecases"
)

func RegisterCourseRouter(baseRouter *gin.Engine) {
	CourseRouters := baseRouter.Group("/courses")
	{
		CourseRouters.GET("/", GetAllCoursesRoute)
		CourseRouters.GET("/:id", GetCourseById)
		CourseRouters.POST("/", CreateCourseRoute)
		CourseRouters.PATCH("/:id", UpdateCourseById)
		CourseRouters.DELETE("/:id", DeleteCourseRoute)
	}
}

func GetAllCoursesRoute(ctx *gin.Context) {
	var response = usecases.GetAllCourses()
	ctx.JSON(http.StatusOK, response)
}

func GetCourseById(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := usecases.GetCourseById(
		dtos.GetCourseByIdRequest{Id: id},
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateCourseRoute(ctx *gin.Context) {
	var request dtos.CreateUserRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "incorrect structure of request"})
		return
	}

	ctx.JSON(http.StatusCreated, usecases.CreateCourse(request))
}

func UpdateCourseById(ctx *gin.Context) {

}

func DeleteCourseRoute(ctx *gin.Context) {
	id := ctx.Param("id")
	err := usecases.DeleteCourse(dtos.DeleteCourseByIdRequest{Id: id})

	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
