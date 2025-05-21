package dtos

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
)

type CreateLessonRequestDto struct {
	Name        string           `db:"name"`
	Description string           `db:"description"`
	Type        bool             `db:"type"`
	Files       []multipart.File `db:"files"`
	Homework    uuid.UUID        `db:"homework"`
}

type CreateLessonResponseDto struct {
	Message string `json:"message"`
}

type GetAllLessonsResponseDto struct {
	Lessons []models.Lesson `json:"lessons"`
}

type GetLessonByIdRequest struct {
	Id string `json:"id"`
}

type GetLessonByIdResponse struct {
	Lesson *models.Lesson `json:"lesson"`
}

type DeleteLessonByIdRequest struct {
	Id string `json:"id"`
}
