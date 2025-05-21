package usecases

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dal"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/models"
	"github.com/mrForza/SaturnLMS/profile_service/internal/validators"
)

func GetAllAdminProfiles() dtos.GetAllAdminProfilesResponseDto {
	var adminProfiles []models.AdminProfile
	adminProfiles, _ = dal.GetAllAdminProfiles()

	return dtos.GetAllAdminProfilesResponseDto{
		AdminProfiles: adminProfiles,
	}
}

func GetAdminProfileById(dto dtos.GetAdminProfileByIdRequest) (dtos.GetAdminProfileByIdResponse, error) {
	adminProfile, err := dal.GetAdminProfileById(dto.Id)

	if err != nil {
		return dtos.GetAdminProfileByIdResponse{AdminProfile: nil}, err
	}

	return dtos.GetAdminProfileByIdResponse{AdminProfile: adminProfile}, nil
}

func CreateAdminProfile(dto dtos.CreateAdminProfileRequestDto) (*string, error) {
	if err := validators.ValidateStringField(dto.Education, 1024, "education"); err != nil {
		return nil, err
	}
	if err := validators.ValidateStringField(dto.WorkExperience, 65536, "work experience"); err != nil {
		return nil, err
	}
	if err := validators.ValidateStringField(dto.Achievements, 65536, "achievements"); err != nil {
		return nil, err
	}
	if err := validators.ValidateStringField(dto.Languages, 128, "languages"); err != nil {
		return nil, err
	}

	userProfileId, err := dal.CreateAdminProfile(uuid.New(), dto)
	if err != nil {
		return nil, err
	}

	var message = fmt.Sprintf("the 'admin profile' with id %s has been seccessfully added", userProfileId)
	return &message, nil
}

func DeleteAdminProfile(dto dtos.DeleteAdminProfileByIdRequest) error {
	return dal.DeleteAdminById(dto.Id)
}
