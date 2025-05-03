package dtos

import (
	"github.com/mrForza/SaturnLMS/organization_service/internal/models"
)

type CreateProgramGroupRequestDto struct {
	Number       uint16 `json:"number"`
	Name         string `json:"name"`
	CourseNumber uint8  `json:"course_number"`
	ProgramName  string `json:"program_name"`
}

type CreateProgramGroupResponseDto struct {
	Message string `json:"message"`
}

type GetProgramGroupByNumberRequestDto struct {
	Number uint16 `json:"number"`
}

type GetProgramGroupByNumberResponseDto struct {
	ProgramGroup *models.ProgramGroup `json:"group"`
}

type GetAllProgramGroupsResponseDto struct {
	ProgramGroups []models.ProgramGroup `json:"groups"`
}

type DeleteProgramGroupRequestDto struct {
	Number uint16 `json:"number"`
}
