package dtos

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
)

type CreateHomeworkRequestDto struct {
	Id          uuid.UUID              `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Files       []multipart.FileHeader `json:"files"`
}

type CreateHomeworkResponseDto struct {
	Message string `json:"message"`
}

type GetAllHomeworksResponseDto struct {
	Homeworks []models.Homework `json:"homeworks"`
}

type GetHomeworkByIdRequest struct {
	Id string `json:"id"`
}

type GetHomeworkByIdResponse struct {
	Homework *models.Homework `json:"homework"`
}

type DeleteHomeworkByIdRequest struct {
	Id string `json:"id"`
}
