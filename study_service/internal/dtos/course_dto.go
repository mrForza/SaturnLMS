package dtos

import (
	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
)

type CreateCourseRequestDto struct {
	Name        string      `db:"name"`
	Description string      `db:"description"`
	Formula     string      `db:"formula"`
	Languages   []string    `db:"languages"`
	Teachers    []uuid.UUID `db:"teachers"`
	Students    []uuid.UUID `db:"students"`
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
