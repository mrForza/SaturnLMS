package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterTeacherProfileRouters(baseRouter *gin.Engine) {
	teacherProfileRouters := baseRouter.Group("/teacher-profile")
	{
		teacherProfileRouters.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Read teacher profile"})
		})
		teacherProfileRouters.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Add teacher profile"})
		})
		teacherProfileRouters.DELETE("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNoContent, gin.H{"message": "Remove teacher profile"})
		})
	}
}
