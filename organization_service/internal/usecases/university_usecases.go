package usecases

import (
	"fmt"

	"github.com/mrForza/SaturnLMS/organization_service/internal/connectors"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dal"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/organization_service/internal/models"
	"github.com/mrForza/SaturnLMS/organization_service/internal/validators"
)

func GetAllUniversities() dtos.GetAllUniversitiesResponseDto {
	var universities []models.University
	universities, _ = dal.GetAllUniversities()

	return dtos.GetAllUniversitiesResponseDto{
		Universities: universities,
	}
}

func GetUniversityByName(dto dtos.GetUniversityByNameRequestDto) (dtos.GetUniversityByNameResponseDto, error) {
	university, err := dal.GetUniversityByName(dto.Name)

	if err != nil {
		return dtos.GetUniversityByNameResponseDto{University: nil}, err
	}

	return dtos.GetUniversityByNameResponseDto{University: university}, nil
}

func CreateUniversityProfile(dto dtos.CreateUniversityRequestDto) (*dtos.CreateUniversityResponseDto, error) {
	if err := validators.ValidateStringField(dto.Name, 64, "name"); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringField(dto.Description, 65536, "description"); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringField(dto.LegalAddress, 512, "legal address"); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringField(dto.ActualAddress, 512, "actual address"); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringField(dto.Inn, 10, "inn"); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringField(dto.BankName, 64, "bank name"); err != nil {
		return nil, err
	}

	if !connectors.DoesAdminProfileExist(dto.OwnerId) {
		return nil, fmt.Errorf("no admin profile with id: %s", dto.OwnerId)
	}

	universityName, err := dal.CreateUniversity(dto)
	if err != nil {
		return nil, err
	}

	return &dtos.CreateUniversityResponseDto{
		Message: fmt.Sprintf("the UniversityProfile with id %s has been seccessfully added", universityName),
	}, nil
}

func DeleteUniversityProfile(dto dtos.DeleteUniversityRequestDto) error {
	return dal.DeleteUniversityByName(dto.Name)
}
