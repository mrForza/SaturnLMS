package usecases

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dal"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/models"
	"github.com/mrForza/SaturnLMS/profile_service/internal/validators"
)

func GetAllTeacherProfiles() dtos.GetAllTeacherProfilesResponseDto {
	var teacherProfiles []models.TeacherProfile
	teacherProfiles, _ = dal.GetAllTeacherProfiles()

	return dtos.GetAllTeacherProfilesResponseDto{
		TeacherProfiles: teacherProfiles,
	}
}

func GetTeacherProfileById(dto dtos.GetTeacherProfileByIdRequest) (dtos.GetTeacherProfileByIdResponse, error) {
	teacherProfile, err := dal.GetTeacherProfileById(dto.Id)

	if err != nil {
		return dtos.GetTeacherProfileByIdResponse{TeacherProfile: nil}, err
	}

	return dtos.GetTeacherProfileByIdResponse{TeacherProfile: teacherProfile}, nil
}

func CreateTeacherProfile(dto dtos.CreateTeacherProfileRequestDto) (*string, error) {
	if err := validators.ValidateStringField(dto.Education, 1024, "education"); err != nil {
		return nil, err
	}
	if err := validators.ValidateStringField(dto.ScientificExperience, 65536, "scientific experience"); err != nil {
		return nil, err
	}
	if err := validators.ValidateStringField(dto.TeachingExperience, 65536, "teaching experience"); err != nil {
		return nil, err
	}
	if err := validators.ValidateStringField(dto.ProfessionalInterests, 65536, "professional interests"); err != nil {
		return nil, err
	}
	if err := validators.ValidateStringField(dto.Achievements, 65536, "achievements"); err != nil {
		return nil, err
	}
	if err := validators.ValidateStringField(dto.Languages, 128, "languages"); err != nil {
		return nil, err
	}

	userProfileId, err := dal.CreateTeacherProfile(uuid.New(), dto)
	if err != nil {
		return nil, err
	}

	var message = fmt.Sprintf("the 'teacher profile' with id %s has been seccessfully added", userProfileId)
	return &message, nil
}

func DeleteTeacherProfile(dto dtos.DeleteTeacherProfileByIdRequest) error {
	return dal.DeleteTeacherById(dto.Id)
}
