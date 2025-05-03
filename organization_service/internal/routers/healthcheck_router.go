package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHealthcheckRouter(baseRouter *gin.Engine) {
	healthcheckRouter := baseRouter.Group("/healthcheck")
	{
		healthcheckRouter.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
		})
	}
}
