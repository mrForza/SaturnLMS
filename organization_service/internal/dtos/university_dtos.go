package dtos

import (
	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/organization_service/internal/models"
)

type CreateUniversityRequestDto struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	LegalAddress  string    `json:"legal_address"`
	ActualAddress string    `json:"actual_address"`
	Inn           string    `json:"inn"`
	BankName      string    `json:"bank_name"`
	OwnerId       uuid.UUID `json:"owner_id"`
}

type CreateUniversityResponseDto struct {
	Message string `json:"message"`
}

type GetUniversityByNameRequestDto struct {
	Name string `json:"name"`
}

type GetUniversityByNameResponseDto struct {
	University *models.University `json:"university"`
}

type GetAllUniversitiesResponseDto struct {
	Universities []models.University `json:"universities"`
}

type DeleteUniversityRequestDto struct {
	Name string `json:"name"`
}
