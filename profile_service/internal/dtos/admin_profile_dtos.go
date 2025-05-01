package dtos

import (
	"github.com/mrForza/SaturnLMS/profile_service/internal/models"
)

type CreateAdminProfileRequestDto struct {
	Education      string `json:"education"`
	WorkExperience string `json:"work_experience"`
	Achievements   string `json:"achievements"`
	Languages      string `json:"languages"`
}

type CreateAdminProfileResponseDto struct {
	Message string `json:"message"`
}

type GetAllAdminProfilesResponseDto struct {
	AdminProfiles []models.AdminProfile `json:"teacher_profiles"`
}

type GetAdminProfileByIdRequest struct {
	Id string `json:"id"`
}

type GetAdminProfileByIdResponse struct {
	AdminProfile *models.AdminProfile `json:"teacher_profile"`
}

type DeleteAdminProfileByIdRequest struct {
	Id string `json:"id"`
}
