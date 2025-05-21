package usecases

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/study_service/internal/dal"
	"github.com/mrForza/SaturnLMS/study_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/study_service/internal/models"
	"github.com/mrForza/SaturnLMS/study_service/internal/validators"
)

func GetAllCourses() dtos.GetAllCoursesResponseDto {
	var Courses []models.Course
	Courses, _ = dal.GetAllCourses()

	return dtos.GetAllCoursesResponseDto{
		Courses: Courses,
	}
}

func GetCourseById(dto dtos.GetCourseByIdRequest) (dtos.GetCourseByIdResponse, error) {
	Course, err := dal.GetCourseById(dto.Id)

	if err != nil {
		return dtos.GetCourseByIdResponse{Course: nil}, err
	}

	return dtos.GetCourseByIdResponse{Course: Course}, nil
}

func CreateCourse(dto dtos.CreateCourseRequestDto) dtos.CreateCourseResponseDto {
	if err := validators.ValidateCourseInitials(dto.FirstName, "FirstName"); err != nil {
		return dtos.CreateCourseResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateCourseInitials(dto.LastName, "LastName"); err != nil {
		return dtos.CreateCourseResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateCourseInitials(dto.FatherName, "FatherName"); err != nil {
		return dtos.CreateCourseResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateAge(dto.Age); err != nil {
		return dtos.CreateCourseResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateGender(bool(dto.Gender)); err != nil {
		return dtos.CreateCourseResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateAboutMe(dto.AboutMe); err != nil {
		return dtos.CreateCourseResponseDto{Message: err.Error()}
	}

	if err := validators.ValidateInterests(dto.Interests); err != nil {
		return dtos.CreateCourseResponseDto{Message: err.Error()}
	}

	CourseId, err := dal.CreateCourse(uuid.New(), dto)
	if err != nil {
		return dtos.CreateCourseResponseDto{Message: err.Error()}
	}

	return dtos.CreateCourseResponseDto{
		Message: fmt.Sprintf("the Course with id %s has been seccessfully added", CourseId),
	}
}

func DeleteCourse(dto dtos.DeleteCourseByIdRequest) error {
	return dal.DeleteCourseById(dto.Id)
}
