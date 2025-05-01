package usecases

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dal"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/models"
	"github.com/mrForza/SaturnLMS/profile_service/internal/validators"
)

func GetAllStudentProfiles() dtos.GetAllStudentProfilesResponseDto {
	var studentProfiles []models.StudentProfile
	studentProfiles, _ = dal.GetAllStudentProfiles()

	return dtos.GetAllStudentProfilesResponseDto{
		StudentProfiles: studentProfiles,
	}
}

func GetStudentProfileById(dto dtos.GetStudentProfileByIdRequest) (dtos.GetStudentProfileByIdResponse, error) {
	studentProfile, err := dal.GetStudentProfileById(dto.Id)

	if err != nil {
		return dtos.GetStudentProfileByIdResponse{StudentProfile: nil}, err
	}

	return dtos.GetStudentProfileByIdResponse{StudentProfile: studentProfile}, nil
}

func CreateStudentProfile(dto dtos.CreateStudentProfileRequestDto) (*string, error) {
	if err := validators.ValidateUniversityName(dto.UniversityName); err != nil {
		return nil, err
	}

	if err := validators.ValidateFacultatyName(dto.FacultatyName); err != nil {
		return nil, err
	}

	if err := validators.ValidateProgramName(dto.ProgramName); err != nil {
		return nil, err
	}

	if err := validators.ValidateCourseNumber(dto.CourseNumber); err != nil {
		return nil, err
	}

	if err := validators.ValidateGroupNumber(dto.GroupNumber); err != nil {
		return nil, err
	}

	userProfileId, err := dal.CreateStudentProfile(uuid.New(), dto)
	if err != nil {
		return nil, err
	}

	var message = fmt.Sprintf("the 'student profile' with id %s has been seccessfully added", userProfileId)
	return &message, nil
}

func DeleteStudentProfile(dto dtos.DeleteStudentProfileByIdRequest) error {
	return dal.DeleteStudentById(dto.Id)
}
