package dal

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mrForza/SaturnLMS/profile_service/internal/dtos"
	"github.com/mrForza/SaturnLMS/profile_service/internal/models"
)

func GetAllStudentProfiles() ([]models.StudentProfile, error) {
	var studentProfiles []models.StudentProfile
	query := "SELECT id, university_name, facultaty_name, program_name, group_number, course_number FROM student_profile;"

	err := Db.Select(&studentProfiles, query)
	if err != nil {
		return nil, err
	}

	return studentProfiles, nil
}

func GetStudentProfileById(id string) (*models.StudentProfile, error) {
	var StudentProfile models.StudentProfile

	query := `
        SELECT id, university_name, facultaty_name, program_name, group_number, course_number
        FROM student_profile
        WHERE id = $1;
    `

	err := Db.Get(&StudentProfile, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("student profile with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to fetch student profile: %w", err)
	}

	return &StudentProfile, nil
}

func CreateStudentProfile(id uuid.UUID, StudentProfile dtos.CreateStudentProfileRequestDto) (uuid.UUID, error) {
	query := "INSERT INTO student_profile (id, university_name, facultaty_name, program_name, group_number, course_number) VALUES ($1, $2, $3, $4, $5, $6);"

	_, err := Db.Exec(
		query,
		id,
		StudentProfile.UniversityName,
		StudentProfile.FacultatyName,
		StudentProfile.ProgramName,
		StudentProfile.GroupNumber,
		StudentProfile.CourseNumber,
	)

	if err != nil {
		log.Fatalf("%s", err.Error())
		return uuid.Nil, errors.New("the student profile cannot be created because this profile has already existed")
	}

	return id, nil
}

func DeleteStudentById(id string) error {
	query := `
        DELETE FROM student_profile
        WHERE id = $1;
    `

	result, err := Db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete student profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("student profile with id %s not found", id)
	}

	fmt.Printf("student profile with ID %s successfully deleted\n", id)
	return nil
}
