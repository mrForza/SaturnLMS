package dtos

import (
	"github.com/mrForza/SaturnLMS/organization_service/internal/models"
)

type CreateProgramRequestDto struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Languages    string `json:"languages"`
	FacultatyNme string `json:"facultaty_name"`
}

type CreateProgramResponseDto struct {
	Message string `json:"message"`
}

type GetProgramByNameRequestDto struct {
	Name string `json:"name"`
}

type GetProgramByNameResponseDto struct {
	Program *models.Program `json:"program"`
}

type GetAllProgramsResponseDto struct {
	Programs []models.Program `json:"programs"`
}

type DeleteProgramRequestDto struct {
	Name string `json:"name"`
}
