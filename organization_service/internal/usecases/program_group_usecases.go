package usecases

import (
	"fmt"

	"github.com/mrForza/SaturnLMS/organization_service/internal/dal"
	"github.com/mrForza/SaturnLMS/organization_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/organization_service/internal/models"
	"github.com/mrForza/SaturnLMS/organization_service/internal/validators"
)

func GetAllProgramGroups() dtos.GetAllProgramGroupsResponseDto {
	var programGroups []models.ProgramGroup
	programGroups, _ = dal.GetAllProgramGroups()

	return dtos.GetAllProgramGroupsResponseDto{
		ProgramGroups: programGroups,
	}
}

func GetProgramGroupByNumber(dto dtos.GetProgramGroupByNumberRequestDto) (dtos.GetProgramGroupByNumberResponseDto, error) {
	programGroup, err := dal.GetProgramGroupByNumber(dto.Number)

	if err != nil {
		return dtos.GetProgramGroupByNumberResponseDto{ProgramGroup: nil}, err
	}

	return dtos.GetProgramGroupByNumberResponseDto{ProgramGroup: programGroup}, nil
}

func CreateProgramGroup(dto dtos.CreateProgramGroupRequestDto) (*dtos.CreateProgramGroupResponseDto, error) {
	if err := validators.ValidateStringField(dto.Name, 64, "name"); err != nil {
		return nil, err
	}

	if err := validators.ValidateCourseNumber(dto.CourseNumber); err != nil {
		return nil, err
	}

	_, err := dal.GetProgramByName(dto.ProgramName)
	if err != nil {
		return nil, fmt.Errorf("no program with name: %s", dto.ProgramName)
	}

	programGroupName, err := dal.CreateProgramGroup(dto)
	if err != nil {
		return nil, err
	}

	return &dtos.CreateProgramGroupResponseDto{
		Message: fmt.Sprintf("the programGroup with name %s has been seccessfully added", programGroupName),
	}, nil
}

func DeleteProgramGroup(dto dtos.DeleteProgramGroupRequestDto) error {
	return dal.DeleteProgramGroupByName(dto.Number)
}
