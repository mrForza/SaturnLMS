package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRouters(baseRouter *gin.Engine) {
	authRouters := baseRouter.Group("/auth")
	{
		authRouters.POST("signup", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{"message": "successful registration"})
		})
		authRouters.POST("login", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{"message": "successful authentication"})
		})
		authRouters.POST("logout", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{"message": "successful logout"})
		})
		authRouters.POST("reset", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{"message": "successful passwordd reset"})
		})
	}
}
