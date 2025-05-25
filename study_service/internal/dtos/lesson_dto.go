package dtos

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
)

type CreateLessonRequestDto struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Type        bool                    `json:"type"`
	Files       []*multipart.FileHeader `json:"files"`
	// Homework    uuid.UUID               `json:"homework"`
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

type UploadFileRequestDto struct {
	BucketId uuid.UUID               `json:"id"`
	Files    []*multipart.FileHeader `json:"file"`
}

type UploadFileResponseDto struct {
	Message string `json:"message"`
	IsError bool   `json:"is_error"`
}

type DeleteFileRequestDto struct {
	BucketId uuid.UUID `json:"id"`
	FileName string    `json:"file_name"`
}

type DeleteFileResponseDto struct {
	Message string `json:"message"`
	IsError bool   `json:"is_error"`
}

type DownloadFileRequestDto struct {
	BucketId uuid.UUID `json:"id"`
	FileName string    `json:"file_name"`
}

type DownloadFileResponseDto struct {
	File    *[]byte `json:"file"`
	IsError bool    `json:"is_error"`
}
