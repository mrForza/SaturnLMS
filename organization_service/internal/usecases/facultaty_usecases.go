package usecases

import (
	"fmt"

	"github.com/mrForza/SaturnLMS/organization_service/internal/dal"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/organization_service/internal/models"
	"github.com/mrForza/SaturnLMS/organization_service/internal/validators"
)

func GetAllFacultaties() dtos.GetAllFacultatiesResponseDto {
	var facultaties []models.Facultaty
	facultaties, _ = dal.GetAllFacultaties()

	return dtos.GetAllFacultatiesResponseDto{
		Facultaties: facultaties,
	}
}

func GetFacultatyByName(dto dtos.GetFacultatyByNameRequestDto) (dtos.GetFacultatyByNameResponseDto, error) {
	facultaty, err := dal.GetFacultatyByName(dto.Name)

	if err != nil {
		return dtos.GetFacultatyByNameResponseDto{Facultaty: nil}, err
	}

	return dtos.GetFacultatyByNameResponseDto{Facultaty: facultaty}, nil
}

func CreateFacultaty(dto dtos.CreateFacultatyRequestDto) (*dtos.CreateFacultatyResponseDto, error) {
	if err := validators.ValidateStringField(dto.Name, 64, "name"); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringField(dto.Description, 65536, "description"); err != nil {
		return nil, err
	}

	_, err := dal.GetUniversityByName(dto.UniversityName)
	if err != nil {
		return nil, fmt.Errorf("no university with name: %s", dto.UniversityName)
	}

	facultatyName, err := dal.CreateFacultaty(dto)
	if err != nil {
		return nil, err
	}

	return &dtos.CreateFacultatyResponseDto{
		Message: fmt.Sprintf("the facultaty with name %s has been seccessfully added", facultatyName),
	}, nil
}

func DeleteFacultaty(dto dtos.DeleteFacultatyRequestDto) error {
	return dal.DeleteFacultatyByName(dto.Name)
}
