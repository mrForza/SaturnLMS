package usecases

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/study_service/internal/dal"
	"github.com/mrForza/SaturnLMS/study_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
)

var lessonRepository = dal.NewLessonRepository(dal.Client.Database("study_db"))

func CreateLesson(dto dtos.CreateLessonRequestDto) dtos.CreateLessonResponseDto {
	id := uuid.New()

	err := dal.CreateLessonBucketByUUID(dal.MinioClient, id)
	if err != nil {
		return dtos.CreateLessonResponseDto{
			Message: "Error while upload a file",
		}
	}

	var _, err2 = dal.UploadFilesToLessonBucket(dal.MinioClient, id, dto.Files)
	if err2 != nil {
		return dtos.CreateLessonResponseDto{
			Message: "Error while upload a file",
		}
	}

	lesson := models.Lesson{}
	id, err = lessonRepository.CreateLesson(id, lesson)
	if err != nil {
		return dtos.CreateLessonResponseDto{
			Message: err.Error(),
		}
	}

	return dtos.CreateLessonResponseDto{
		Message: fmt.Sprintf("The lesson with id: %s was successfully added", id),
	}
}

func DeleteLesson(dto dtos.DeleteLessonByIdRequest) error {
	return lessonRepository.DeleteLessonById(dto.Id)
}

func UploadFileIntoLessonById(dto dtos.UploadFileRequestDto) dtos.UploadFileResponseDto {
	err := dal.UploadFileToLesson(dal.MinioClient, dto.BucketId.String(), dto.Files[0])
	if err != nil {
		return dtos.UploadFileResponseDto{Message: err.Error(), IsError: true}
	}

	return dtos.UploadFileResponseDto{Message: "The file was successfully added into Lesson", IsError: false}
}

func DeleteFileFromLessonById(dto dtos.DeleteFileRequestDto) dtos.DeleteFileResponseDto {
	err := dal.DeleteFileFromMinIO(dal.MinioClient, dto.BucketId.String(), dto.FileName)
	if err != nil {
		return dtos.DeleteFileResponseDto{Message: err.Error(), IsError: true}
	}

	return dtos.DeleteFileResponseDto{Message: "The file was successfully removed from Lesson", IsError: false}
}

func DownloadFileFromLessonById(dto dtos.DownloadFileRequestDto) dtos.DownloadFileResponseDto {
	file, err := dal.DownloadFileFromMinIO(dal.MinioClient, dto.BucketId.String(), dto.FileName)
	if err != nil {
		return dtos.DownloadFileResponseDto{File: nil, IsError: true}
	}

	return dtos.DownloadFileResponseDto{File: &file, IsError: false}
}
