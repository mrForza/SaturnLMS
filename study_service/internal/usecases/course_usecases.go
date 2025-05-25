package usecases

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/study_service/internal/dal"
	"github.com/mrForza/SaturnLMS/study_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
)

var courseRepository = dal.NewCourseRepository(dal.Client.Database("study_db"))

func GetAllCourses() dtos.GetAllCoursesResponseDto {
	var Courses []models.Course
	Courses, _ = courseRepository.GetAllCourses()

	return dtos.GetAllCoursesResponseDto{
		Courses: Courses,
	}
}

func GetCourseById(dto dtos.GetCourseByIdRequest) (dtos.GetCourseByIdResponse, error) {
	Course, err := courseRepository.GetCourseById(dto.Id)

	if err != nil {
		return dtos.GetCourseByIdResponse{Course: nil}, err
	}

	return dtos.GetCourseByIdResponse{Course: Course}, nil
}

func CreateCourse(dto dtos.CreateCourseRequestDto) dtos.CreateCourseResponseDto {
	var course = models.Course{
		Id:          uuid.New(),
		Name:        dto.Name,
		Description: dto.Description,
		Formula:     dto.Formula,
		Languages:   dto.Languages,
		Teachers:    dto.Teachers,
		Students:    dto.Students,
		Lessons:     dto.Lessons,
	}
	CourseId, err := courseRepository.CreateCourse(uuid.New(), course)
	if err != nil {
		return dtos.CreateCourseResponseDto{Message: err.Error()}
	}

	return dtos.CreateCourseResponseDto{
		Message: fmt.Sprintf("the Course with id %s has been seccessfully added", CourseId),
	}
}

func DeleteCourse(dto dtos.DeleteCourseByIdRequest) error {
	return courseRepository.DeleteCourseById(dto.Id)
}
