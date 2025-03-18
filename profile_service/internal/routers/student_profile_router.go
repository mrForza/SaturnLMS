package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterStudentProfileRouters(baseRouter *gin.Engine) {
	studentProfileRouters := baseRouter.Group("/student-profile")
	{
		studentProfileRouters.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Read student profile"})
		})
		studentProfileRouters.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Add student profile"})
		})
		studentProfileRouters.DELETE("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNoContent, gin.H{"message": "Remove student profile"})
		})
	}
}
