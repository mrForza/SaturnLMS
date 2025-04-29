package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(baseRouter *gin.Engine) {
	baseRouter.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})

	RegisterHealthcheckRouter(baseRouter)

	RegisterUserProfileRouters(baseRouter)

	RegisterStudentProfileRouters(baseRouter)

	RegisterTeacherProfileRouters(baseRouter)

	RegisterAdminProfileRouters(baseRouter)
}
