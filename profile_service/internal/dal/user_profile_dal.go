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

func GetAllUserProfiles() ([]models.UserProfile, error) {
	var userProfiles []models.UserProfile
	query := "SELECT id, first_name, last_name, father_name, age, gender, about_me, interests FROM user_profile;"

	err := Db.Select(&userProfiles, query)
	if err != nil {
		return nil, err // errors.New("failed to fetch users")
	}

	// log.Fatal(userProfiles)
	return userProfiles, nil
}

func GetUserProfileById(id string) (*models.UserProfile, error) {
	var userProfile models.UserProfile

	query := `
        SELECT id, first_name, last_name, father_name, age, gender, about_me, interests
        FROM user_profile
        WHERE id = $1;
    `

	err := Db.Get(&userProfile, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to fetch user profile: %w", err)
	}

	return &userProfile, nil
}

func CreateUserProfile(id uuid.UUID, userProfile dtos.CreateUserRequestDto) (uuid.UUID, error) {
	query := "INSERT INTO user_profile (id, first_name, last_name, father_name, age, gender, about_me, interests) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);"

	_, err := Db.Exec(
		query,
		id,
		userProfile.FirstName,
		userProfile.LastName,
		userProfile.FatherName,
		userProfile.Age,
		userProfile.Gender,
		userProfile.AboutMe,
		userProfile.Interests,
	)

	if err != nil {
		log.Fatalf("%s", err.Error())
		return uuid.Nil, errors.New("the UserProfile cannot be created because this user has already existed")
	}

	return id, nil
}

func DeleteUserById(id string) error {
	query := `
        DELETE FROM user_profile
        WHERE id = $1;
    `

	result, err := Db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	fmt.Printf("User with ID %s successfully deleted\n", id)
	return nil
}

func UpdateUserProfileById(id uuid.UUID, updates map[string]interface{}) error {
	query := "UPDATE user_profile SET "
	var args []interface{}
	i := 1

	for key, value := range updates {
		query += fmt.Sprintf("%s = $%d, ", key, i)
		args = append(args, value)
		i++
	}

	query = query[:len(query)-2] + " WHERE id = $" + fmt.Sprint(i)
	args = append(args, id)

	result, err := Db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user profile: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	fmt.Printf("User with ID %s successfully updated\n", id)
	return nil
}
