package dtos

import (
	"github.com/mrForza/SaturnLMS/organization_service/internal/models"
)

type CreateFacultatyRequestDto struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	UniversityName string `json:"university_name"`
}

type CreateFacultatyResponseDto struct {
	Message string `json:"message"`
}

type GetFacultatyByNameRequestDto struct {
	Name string `json:"name"`
}

type GetFacultatyByNameResponseDto struct {
	Facultaty *models.Facultaty `json:"facultaty"`
}

type GetAllFacultatiesResponseDto struct {
	Facultaties []models.Facultaty `json:"facultaties"`
}

type DeleteFacultatyRequestDto struct {
	Name string `json:"name"`
}
