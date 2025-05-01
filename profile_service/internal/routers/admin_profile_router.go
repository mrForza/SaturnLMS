package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/usecases"
)

func RegisterAdminProfileRouters(baseRouter *gin.Engine) {
	adminProfileRouters := baseRouter.Group("/admin-profile")
	{
		adminProfileRouters.GET("/", GetAllAdminProfilesRoute)
		adminProfileRouters.GET("/:id", GetAdminProfileById)
		adminProfileRouters.POST("/", CreateAdminProfileRoute)
		adminProfileRouters.PATCH("/:id", UpdateAdminProfileById)
		adminProfileRouters.DELETE("/:id", DeleteAdminProfileRoute)
	}
}

func GetAllAdminProfilesRoute(ctx *gin.Context) {
	var response = usecases.GetAllAdminProfiles()
	ctx.JSON(http.StatusOK, response)
}

func GetAdminProfileById(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := usecases.GetAdminProfileById(
		dtos.GetAdminProfileByIdRequest{Id: id},
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateAdminProfileRoute(ctx *gin.Context) {
	var request dtos.CreateAdminProfileRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "incorrect structure of request"})
		return
	}

	message, err := usecases.CreateAdminProfile(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.CreateAdminProfileResponseDto{Message: err.Error()})
	}
	ctx.JSON(http.StatusCreated, dtos.CreateAdminProfileResponseDto{Message: *message})
}

func UpdateAdminProfileById(ctx *gin.Context) {

}

func DeleteAdminProfileRoute(ctx *gin.Context) {
	id := ctx.Param("id")
	err := usecases.DeleteAdminProfile(dtos.DeleteAdminProfileByIdRequest{Id: id})

	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
