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

func GetAllAdminProfiles() ([]models.AdminProfile, error) {
	var adminProfiles []models.AdminProfile
	query := "SELECT id, education, work_experience, achievements, languages FROM administrative_profile;"

	err := Db.Select(&adminProfiles, query)
	if err != nil {
		return nil, err
	}

	return adminProfiles, nil
}

func GetAdminProfileById(id string) (*models.AdminProfile, error) {
	var AdminProfile models.AdminProfile

	query := `
        SELECT id, education, work_experience, achievements, languages
        FROM administrative_profile
        WHERE id = $1;
    `

	err := Db.Get(&AdminProfile, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("admin profile with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to fetch admin profile: %w", err)
	}

	return &AdminProfile, nil
}

func CreateAdminProfile(id uuid.UUID, AdminProfile dtos.CreateAdminProfileRequestDto) (uuid.UUID, error) {
	query := "INSERT INTO administrative_profile (id, education, work_experience, achievements, languages) VALUES ($1, $2, $3, $4, $5);"

	_, err := Db.Exec(
		query,
		id,
		AdminProfile.Education,
		AdminProfile.WorkExperience,
		AdminProfile.Achievements,
		AdminProfile.Languages,
	)

	if err != nil {
		log.Fatalf("%s", err.Error())
		return uuid.Nil, errors.New("the admin profile cannot be created because this profile has already existed")
	}

	return id, nil
}

func DeleteAdminById(id string) error {
	query := `
        DELETE FROM administrative_profile
        WHERE id = $1;
    `

	result, err := Db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete admin profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("admin profile with id %s not found", id)
	}

	fmt.Printf("admin profile with ID %s successfully deleted\n", id)
	return nil
}
