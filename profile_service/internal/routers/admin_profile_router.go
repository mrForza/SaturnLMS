package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdminProfileRouters(baseRouter *gin.Engine) {
	adminProfileRouters := baseRouter.Group("/admin-profile")
	{
		adminProfileRouters.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "Read admin profile"})
		})
		adminProfileRouters.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{"message": "Add admin profile"})
		})
		adminProfileRouters.DELETE("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusNoContent, gin.H{"message": "Remove admin profile"})
		})
	}
}
