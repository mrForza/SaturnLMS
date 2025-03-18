package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserProfileRouters(baseRouter *gin.Engine) {
	userProfileRouters := baseRouter.Group("/user-profile")
	{
		userProfileRouters.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Read user profile"})
		})
		userProfileRouters.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Add user profile"})
		})
		userProfileRouters.DELETE("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNoContent, gin.H{"message": "Remove user profile"})
		})
	}
}
