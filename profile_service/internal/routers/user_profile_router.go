package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/usecases"
)

func RegisterUserProfileRouters(baseRouter *gin.Engine) {
	userProfileRouters := baseRouter.Group("/user-profile")
	{
		userProfileRouters.GET("/", GetAllUserProfilesRoute)
		userProfileRouters.GET("/:id", GetUserProfileById)
		userProfileRouters.POST("/", CreateUserProfileRoute)
		userProfileRouters.PATCH("/:id", UpdateUserProfileById)
		userProfileRouters.DELETE("/:id", DeleteUserProfileRoute)
	}
}

func GetAllUserProfilesRoute(ctx *gin.Context) {
	var response = usecases.GetAllUserProfiles()
	ctx.JSON(http.StatusOK, response)
}

func GetUserProfileById(ctx *gin.Context) {
	id := ctx.Param("id")
	response, err := usecases.GetUserProfileById(
		dtos.GetUserProfileByIdRequest{Id: id},
	)

	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateUserProfileRoute(ctx *gin.Context) {
	var request dtos.CreateUserRequestDto
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": "incorrect structure of request"})
		return
	}

	ctx.JSON(http.StatusCreated, usecases.CreateUserProfile(request))
}

func UpdateUserProfileById(ctx *gin.Context) {

}

func DeleteUserProfileRoute(ctx *gin.Context) {
	id := ctx.Param("id")
	err := usecases.DeleteUserProfile(dtos.DeleteUserProfileByIdRequest{Id: id})

	if err == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
}
