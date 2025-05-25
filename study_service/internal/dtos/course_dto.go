package dtos

import (
	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
)

type CreateCourseRequestDto struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Formula     string      `json:"formula"`
	Languages   []string    `json:"languages"`
	Teachers    []uuid.UUID `json:"teachers"`
	Students    []uuid.UUID `json:"students"`
	Lessons     []uuid.UUID `json:"lessons"`
}

type CreateCourseResponseDto struct {
	Message string `json:"message"`
}

type GetAllCoursesResponseDto struct {
	Courses []models.Course `json:"courses"`
}

type GetCourseByIdRequest struct {
	Id string `json:"id"`
}

type GetCourseByIdResponse struct {
	Course *models.Course `json:"course"`
}

type DeleteCourseByIdRequest struct {
	Id string `json:"id"`
}
