package usecases

import (
	"fmt"

	"github.com/mrForza/SaturnLMS/organization_service/internal/dal"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/organization_service/internal/models"
	"github.com/mrForza/SaturnLMS/organization_service/internal/validators"
)

func GetAllPrograms() dtos.GetAllProgramsResponseDto {
	var programs []models.Program
	programs, _ = dal.GetAllPrograms()

	return dtos.GetAllProgramsResponseDto{
		Programs: programs,
	}
}

func GetProgramByName(dto dtos.GetProgramByNameRequestDto) (dtos.GetProgramByNameResponseDto, error) {
	program, err := dal.GetProgramByName(dto.Name)

	if err != nil {
		return dtos.GetProgramByNameResponseDto{Program: nil}, err
	}

	return dtos.GetProgramByNameResponseDto{Program: program}, nil
}

func CreateProgram(dto dtos.CreateProgramRequestDto) (*dtos.CreateProgramResponseDto, error) {
	if err := validators.ValidateStringField(dto.Name, 64, "name"); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringField(dto.Description, 65536, "description"); err != nil {
		return nil, err
	}

	if err := validators.ValidateType(dto.Type); err != nil {
		return nil, err
	}

	if err := validators.ValidateStringField(dto.Description, 64, "languages"); err != nil {
		return nil, err
	}

	_, err := dal.GetFacultatyByName(dto.FacultatyNme)
	if err != nil {
		return nil, fmt.Errorf("no facultaty with name: %s", dto.FacultatyNme)
	}

	programName, err := dal.CreateProgram(dto)
	if err != nil {
		return nil, err
	}

	return &dtos.CreateProgramResponseDto{
		Message: fmt.Sprintf("the program with name %s has been seccessfully added", programName),
	}, nil
}

func DeleteProgram(dto dtos.DeleteProgramRequestDto) error {
	return dal.DeleteProgramByName(dto.Name)
}
