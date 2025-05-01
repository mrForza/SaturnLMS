package dtos

import (
	"github.com/mrForza/SaturnLMS/profile_service/internal/models"
)

type CreateStudentProfileRequestDto struct {
	UniversityName string `json:"university_name"`
	FacultatyName  string `json:"facultaty_name"`
	ProgramName    string `json:"program_name"`
	GroupNumber    uint8  `json:"group_number"`
	CourseNumber   uint8  `json:"course_number"`
}

type CreateStudentProfileResponseDto struct {
	Message string `json:"message"`
}

type GetAllStudentProfilesResponseDto struct {
	StudentProfiles []models.StudentProfile `json:"student_profiles"`
}

type GetStudentProfileByIdRequest struct {
	Id string `json:"id"`
}

type GetStudentProfileByIdResponse struct {
	StudentProfile *models.StudentProfile `json:"student_profile"`
}

type DeleteStudentProfileByIdRequest struct {
	Id string `json:"id"`
}
