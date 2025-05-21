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

func GetAllTeacherProfiles() ([]models.TeacherProfile, error) {
	var teacherProfiles []models.TeacherProfile
	query := "SELECT id, education, scientific_experience, teaching_experience, professional_interests, achievements, languages FROM teacher_profile;"

	err := Db.Select(&teacherProfiles, query)
	if err != nil {
		return nil, err
	}

	return teacherProfiles, nil
}

func GetTeacherProfileById(id string) (*models.TeacherProfile, error) {
	var TeacherProfile models.TeacherProfile

	query := `
        SELECT id, education, scientific_experience, teaching_experience, professional_interests, achievements, languages
        FROM teacher_profile
        WHERE id = $1;
    `

	err := Db.Get(&TeacherProfile, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("teacher profile with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to fetch teacher profile: %w", err)
	}

	return &TeacherProfile, nil
}

func CreateTeacherProfile(id uuid.UUID, TeacherProfile dtos.CreateTeacherProfileRequestDto) (uuid.UUID, error) {
	query := "INSERT INTO teacher_profile (id, education, scientific_experience, teaching_experience, professional_interests, achievements, languages) VALUES ($1, $2, $3, $4, $5, $6, $7);"

	_, err := Db.Exec(
		query,
		id,
		TeacherProfile.Education,
		TeacherProfile.ScientificExperience,
		TeacherProfile.TeachingExperience,
		TeacherProfile.ProfessionalInterests,
		TeacherProfile.Achievements,
		TeacherProfile.Languages,
	)

	if err != nil {
		log.Fatalf("%s", err.Error())
		return uuid.Nil, errors.New("the teacher profile cannot be created because this profile has already existed")
	}

	return id, nil
}

func DeleteTeacherById(id string) error {
	query := `
        DELETE FROM teacher_profile
        WHERE id = $1;
    `

	result, err := Db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete teacher profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("teacher profile with id %s not found", id)
	}

	fmt.Printf("teacher profile with ID %s successfully deleted\n", id)
	return nil
}
