package dtos

import (
	"github.com/mrForza/SaturnLMS/profile_service/internal/models"
)

type CreateUserRequestDto struct {
	FirstName  string        `json:"first_name"`
	LastName   string        `json:"last_name"`
	FatherName string        `json:"father_name"`
	Age        uint8         `json:"age"`
	Gender     models.Gender `json:"gender"`
	AboutMe    string        `json:"about_me"`
	Interests  string        `json:"interests"`
}

type CreateUserResponseDto struct {
	Message string `json:"message"`
}

type GetAllUserProfilesResponseDto struct {
	UserProfiles []models.UserProfile `json:"user_profiles"`
}

type GetUserProfileByIdRequest struct {
	Id string `json:"id"`
}

type GetUserProfileByIdResponse struct {
	UserProfile *models.UserProfile `json:"user_profile"`
}

type DeleteUserProfileByIdRequest struct {
	Id string `json:"id"`
}
