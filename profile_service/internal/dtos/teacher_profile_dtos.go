package dtos

import (
	"github.com/mrForza/SaturnLMS/profile_service/internal/models"
)

type CreateTeacherProfileRequestDto struct {
	Education             string `json:"education"`
	ScientificExperience  string `json:"scientific_experience"`
	TeachingExperience    string `json:"teaching_experience"`
	ProfessionalInterests string `json:"professional_interests"`
	Achievements          string `json:"achievements"`
	Languages             string `json:"languages"`
}

type CreateTeacherProfileResponseDto struct {
	Message string `json:"message"`
}

type GetAllTeacherProfilesResponseDto struct {
	TeacherProfiles []models.TeacherProfile `json:"teacher_profiles"`
}

type GetTeacherProfileByIdRequest struct {
	Id string `json:"id"`
}

type GetTeacherProfileByIdResponse struct {
	TeacherProfile *models.TeacherProfile `json:"teacher_profile"`
}

type DeleteTeacherProfileByIdRequest struct {
	Id string `json:"id"`
}
