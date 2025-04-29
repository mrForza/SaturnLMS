package usecases

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dal"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/models"
	"github.com/mrForza/SaturnLMS/profile_service/internal/validators"
)

func GetAllUserProfiles() dtos.GetAllUserProfilesResponseDto {
	var userProfiles []models.UserProfile
	userProfiles, _ = dal.GetAllUserProfiles()

	return dtos.GetAllUserProfilesResponseDto{
		UserProfiles: userProfiles,
	}
}

func GetUserProfileById(dto dtos.GetUserProfileByIdRequest) (dtos.GetUserProfileByIdResponse, error) {
	userProfile, err := dal.GetUserProfileById(dto.Id)

	if err != nil {
		return dtos.GetUserProfileByIdResponse{UserProfile: nil}, err
	}

	return dtos.GetUserProfileByIdResponse{UserProfile: userProfile}, nil
}

func CreateUserProfile(dto dtos.CreateUserRequestDto) dtos.CreateUserResponseDto {
	if err := validators.ValidateUserInitials(dto.FirstName, "FirstName"); err != nil {
		return dtos.CreateUserResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateUserInitials(dto.LastName, "LastName"); err != nil {
		return dtos.CreateUserResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateUserInitials(dto.FatherName, "FatherName"); err != nil {
		return dtos.CreateUserResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateAge(dto.Age); err != nil {
		return dtos.CreateUserResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateGender(bool(dto.Gender)); err != nil {
		return dtos.CreateUserResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateAboutMe(dto.AboutMe); err != nil {
		return dtos.CreateUserResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateInterests(dto.Interests); err != nil {
		return dtos.CreateUserResponseDto{Message: err.Error()}
	}

	userProfileId, err := dal.CreateUserProfile(uuid.New(), dto)
	if err != nil {
		return dtos.CreateUserResponseDto{Message: err.Error()}
	}

	return dtos.CreateUserResponseDto{
		Message: fmt.Sprintf("the UserProfile with id %s has been seccessfully added", userProfileId),
	}
}

func DeleteUserProfile(dto dtos.DeleteUserProfileByIdRequest) error {
	return dal.DeleteUserById(dto.Id)
}
